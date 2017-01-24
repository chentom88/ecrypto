package main

import (
	"fmt"
	"os"

	"github.com/chentom88/ecrypto/base64"
	"github.com/chentom88/ecrypto/hex"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:    "base64",
			Aliases: []string{"b64"},
			Usage:   "base64 operations",
			Subcommands: []cli.Command{
				{
					Name:  "encode",
					Usage: "encode string to base64",
					Action: func(c *cli.Context) error {
						return runBasicCommand(base64.Encode, c.Args().First())
					},
				},
				{
					Name:  "decode",
					Usage: "decode base64 to string",
					Action: func(c *cli.Context) error {
						return runBasicCommand(base64.Decode, c.Args().First())
					},
				},
				{
					Name:  "encode_hex",
					Usage: "encode hex string to base64",
					Action: func(c *cli.Context) error {
						return runBasicCommand(base64.EncodeHex, c.Args().First())
					},
				},
				{
					Name:  "decode_hex",
					Usage: "decode base64 to hex string",
					Action: func(c *cli.Context) error {
						return runBasicCommand(base64.DecodeToHex, c.Args().First())
					},
				},
			},
		},
		{
			Name:    "hex",
			Aliases: []string{"h"},
			Usage:   "hex string operations",
			Subcommands: []cli.Command{
				{
					Name:  "encode",
					Usage: "encode string to hex string",
					Action: func(c *cli.Context) error {
						return runBasicCommand(hex.StringToHexString, c.Args().First())
					},
				},
				{
					Name:  "decode",
					Usage: "decode hex string to string",
					Action: func(c *cli.Context) error {
						return runBasicCommand(hex.HexStringToString, c.Args().First())
					},
				},
			},
		},
	}

	app.Run(os.Args)
}

func runBasicCommand(operation func(string) (string, error), input string) error {
	output, err := operation(input)
	if err == nil {
		fmt.Println(output)
	}

	return err
}
