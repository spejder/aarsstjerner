package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/howeyc/gopass"
	"github.com/urfave/cli/v2"
)

func credentials(ctx *cli.Context) (string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	username := ctx.String("username")
	if username == "" {
		fmt.Fprint(os.Stderr, "Username: ")

		usernameInput, err := reader.ReadString('\n')
		if err != nil {
			return "", "", fmt.Errorf("reading username: %w", err)
		}

		username = strings.TrimSpace(usernameInput)
	}

	var password string

	if opPath, err := exec.LookPath("op"); err == nil && ctx.String("1pass") != "" {
		fmt.Fprintf(os.Stderr, "Henter Medlemsservice-adgangskode for %s fra 1Password...\n", username)

		bytePassword, err := exec.Command(opPath, "read", ctx.String("1pass"), "--no-newline").Output()
		if err != nil {
			log.Fatal(err)
		}

		password = string(bytePassword)
	} else {
		bytePassword, err := gopass.GetPasswdPrompt("Password: ", true, os.Stdin, os.Stderr)
		if err != nil {
			log.Fatal(err)
		}

		password = string(bytePassword)
	}

	return username, password, nil
}
