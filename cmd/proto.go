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
package cmd

import (
	"os"

	"github.com/slavovojacek/cloudstd/internal/proto"
	"github.com/spf13/cobra"
)

var (
	packageName         string
	resourceNames       []string
	parentResourceNames []string
	dryrun              bool
	googleStyleguide    bool
)

var (
	styleguide = "modern"
)

// protoCmd represents the proto command
var protoCmd = &cobra.Command{
	Use:   "proto",
	Short: "Create standard Protobuf service definition",
	RunE: func(cmd *cobra.Command, args []string) error {
		resource, err := proto.NewResourceName(resourceNames...)
		if err != nil {
			return err
		}

		parentResource, err := proto.NewResourceName(parentResourceNames...)
		if err != nil {
			return err
		}

		if googleStyleguide {
			styleguide = "google"
		}

		tmpl := proto.NewProtoTemplate(packageName, resource, parentResource, styleguide)

		if dryrun {
			writer := os.Stdout
			defer writer.Close()
			return tmpl.Write(writer)
		}

		return tmpl.WriteFiles()
	},
}

func init() {
	rootCmd.AddCommand(protoCmd)
	protoCmd.Flags().StringVar(&packageName, "package", "", "package name, for example 'acme.resource.v1'")
	protoCmd.MarkFlagRequired("package")
	protoCmd.Flags().StringSliceVar(&resourceNames, "resource", []string{}, "resource name configuration, for example 'book'")
	protoCmd.MarkFlagRequired("resource")
	protoCmd.Flags().StringSliceVar(&parentResourceNames, "parent", []string{"example"}, "optional parent name configuration, for example 'shelf,shelves'")
	protoCmd.Flags().BoolVar(&dryrun, "dry-run", false, "execute as a dry run")
	protoCmd.Flags().BoolVar(&googleStyleguide, "google", false, "use the google styleguide")
}
