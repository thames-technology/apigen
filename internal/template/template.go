package template

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed proto/*.tmpl ts-rest/*.tmpl
var tmplFS embed.FS

var caser = cases.Title(language.Und)

var funcs = template.FuncMap{
	"title": caser.String,
}

type ProtoTmplData struct {
	Resource  string
	Resources string
	Parent    string
	Package   string
}

type TsRestTmplData struct {
	Resource  string
	Resources string
	Parent    string
	IdType    string
}

type WriteOpts struct {
	OutDir string
	Writer io.Writer
}

func WriteTsRest(data *TsRestTmplData, opts *WriteOpts) error {
	if opts.Writer != nil {
		return writeTsRestTo(opts.Writer, data)
	}

	outFilePath := filepath.Join(opts.OutDir, fmt.Sprintf("%sContract.ts", data.Resource))
	if err := os.MkdirAll(filepath.Dir(outFilePath), 0755); err != nil {
		return err
	}

	outFile, err := os.Create(outFilePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return writeTsRestTo(outFile, data)
}

func writeTsRestTo(w io.Writer, data *TsRestTmplData) error {
	tmpl, err := template.New("ts-rest").Funcs(funcs).ParseFS(tmplFS, "ts-rest/contract.ts.tmpl")
	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(w, "contract.ts.tmpl", data)
}

func WriteProto(data *ProtoTmplData, opts *WriteOpts) error {
	if opts.Writer != nil {
		return writeProtoTo(opts.Writer, data)
	}

	outFilePath := filepath.Join(opts.OutDir, strings.ReplaceAll(data.Package, ".", "/"), "service.proto")
	if err := os.MkdirAll(filepath.Dir(outFilePath), 0755); err != nil {
		return err
	}

	outFile, err := os.Create(outFilePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return writeProtoTo(outFile, data)
}

func writeProtoTo(w io.Writer, data *ProtoTmplData) error {
	tmpl, err := template.New("proto").Funcs(funcs).ParseFS(tmplFS, "proto/service.proto.tmpl")
	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(w, "service.proto.tmpl", data)
}
