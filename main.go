package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"

	"github.com/jcbritobr/sbox/sbox"
	"github.com/urfave/cli/v2"
)

const (
	applicationName  = "sbox"
	applicationUsage = "A symmetric secret box tool for seal and open documents"
)

var (
	key             string
	errWrongKeySize = errors.New("key size must be 32 bytes")
	flagFilename    = &cli.StringFlag{Name: "key", Aliases: []string{"k"}, Usage: "a 32 byte `KEY`", Destination: &key}

	commandOpen = cli.Command{
		Name:    "open",
		Aliases: []string{"o"},
		Usage:   "open a sealed message",
		Flags:   []cli.Flag{flagFilename},
		Action: func(ctx *cli.Context) error {
			if len(key) < sbox.KeySize {
				return errWrongKeySize
			}

			data, err := readFromStdin()
			if err != nil {
				return err
			}
			buffer := strKeyToArray(key)
			openData, err := sbox.Open(&buffer, data)
			if err != nil {
				return err
			}
			io.Copy(os.Stdout, bytes.NewBuffer(openData))

			return nil
		},
	}
	commandSeal = cli.Command{
		Name:    "seal",
		Aliases: []string{"s"},
		Usage:   "seal an open message",
		Flags:   []cli.Flag{flagFilename},
		Action: func(ctx *cli.Context) error {
			if len(key) < sbox.KeySize {
				return errWrongKeySize
			}

			data, err := readFromStdin()
			if err != nil {
				return err
			}
			buffer := strKeyToArray(key)
			cipherData, err := sbox.Seal(&buffer, data)
			if err != nil {
				return err
			}
			io.Copy(os.Stdout, bytes.NewBuffer(cipherData))
			return nil
		},
	}
)

func strKeyToArray(key string) [32]byte {
	var buffer [32]byte
	copy(buffer[:], key)
	return buffer
}

func readFromStdin() ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func buildApplication() *cli.App {
	app := &cli.App{
		Name:     applicationName,
		Usage:    applicationUsage,
		Commands: []*cli.Command{&commandOpen, &commandSeal},
	}

	return app
}

func main() {
	if err := buildApplication().Run(os.Args); err != nil {
		panic(err)
	}
}
