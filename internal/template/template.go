package template

import (
	"embed"
	"io"
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

type WriteOpts struct {
	OutDir string
	Write  bool
}

func Write(data *ProtoData, opts *WriteOpts) error {
	if !opts.Write {
		return writeTo(os.Stdout, data)
	}

	// Construct the output file path
	outFilePath := filepath.Join(opts.OutDir, strings.ReplaceAll(data.Package, ".", "/"), "service.proto")

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

	return writeTo(outFile, data)
}

func writeTo(w io.Writer, data *ProtoData) error {
	tmpl, err := template.New("proto").Funcs(funcs).ParseFS(tmplFS, "service.tmpl")
	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(w, "service.tmpl", data)
}
