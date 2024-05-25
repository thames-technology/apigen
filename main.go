package main

import (
	"fmt"
	"log"
	"os"

	"github.com/thames-technology/apigen/internal/template"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "apigen",
		Usage: "Generate Protobuf API definitions using best practices",
		Commands: []*cli.Command{
			{
				Name: "proto",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "resource",
						Usage:    "The name of the resource, e.g. book",
						Required: true,
					},
					&cli.StringFlag{
						Name:  "resources",
						Usage: "The plural name of the resource, e.g. books",
					},
					&cli.StringFlag{
						Name:  "parent",
						Usage: "The name of the parent resource, e.g. author",
					},
					&cli.StringFlag{
						Name:  "parents",
						Usage: "The plural name of the parent resource, e.g. authors",
					},
					&cli.StringFlag{
						Name:     "package",
						Usage:    "The package name, e.g. bookservice.v1alpha1",
						Required: true,
					},
					&cli.StringFlag{
						Name:  "dir",
						Usage: "The directory to write the proto files to",
						Value: "proto",
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

					fmt.Printf("Data: %+v\n", data)

					return template.Write(data, c.String("dir"))
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
