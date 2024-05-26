package main

import (
	"errors"
	"io"
	"log"
	"os"
	"strings"

	"github.com/thames-technology/apigen/internal/template"
	"github.com/urfave/cli/v2"
)

// Changed at build time via -ldflags
var version = "debug"

func main() {
	app := &cli.App{
		Name:    "apigen",
		Usage:   "Generate best-practice Protobuf APIs following design patterns",
		Version: version,
		Commands: []*cli.Command{
			{
				Name: "proto",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "resource",
						Aliases:  []string{"r"},
						Usage:    "The name of the resource, e.g. book",
						Required: true,
					},
					&cli.StringFlag{
						Name:    "resources",
						Aliases: []string{"rs"},
						Usage:   "The plural name of the resource, e.g. books",
					},
					&cli.StringFlag{
						Name:    "parent",
						Aliases: []string{"p"},
						Usage:   "The name of the parent resource, e.g. author",
					},
					&cli.StringFlag{
						Name:     "package",
						Aliases:  []string{"pkg"},
						Usage:    "The package name, e.g. bookservice.v1alpha1",
						Required: true,
					},
					&cli.StringFlag{
						Name:    "out-dir",
						Aliases: []string{"o"},
						Usage:   "The directory to write the proto files to",
						Value:   "proto",
					},
					&cli.BoolFlag{
						Name:    "write",
						Aliases: []string{"w"},
						Usage:   "Write the output to a file",
					},
				},
				Action: func(c *cli.Context) error {
					resource := strings.ToLower(c.String("resource"))
					resources := strings.ToLower(c.String("resources"))
					if resources == "" {
						resources = resource + "s"
					}

					data := &template.ProtoTmplData{
						Resource:  resource,
						Resources: resources,
						Parent:    strings.ToLower(c.String("parent")),
						Package:   strings.ToLower(c.String("package")),
					}

					var writer io.Writer
					if !c.Bool("write") {
						writer = os.Stdout
					}

					opts := &template.WriteOpts{
						OutDir: c.String("out-dir"),
						Writer: writer,
					}

					return template.WriteProto(data, opts)
				},
			},
			{
				Name: "ts-rest",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "resource",
						Aliases:  []string{"r"},
						Usage:    "The name of the resource, e.g. book",
						Required: true,
					},
					&cli.StringFlag{
						Name:    "resources",
						Aliases: []string{"rs"},
						Usage:   "The plural name of the resource, e.g. books",
					},
					&cli.StringFlag{
						Name:    "parent",
						Aliases: []string{"p"},
						Usage:   "The name of the parent resource, e.g. author",
					},
					&cli.StringFlag{
						Name:  "id",
						Usage: "The type of the ID field, ulid or uuid",
						Value: "ulid",
					},
					&cli.StringFlag{
						Name:    "out-dir",
						Aliases: []string{"o"},
						Usage:   "The directory to write the TypeScript files to",
						Value:   "contracts",
					},
					&cli.BoolFlag{
						Name:    "write",
						Aliases: []string{"w"},
						Usage:   "Write the output to a file",
					},
				},
				Action: func(c *cli.Context) error {
					resource := strings.ToLower(c.String("resource"))
					resources := strings.ToLower(c.String("resources"))
					if resources == "" {
						resources = resource + "s"
					}

					idType := c.String("id")
					if idType != "ulid" && idType != "uuid" {
						return errors.New("id must be ulid or uuid")
					}

					data := &template.TsRestTmplData{
						Resource:  resource,
						Resources: resources,
						Parent:    strings.ToLower(c.String("parent")),
						IdType:    idType,
					}

					var writer io.Writer
					if !c.Bool("write") {
						writer = os.Stdout
					}

					opts := &template.WriteOpts{
						OutDir: c.String("out-dir"),
						Writer: writer,
					}

					return template.WriteTsRest(data, opts)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
