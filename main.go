package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"

	"github.com/qnepff/qne-node-v12/internal/rest"
)

const (
	addr = ":4445"
	gatewayURL = "https://qne.name" // QNE gateway server URL
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	nodeID int64
	nodeName string
	segmentID string
	certificate string
	restClient *rest.Client
	mu sync.RWMutex
)

type WebRTCMessage struct {
	Type      string      `json:"type"`
	Offer     interface{} `json:"offer,omitempty"`
	Answer    interface{} `json:"answer,omitempty"`
	Candidate interface{} `json:"candidate,omitempty"`
}

func start() error {
	// Initialize REST client if not already done
	if restClient == nil {
		restClient = rest.NewClient(gatewayURL)
	}

	// Register in a segment and get assigned a node ID and temporary name
	resp, err := restClient.RegisterInSegment()
	if err != nil {
		return fmt.Errorf("failed to register in segment: %v", err)
	}

	mu.Lock()
	nodeID = resp.NodeID
	nodeName = resp.NodeName
	segmentID = resp.SegmentID
	mu.Unlock()

	log.Printf("Registered with node ID %d and name '%s' in segment: %s", nodeID, nodeName, segmentID)

	// Get QNE certificate
	certResp, err := restClient.GetQNECertificate(nodeID, nodeName, segmentID)
	if err != nil {
		return fmt.Errorf("failed to get QNE certificate: %v", err)
	}

	mu.Lock()
	certificate = certResp.Certificate
	mu.Unlock()

	log.Printf("Retrieved QNE certificate")
	return nil
}

func restart() error {
	mu.Lock()
	// Clear existing state
	nodeID = 0
	nodeName = ""
	segmentID = ""
	certificate = ""
	mu.Unlock()

	// Start fresh
	return start()
}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	tlsConfig, err := generateTLSConfig()
	if err != nil {
		log.Fatalf("Failed to generate TLS config: %v", err)
	}

	quicConfig := &quic.Config{
		EnableDatagrams:       true,
		MaxIdleTimeout:        30 * time.Second,
		KeepAlivePeriod:      10 * time.Second,
		HandshakeIdleTimeout:  5 * time.Second,
		MaxIncomingStreams:   1000,
		MaxIncomingUniStreams: 1000,
		Allow0RTT:            true,
		Versions:             []quic.VersionNumber{quic.Version1},
	}

	mux := http.NewServeMux()

	// Handle WebSocket endpoint
	mux.HandleFunc("/ws", handleWebSocket)

	// Handle static files
	fileHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Add Alt-Svc header for HTTP/3
		w.Header().Set("Alt-Svc", fmt.Sprintf(`h3=":%s"`, addr))

		// Set content type based on file extension
		if strings.HasSuffix(r.URL.Path, ".css") {
			w.Header().Set("Content-Type", "text/css")
		} else if strings.HasSuffix(r.URL.Path, ".js") {
			w.Header().Set("Content-Type", "application/javascript")
		} else if strings.HasSuffix(r.URL.Path, ".html") {
			w.Header().Set("Content-Type", "text/html")
		}

		// Serve static files from the generated output directory
		fs := http.FileServer(http.Dir("frontend/.output/public"))
		fs.ServeHTTP(w, r)
	})

	mux.Handle("/", fileHandler)

	// Create HTTP/3 server
	http3Server := &http3.Server{
		Handler:         mux,
		Addr:           addr,
		QuicConfig:     quicConfig,
		TLSConfig:      tlsConfig,
		EnableDatagrams: true,
	}

	// Create HTTP/2 server
	http2Server := &http.Server{
		Addr:      addr,
		Handler:   mux,
		TLSConfig: tlsConfig,
	}

	// Start HTTP/3 server
	go func() {
		fmt.Printf("Starting HTTP/3 server on %s\n", addr)
		if err := http3Server.ListenAndServeTLS("localhost.crt", "localhost.key"); err != nil {
			log.Printf("HTTP/3 server error: %v", err)
		}
	}()

	// Start HTTP/2 server
	go func() {
		fmt.Printf("Starting HTTP/2 server on %s\n", addr)
		if err := http2Server.ListenAndServeTLS("localhost.crt", "localhost.key"); err != nil {
			log.Printf("HTTP/2 server error: %v", err)
		}
	}()

	if err := start(); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}

	<-sigChan
	fmt.Println("\nShutting down gracefully...")
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade WebSocket connection: %v", err)
		return
	}
	defer conn.Close()

	for {
		var msg WebRTCMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Failed to read WebSocket message: %v", err)
			break
		}

		// Echo the message back
		if err := conn.WriteJSON(msg); err != nil {
			log.Printf("Failed to write WebSocket message: %v", err)
			break
		}
	}
}

func generateTLSConfig() (*tls.Config, error) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(24 * time.Hour),
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
		DNSNames:     []string{"localhost"},
	}

	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		return nil, err
	}

	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	// Write the certificate and key to files
	if err := os.WriteFile("localhost.crt", certPEM, 0644); err != nil {
		return nil, fmt.Errorf("failed to write certificate file: %v", err)
	}
	if err := os.WriteFile("localhost.key", keyPEM, 0600); err != nil {
		return nil, fmt.Errorf("failed to write key file: %v", err)
	}

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return nil, err
	}

	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:  []string{"h3", "h2"},
	}, nil
}
