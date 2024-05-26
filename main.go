package main

import (
	"log"
	"os"

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
						Name:    "parents",
						Aliases: []string{"ps"},
						Usage:   "The plural name of the parent resource, e.g. authors",
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
					resource := c.String("resource")
					resources := c.String("resources")
					if resources == "" {
						resources = resource + "s"
					}

					parent := c.String("parent")
					parents := c.String("parents")
					if parent != "" && parents == "" {
						parents = parent + "s"
					}

					data := &template.ProtoData{
						Resource:  resource,
						Resources: resources,
						Parent:    parent,
						Parents:   parents,
						Package:   c.String("package"),
					}

					opts := &template.WriteOpts{
						OutDir: c.String("out-dir"),
						Write:  c.Bool("write"),
					}

					return template.Write(data, opts)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
