package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"time"

	"bitbucket.org/long174/go-odoo"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/pkg/browser"
	"github.com/spejder/aarsstjerner/internal/ms"
	"github.com/urfave/cli/v2"
)

type otherInfo struct {
	Aarstjerne int `yaml:"årsstjerne"`
}

func getHTML(ctx *cli.Context) (string, error) {
	listInMarkdown, err := getMarkdown(ctx)
	if err != nil {
		return "", err
	}

	// create markdown parser with extensions
	p := parser.New()
	doc := p.Parse([]byte(listInMarkdown))

	// create HTML renderer with extensions
	htmlFlags := html.CompletePage | html.Smartypants
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	output := markdown.Render(doc, renderer)

	return string(output), nil
}

func runBrowser(ctx *cli.Context) error {
	output, err := getHTML(ctx)
	if err != nil {
		return fmt.Errorf("building html: %w", err)
	}

	tmpFile, err := os.CreateTemp("", "aarsstjerner.*.html")
	if err != nil {
		return fmt.Errorf("could not create temporary file: %w", err)
	}

	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(output)
	if err != nil {
		return fmt.Errorf("writing temporary file: %w", err)
	}

	err = tmpFile.Close()
	if err != nil {
		return fmt.Errorf("closing temporary file: %w", err)
	}

	err = browser.OpenFile(tmpFile.Name())
	if err != nil {
		return fmt.Errorf("opening browser: %w", err)
	}

	time.Sleep(1 * time.Second)

	return nil
}

func list(ctx *cli.Context) error {
	md, err := getMarkdown(ctx)
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, md)

	return nil
}

//nolint:funlen,cyclop
func getMarkdown(ctx *cli.Context) (string, error) {
	username, password, err := credentials(ctx)
	if err != nil {
		return "", fmt.Errorf("getting credentials: %w", err)
	}

	config := &odoo.ClientConfig{
		Admin:    username,
		Password: password,
		Database: ctx.String("ms-database"),
		URL:      ctx.String("ms-url"),
	}

	oc, err := odoo.NewClient(config)
	if err != nil {
		return "", fmt.Errorf("creating Odoo client: %w", err)
	}

	client := &ms.Client{Client: *oc}

	fmt.Fprintln(os.Stderr, "Henter medlemmer...")

	profiles, err := profiles(client)
	if err != nil {
		return "", err
	}

	fmt.Fprintln(os.Stderr, "Henter medlemskaber...")

	res := make(map[int][]string)

	ids := make([]int64, 0, len(*profiles))
	for _, profile := range *profiles {
		ids = append(ids, profile.MembershipIds.Get()...)
	}

	memberships, err := memberships(client, ids)
	if err != nil {
		return "", err
	}

	membershipsMap := make(map[int64]ms.MemberMembership, len(*memberships))
	for _, membership := range *memberships {
		membershipsMap[membership.Id.Get()] = membership
	}

	for _, profile := range *profiles {
		years, name := calculate(ctx, profile, membershipsMap)
		if years == 0 {
			continue
		}

		res[years] = append(res[years], name)
	}

	keys := make([]int, 0, len(res))
	for k := range res {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	slices.Reverse(keys)

	output := ""
	for _, years := range keys {
		output += fmt.Sprintf("\n# %d-årsstjerner (%d stk)\n", years, len(res[years]))

		for _, name := range res[years] {
			output += fmt.Sprintf("- %s\n", name)
		}
	}

	output += fmt.Sprintf("\n---\n_%s_\n", time.Now().Format("2006-01-02 15:04:05"))

	return output, nil
}
