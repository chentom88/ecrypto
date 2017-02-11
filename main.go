package main

import (
	"fmt"
	"os"

	"github.com/chentom88/ecrypto/base64"
	"github.com/chentom88/ecrypto/cipher"
	"github.com/chentom88/ecrypto/hex"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Usage = "A basic application for simple crypto"

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
						return runBasic(base64.Encode, c.Args().First())
					},
				},
				{
					Name:  "decode",
					Usage: "decode base64 to string",
					Action: func(c *cli.Context) error {
						return runBasic(base64.Decode, c.Args().First())
					},
				},
				{
					Name:  "encode_hex",
					Usage: "encode hex string to base64",
					Action: func(c *cli.Context) error {
						return runBasic(base64.EncodeHex, c.Args().First())
					},
				},
				{
					Name:  "decode_hex",
					Usage: "decode base64 to hex string",
					Action: func(c *cli.Context) error {
						return runBasic(base64.DecodeToHex, c.Args().First())
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
						return runBasic(hex.StringToHexString, c.Args().First())
					},
				},
				{
					Name:  "decode",
					Usage: "decode hex string to string",
					Action: func(c *cli.Context) error {
						return runBasic(hex.HexStringToString, c.Args().First())
					},
				},
				{
					Name:  "fixedxor",
					Usage: "perform a fixed xor on two hex strings of equivalent length",
					Action: func(c *cli.Context) error {
						output, err := hex.FixedXOR(c.Args().First(), c.Args().Get(1))
						if err == nil {
							fmt.Println(output)
						}

						return err
					},
				},
			},
		},
		{
			Name:    "sbxor",
			Aliases: []string{"sx"},
			Usage:   "single byte xor",
			Subcommands: []cli.Command{
				{
					Name:  "encrypt",
					Usage: "encode string to hex string and then encrypt with single byte",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "text",
							Usage: "the text to encrypt",
						},
						cli.StringFlag{
							Name:  "key",
							Value: "A",
							Usage: "the key to use for encryption, only the first byte will be used",
						},
					},
					Action: func(c *cli.Context) error {
						output, err := cipher.EncodeSingleByteXORString(c.String("text"), c.String("key"))
						if err == nil {
							fmt.Println(output)
						}

						return err
					},
				},
			},
		},
	}

	app.Run(os.Args)
}

func runBasic(operation func(string) (string, error), input string) error {
	output, err := operation(input)
	if err == nil {
		fmt.Println(output)
	}

	return err
}
