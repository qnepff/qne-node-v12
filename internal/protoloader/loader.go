package protoloader

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

type Proto struct {
	ID        int64  `json:"id"`
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Content   string `json:"content"`
	Version   string `json:"version"`
}

type ProtoLoader struct {
	serverURL    string
	cacheDir     string
	compiledDir  string
	protoCache   map[int64]*Proto
	cacheMutex   sync.RWMutex
	httpClient   *http.Client
}

func New(serverURL, cacheDir, compiledDir string) (*ProtoLoader, error) {
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create cache directory: %v", err)
	}
	if err := os.MkdirAll(compiledDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create compiled directory: %v", err)
	}

	return &ProtoLoader{
		serverURL:    serverURL,
		cacheDir:     cacheDir,
		compiledDir:  compiledDir,
		protoCache:   make(map[int64]*Proto),
		httpClient:   &http.Client{},
	}, nil
}

func (l *ProtoLoader) GetProto(ctx context.Context, id int64) (*Proto, error) {
	// Check memory cache first
	l.cacheMutex.RLock()
	if proto, ok := l.protoCache[id]; ok {
		l.cacheMutex.RUnlock()
		return proto, nil
	}
	l.cacheMutex.RUnlock()

	// Check disk cache
	cachePath := filepath.Join(l.cacheDir, fmt.Sprintf("proto_%d.json", id))
	if data, err := ioutil.ReadFile(cachePath); err == nil {
		var proto Proto
		if err := json.Unmarshal(data, &proto); err == nil {
			l.cacheMutex.Lock()
			l.protoCache[id] = &proto
			l.cacheMutex.Unlock()
			return &proto, nil
		}
	}

	// Fetch from server
	url := fmt.Sprintf("%s/api/v1/protos/%d", l.serverURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := l.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch proto: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("proto not found: %d", id)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned status %d", resp.StatusCode)
	}

	var proto Proto
	if err := json.NewDecoder(resp.Body).Decode(&proto); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	// Cache the proto
	data, err := json.Marshal(proto)
	if err == nil {
		_ = ioutil.WriteFile(cachePath, data, 0644)
	}

	l.cacheMutex.Lock()
	l.protoCache[id] = &proto
	l.cacheMutex.Unlock()

	return &proto, nil
}

func (l *ProtoLoader) CompileProto(ctx context.Context, id int64) error {
	proto, err := l.GetProto(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get proto: %v", err)
	}

	// Create a temporary directory for proto files
	tmpDir, err := ioutil.TempDir("", "proto_compile_*")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Write the proto file
	protoPath := filepath.Join(tmpDir, proto.Name)
	if err := ioutil.WriteFile(protoPath, []byte(proto.Content), 0644); err != nil {
		return fmt.Errorf("failed to write proto file: %v", err)
	}

	// Create output directory
	outputDir := filepath.Join(l.compiledDir, fmt.Sprintf("proto_%d", id))
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// Compile the proto file
	cmd := exec.CommandContext(ctx, "protoc",
		"--go_out="+outputDir,
		"--go_opt=paths=source_relative",
		"--go-grpc_out="+outputDir,
		"--go-grpc_opt=paths=source_relative",
		"--proto_path="+tmpDir,
		proto.Name,
	)

	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to compile proto: %v\nOutput: %s", err, output)
	}

	return nil
}

func (l *ProtoLoader) GetCompiledProtoPath(id int64) (string, error) {
	outputDir := filepath.Join(l.compiledDir, fmt.Sprintf("proto_%d", id))
	if _, err := os.Stat(outputDir); err != nil {
		return "", fmt.Errorf("compiled proto not found: %v", err)
	}
	return outputDir, nil
}

// ListProtos returns a list of available proto definitions
func (l *ProtoLoader) ListProtos(ctx context.Context, namespace string, pageSize int32, lastSeenID int64) ([]*Proto, bool, int64, error) {
	url := fmt.Sprintf("%s/api/v1/protos?page_size=%d&last_seen_id=%d", l.serverURL, pageSize, lastSeenID)
	if namespace != "" {
		url += "&namespace=" + namespace
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, false, 0, fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := l.httpClient.Do(req)
	if err != nil {
		return nil, false, 0, fmt.Errorf("failed to fetch protos: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, false, 0, fmt.Errorf("server returned status %d", resp.StatusCode)
	}

	var result struct {
		Protos   []*Proto `json:"protos"`
		HasMore  bool     `json:"has_more"`
		LastID   int64    `json:"last_id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, false, 0, fmt.Errorf("failed to decode response: %v", err)
	}

	return result.Protos, result.HasMore, result.LastID, nil
}

// ListNamespaces returns all available namespaces
func (l *ProtoLoader) ListNamespaces(ctx context.Context, pageSize int32, page int32) ([]string, bool, error) {
	url := fmt.Sprintf("%s/api/v1/namespaces?page_size=%d&page=%d", l.serverURL, pageSize, page)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, false, fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := l.httpClient.Do(req)
	if err != nil {
		return nil, false, fmt.Errorf("failed to fetch namespaces: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, false, fmt.Errorf("server returned status %d", resp.StatusCode)
	}

	var result struct {
		Namespaces []string `json:"namespaces"`
		HasMore    bool     `json:"has_more"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, false, fmt.Errorf("failed to decode response: %v", err)
	}

	return result.Namespaces, result.HasMore, nil
}
