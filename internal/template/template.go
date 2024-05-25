package template

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed service.tmpl
var tmplFS embed.FS

var caser = cases.Title(language.Und)

var funcs = template.FuncMap{
	"title": caser.String,
}

type ProtoData struct {
	Resource  string
	Resources string
	Parent    string
	Parents   string
	Package   string
}

func Write(data *ProtoData, outDir string) error {
	tmpl, err := template.New("proto").Funcs(funcs).ParseFS(tmplFS, "service.tmpl")
	if err != nil {
		return err
	}

	// Construct the output file path
	outFilePath := fmt.Sprintf("%s/%s/%s", outDir, strings.ReplaceAll(data.Package, ".", "/"), "service.proto")

	// Create the directory structure if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(outFilePath), 0755); err != nil {
		return err
	}

	// Create the output file
	outFile, err := os.Create(outFilePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return tmpl.ExecuteTemplate(outFile, "service.tmpl", data)
}
