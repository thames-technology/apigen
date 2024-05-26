package template

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWriteTsRest(t *testing.T) {
	tests := []struct {
		name             string
		data             *TsRestTmplData
		fileName         string
		expectedFilePath string
	}{
		{
			name: "basic test with ulid",
			data: &TsRestTmplData{
				Resource:  "book",
				Resources: "books",
				Parent:    "author",
				IdType:    "ulid",
			},
			fileName:         "bookContract.ts",
			expectedFilePath: "../../examples/ts-rest/contracts/bookContract.ts",
		},
		{
			name: "test without parent",
			data: &TsRestTmplData{
				Resource:  "author",
				Resources: "authors",
				IdType:    "uuid",
			},
			fileName:         "authorContract.ts",
			expectedFilePath: "../../examples/ts-rest/contracts/authorContract.ts",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("write to stdout", func(t *testing.T) {
				var buf bytes.Buffer
				opts := &WriteOpts{Writer: &buf}

				err := WriteTsRest(tt.data, opts)
				if err != nil {
					t.Fatalf("WriteTsRest() error: %v", err)
				}

				expectedOutput, err := os.ReadFile(tt.expectedFilePath)
				if err != nil {
					t.Fatalf("Error reading expected output file: %v", err)
				}

				if diff := cmp.Diff(string(expectedOutput), buf.String()); diff != "" {
					t.Fatalf("Mismatch (-want +got):\n%s", diff)
				}
			})

			t.Run("write to file", func(t *testing.T) {
				tmpDir := t.TempDir()
				opts := &WriteOpts{OutDir: tmpDir}

				err := WriteTsRest(tt.data, opts)
				if err != nil {
					t.Fatalf("WriteTsRest() error: %v", err)
				}

				content, err := os.ReadFile(filepath.Join(tmpDir, tt.fileName))
				if err != nil {
					t.Fatalf("Error reading output file: %v", err)
				}

				expectedOutput, err := os.ReadFile(tt.expectedFilePath)
				if err != nil {
					t.Fatalf("Error reading expected output file: %v", err)
				}

				if diff := cmp.Diff(string(expectedOutput), string(content)); diff != "" {
					t.Fatalf("Mismatch (-want +got):\n%s", diff)
				}
			})
		})
	}
}

func TestWriteProto(t *testing.T) {
	tests := []struct {
		name             string
		data             *ProtoTmplData
		fileName         string
		expectedFilePath string
	}{
		{
			name: "basic test",
			data: &ProtoTmplData{
				Resource:  "book",
				Resources: "books",
				Parent:    "author",
				Package:   "bookservice.v1alpha1",
			},
			fileName:         "bookservice/v1alpha1/service.proto",
			expectedFilePath: "../../examples/proto/bookservice/v1alpha1/service.proto",
		},
		{
			name: "test without parent",
			data: &ProtoTmplData{
				Resource:  "author",
				Resources: "authors",
				Package:   "authorservice.v1alpha1",
			},
			fileName:         "authorservice/v1alpha1/service.proto",
			expectedFilePath: "../../examples/proto/authorservice/v1alpha1/service.proto",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("write to stdout", func(t *testing.T) {
				var buf bytes.Buffer
				opts := &WriteOpts{Writer: &buf}

				err := WriteProto(tt.data, opts)
				if err != nil {
					t.Fatalf("WriteProto() error: %v", err)
				}

				expectedOutput, err := os.ReadFile(tt.expectedFilePath)
				if err != nil {
					t.Fatalf("Error reading expected output file: %v", err)
				}

				if diff := cmp.Diff(string(expectedOutput), buf.String()); diff != "" {
					t.Fatalf("Mismatch (-want +got):\n%s", diff)
				}
			})

			t.Run("write to file", func(t *testing.T) {
				tmpDir := t.TempDir()
				opts := &WriteOpts{OutDir: tmpDir}

				err := WriteProto(tt.data, opts)
				if err != nil {
					t.Fatalf("WriteProto() error: %v", err)
				}

				content, err := os.ReadFile(filepath.Join(tmpDir, tt.fileName))
				if err != nil {
					t.Fatalf("Error reading output file: %v", err)
				}

				expectedOutput, err := os.ReadFile(tt.expectedFilePath)
				if err != nil {
					t.Fatalf("Error reading expected output file: %v", err)
				}

				if diff := cmp.Diff(string(expectedOutput), string(content)); diff != "" {
					t.Fatalf("Mismatch (-want +got):\n%s", diff)
				}
			})
		})
	}
}
