package template

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWrite(t *testing.T) {
	expectedDir := "../../proto"

	tests := []struct {
		name string
		data *ProtoData
		opts *WriteOpts
	}{
		{
			name: "Test Book Service",
			data: &ProtoData{
				Resource:  "book",
				Resources: "books",
				Parent:    "author",
				Parents:   "authors",
				Package:   "bookservice.v1alpha1",
			},
			opts: &WriteOpts{Write: true, OutDir: "testdata"},
		},
		{
			name: "Test Author Service",
			data: &ProtoData{
				Resource:  "author",
				Resources: "authors",
				Package:   "authorservice.v1alpha1",
			},
			opts: &WriteOpts{Write: true, OutDir: "testdata"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Ensure the output directory exists
			if err := os.MkdirAll(tt.opts.OutDir, 0755); err != nil {
				t.Fatalf("Failed to create output directory: %v", err)
			}

			// Run the Write function
			if err := Write(tt.data, tt.opts); err != nil {
				t.Fatalf("Write() error: %v", err)
			}

			// Compare the output file with the expected file
			outputFile := filepath.Join(tt.opts.OutDir, strings.ReplaceAll(tt.data.Package, ".", "/"), "service.proto")
			expectedFile := filepath.Join(expectedDir, strings.ReplaceAll(tt.data.Package, ".", "/"), "service.proto")

			output, err := os.ReadFile(outputFile)
			if err != nil {
				t.Fatalf("Failed to read output file: %v", err)
			}

			expected, err := os.ReadFile(expectedFile)
			if err != nil {
				t.Fatalf("Failed to read expected file: %v", err)
			}

			if diff := cmp.Diff(string(output), string(expected)); diff != "" {
				t.Errorf("Mismatch (-output +expected):\n%s", diff)
			}

			// Clean up the output directory
			os.RemoveAll(tt.opts.OutDir)
		})
	}
}
