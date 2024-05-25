package template

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWrite(t *testing.T) {
	tests := []struct {
		name        string
		data        *ProtoData
		outDir      string
		expectedDir string
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
			outDir:      "testdata/output",
			expectedDir: "testdata/expected",
		},
		{
			name: "Test Author Service",
			data: &ProtoData{
				Resource:  "author",
				Resources: "authors",
				Package:   "authorservice.v1alpha1",
			},
			outDir:      "testdata/output",
			expectedDir: "testdata/expected",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Ensure the output directory exists
			if err := os.MkdirAll(tt.outDir, 0755); err != nil {
				t.Fatalf("Failed to create output directory: %v", err)
			}

			// Run the Write function
			if err := Write(tt.data, tt.outDir); err != nil {
				t.Fatalf("Write() error: %v", err)
			}

			// Compare the output file with the expected file
			outputFile := filepath.Join(tt.outDir, strings.ReplaceAll(tt.data.Package, ".", "/"), "service.proto")
			expectedFile := filepath.Join(tt.expectedDir, strings.ReplaceAll(tt.data.Package, ".", "/"), "service.proto")

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
			os.RemoveAll(tt.outDir)
		})
	}
}
