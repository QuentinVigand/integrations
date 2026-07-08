package github

import (
	"context"
	"fmt"

	"github.com/PlakarKorp/kloset/connectors"
	"github.com/PlakarKorp/kloset/connectors/importer"
	"github.com/PlakarKorp/kloset/location"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type github struct {
	organization string
	client       *githubv4.Client
}

func (g *github) Origin() string { return "https://api.github.com" }

func (g *github) Type() string { return "github" }

func (g *github) Root() string { return "/" }

func (g *github) Flags() location.Flags { return 0 }

func (g *github) Ping(ctx context.Context) error {
	var ping QPing
	err := g.client.Query(ctx, &ping, nil)
	if err != nil {
		return fmt.Errorf("running ping query: %w", err)
	}
	// TODO handle invalid or expired token / missing permissions
	return nil
}

func (g *github) Import(context.Context, chan<- *connectors.Record, <-chan *connectors.Result) error {
	return nil
}
func (g *github) Close(context.Context) error {
	return nil
}

func init() {
	importer.Register("github", 0, NewImporter)
}

func NewImporter(ctx context.Context, opts *connectors.Options, str string, config map[string]string) (importer.Importer, error) {
	token, ok := config["token"]
	if !ok {
		return nil, fmt.Errorf("missing token in config")
	}
	org, ok := config["organization"]
	if !ok {
		return nil, fmt.Errorf("missing organization in config")
	}
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(ctx, src)

	return &github{
		organization: org,
		client:       githubv4.NewClient(httpClient),
	}, nil
}
