package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/qnepff/qne-node-v12/internal/protoloader"
)

func main() {
	// Create cache and compiled directories in the user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get home directory: %v", err)
	}

	cacheDir := filepath.Join(homeDir, ".qne", "proto-cache")
	compiledDir := filepath.Join(homeDir, ".qne", "proto-compiled")

	// Create the proto loader
	loader, err := protoloader.New(
		"http://localhost:4444",
		cacheDir,
		compiledDir,
	)
	if err != nil {
		log.Fatalf("Failed to create proto loader: %v", err)
	}

	// Example: Download and compile a proto by ID
	protoID := int64(123)
	
	// This will download the proto if not cached
	proto, err := loader.GetProto(context.Background(), protoID)
	if err != nil {
		log.Fatalf("Failed to get proto: %v", err)
	}
	log.Printf("Got proto: %s (version %s)", proto.Name, proto.Version)

	// This will compile the proto if not already compiled
	if err := loader.CompileProto(context.Background(), protoID); err != nil {
		log.Fatalf("Failed to compile proto: %v", err)
	}

	// Get the path to the compiled proto
	compiledPath, err := loader.GetCompiledProtoPath(protoID)
	if err != nil {
		log.Fatalf("Failed to get compiled proto path: %v", err)
	}
	log.Printf("Compiled proto available at: %s", compiledPath)
}
