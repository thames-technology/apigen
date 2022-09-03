/*
Copyright © 2022 Slavo Vojacek <public@slavovojacek.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package proto

import (
	"io"
	"os"
	"text/template"

	"github.com/slavovojacek/cloudstd/logger"
)

type ProtoTemplate struct {
	data      *templateData
	templates [3]*template.Template
}

// NewProtoTemplate creates the necessary template files and template data using the provided parameters.
func NewProtoTemplate(pkg, resource, resourcePlural, parent, parentPlural string) *ProtoTemplate {
	data := newTemplateData(pkg, resource, resourcePlural, parent, parentPlural)

	logger.Default().Infof("using template data %v", data)

	templates := [3]*template.Template{
		newResourceTemplate(),
		newRequestResponseTemplate(),
		newServiceTemplate(),
	}

	return &ProtoTemplate{
		data:      data,
		templates: templates,
	}
}

// WriteFiles creates the directory for the desired package and executes the templates using individual file writers.
func (t *ProtoTemplate) WriteFiles() error {
	pkgPath := t.data.PackagePath

	logger.Default().Infof("creating directory %s", pkgPath)

	if err := os.MkdirAll(pkgPath, 0770); err != nil {
		return err
	}

	for _, tmpl := range t.templates {
		if err := writeTemplate(tmpl, t.data); err != nil {
			return err
		}
	}

	return nil
}

// Write writes the templates using the provided io.Writer, for example, os.Stdout.
func (t *ProtoTemplate) Write(wr io.Writer) error {
	for _, tmpl := range t.templates {
		if err := tmpl.Execute(wr, t.data); err != nil {
			return err
		}
	}

	return nil
}
