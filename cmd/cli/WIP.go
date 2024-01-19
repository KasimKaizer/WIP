package main

import (
	"log"
	"os"

	"github.com/KasimKaizer/WIP/internal/network"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = "Lets you do stuff with host"
	app.Version = "v0.0.1"

	cli.VersionFlag = cli.BoolFlag{
		Name:  "print-version, V",
		Usage: "print only the version",
	}

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "youtube.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ns",
			Usage: "Lookup the named server for the given host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				return network.NSLookup(c.String("host"))
			},
		},
		{
			Name:  "ip",
			Usage: "Lookup the IPs for the given host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				return network.IPLookup(c.String("host"))
			},
		},
		{
			Name:  "cname",
			Usage: "Lookup the CNAME for the given host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				return network.CnameLookup(c.String("host"))
			},
		},
		{
			Name:  "mx",
			Usage: "Lookup the MX Records for the given host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				return network.MXLookup(c.String("host"))
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
