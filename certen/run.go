package certen

import (
	"certen"
	_ "certen/providers"
	"certen/utils"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
)

func Run() {
	app := cli.NewApp()
	app.Name = "certen"
	app.Version = "0.1.2"

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "source, s", Usage: "The source certifications directory"},
		cli.StringFlag{Name: "names, n", Value: "*", Usage: "The names of domains to export"},
		cli.StringFlag{Name: "output, o", Usage: "Output directory for certs"},
		cli.BoolTFlag{Name: "assemble, a", Usage: "Assemble certs to pem file"},
	}

	app.Action = func(c *cli.Context) error {
		output, err := utils.ResolvePath(c.String("output"))
		if err != nil {
			panic(err)
			return err
		}
		if !utils.Exists(output) {
			os.MkdirAll(output, os.ModePerm)
		}

		dir, err := utils.ResolvePath(c.String("source"))
		if err != nil {
			log.Fatal(err)
			return err
		}

		names := c.String("names")
		assemble := c.Bool("assemble")

		provider, err := certen.CreateProvider("caddy", dir)
		if err != nil {
			if dir == "" {
				log.Fatal("Can not find certs direcotry from defaults, Please specify the right dir by -s or --source")
			} else {
				log.Fatal("The source direcotry is not valid: " + dir)
			}
			return err
		}

		exported, err := provider.ExportByName(names, output, assemble)
		if err != nil {
			log.Fatal(err)
			return err
		}

		fmt.Printf("Exported certs of [%s] to \"%s\"", strings.Join(exported, ", "), output)
		return nil
	}

	app.Run(os.Args)
}
