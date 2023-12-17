package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
	"tailscale.com/atomicfile"
)

func editConfig(ctx *cli.Context) error {
	configFile := ctx.String("config")

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		err := createConfig(configFile)
		if err != nil {
			return err
		}
	}

	editCmd := ctx.String("editor") + " '" + configFile + "'"
	cmd := exec.Command("sh", "-c", editCmd)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run editor command: %w", err)
	}

	return nil
}

func createConfig(configFile string) error {
	//nolint:gomnd
	err := atomicfile.WriteFile(configFile, []byte{}, 0o600)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
