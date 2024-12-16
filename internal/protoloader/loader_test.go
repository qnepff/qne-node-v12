package protoloader

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestProtoLoader(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/v1/protos/1":
			json.NewEncoder(w).Encode(Proto{
				ID:        1,
				Namespace: "test",
				Name:      "example.proto",
				Content:   `syntax = "proto3";

message Example {
  string name = 1;
}`,
				Version:   "v1",
			})
		case "/api/v1/protos":
			json.NewEncoder(w).Encode(map[string]interface{}{
				"protos": []Proto{
					{
						ID:        1,
						Namespace: "test",
						Name:      "example.proto",
						Content:   `syntax = "proto3";

message Example {
  string name = 1;
}`,
						Version:   "v1",
					},
					{
						ID:        2,
						Namespace: "test",
						Name:      "example2.proto",
						Content:   `syntax = "proto3";

message Example2 {
  string name = 1;
}`,
						Version:   "v1",
					},
				},
				"has_more": false,
				"last_id":  2,
			})
		case "/api/v1/namespaces":
			json.NewEncoder(w).Encode(map[string]interface{}{
				"namespaces": []string{"test", "prod"},
				"has_more":   false,
			})
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	// Create temporary directories for testing
	tempDir, err := os.MkdirTemp("", "protoloader_test_*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	cacheDir := filepath.Join(tempDir, "cache")
	compiledDir := filepath.Join(tempDir, "compiled")

	// Create a new ProtoLoader
	loader, err := New(server.URL, cacheDir, compiledDir)
	if err != nil {
		t.Fatal(err)
	}

	// Test GetProto
	t.Run("GetProto", func(t *testing.T) {
		proto, err := loader.GetProto(context.Background(), 1)
		if err != nil {
			t.Fatal(err)
		}

		if proto.ID != 1 || proto.Namespace != "test" || proto.Name != "example.proto" {
			t.Errorf("unexpected proto: %+v", proto)
		}

		// Test cache
		proto2, err := loader.GetProto(context.Background(), 1)
		if err != nil {
			t.Fatal(err)
		}

		if proto2.ID != proto.ID {
			t.Error("cache miss")
		}
	})

	// Test ListProtos
	t.Run("ListProtos", func(t *testing.T) {
		protos, hasMore, lastID, err := loader.ListProtos(context.Background(), "", 10, 0)
		if err != nil {
			t.Fatal(err)
		}

		if len(protos) != 2 {
			t.Errorf("expected 2 protos, got %d", len(protos))
		}

		if hasMore {
			t.Error("expected hasMore to be false")
		}

		if lastID != 2 {
			t.Errorf("expected lastID to be 2, got %d", lastID)
		}
	})

	// Test ListNamespaces
	t.Run("ListNamespaces", func(t *testing.T) {
		namespaces, hasMore, err := loader.ListNamespaces(context.Background(), 10, 1)
		if err != nil {
			t.Fatal(err)
		}

		if len(namespaces) != 2 {
			t.Errorf("expected 2 namespaces, got %d", len(namespaces))
		}

		if hasMore {
			t.Error("expected hasMore to be false")
		}

		if namespaces[0] != "test" || namespaces[1] != "prod" {
			t.Errorf("unexpected namespaces: %v", namespaces)
		}
	})

	// Test CompileProto
	t.Run("CompileProto", func(t *testing.T) {
		if os.Getenv("SKIP_PROTOC_TESTS") != "" {
			t.Skip("Skipping protoc tests")
		}

		err := loader.CompileProto(context.Background(), 1)
		if err != nil {
			t.Fatal(err)
		}

		outputDir, err := loader.GetCompiledProtoPath(1)
		if err != nil {
			t.Fatal(err)
		}

		if _, err := os.Stat(outputDir); err != nil {
			t.Errorf("compiled proto not found: %v", err)
		}
	})
}
