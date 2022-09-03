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
	"fmt"
	"os"

	"github.com/slavovojacek/cloudstd/proto"
	"github.com/spf13/cobra"
)

var _dryrun bool

// protoCmd represents the proto command
var protoCmd = &cobra.Command{
	Use:   "proto [pkg] [resource-name] [resource-plural-name]",
	Short: "Create standard Protobuf service definition",
	Args:  cobra.RangeArgs(2, 3),
	RunE: func(cmd *cobra.Command, args []string) error {
		pkg := args[0]
		resource := args[1]
		resources := fmt.Sprintf("%ss", resource)

		if len(args) == 3 {
			resources = args[2]
		}

		tmpl := proto.NewProtoTemplate(pkg, resource, resources)

		if _dryrun {
			writer := os.Stdout
			defer writer.Close()
			return tmpl.Write(writer)
		}

		return tmpl.WriteFiles()
	},
}

func init() {
	rootCmd.AddCommand(protoCmd)
	protoCmd.Flags().BoolVar(&_dryrun, "dry-run", false, "Dry run")
}
