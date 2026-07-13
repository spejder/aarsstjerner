package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	onepassword "github.com/1password/onepassword-sdk-go"
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

	if account := ctx.String("1pass-account"); ctx.String("1pass") != "" && account != "" {
		fmt.Fprintf(os.Stderr, "Henter Medlemsservice-adgangskode for %s fra 1Password...\n", username)

		client, err := onepassword.NewClient(
			ctx.Context,
			onepassword.WithDesktopAppIntegration(account),
			onepassword.WithIntegrationInfo("aarsstjerner", getVersion()),
		)
		if err != nil {
			return "", "", fmt.Errorf("creating 1Password client: %w", err)
		}

		password, err = client.Secrets().Resolve(ctx.Context, ctx.String("1pass"))
		if err != nil {
			return "", "", fmt.Errorf("resolving 1Password secret: %w", err)
		}
	} else {
		bytePassword, err := gopass.GetPasswdPrompt("Password: ", true, os.Stdin, os.Stderr)
		if err != nil {
			log.Fatal(err)
		}

		password = string(bytePassword)
	}

	return username, password, nil
}
