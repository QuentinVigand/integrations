package github

import (
	"context"

	"github.com/PlakarKorp/kloset/connectors"
	"github.com/PlakarKorp/kloset/connectors/importer"
	"github.com/PlakarKorp/kloset/location"
)

type github struct {
}

func (g *github) Origin() string {
	return "TODO"
}

func (g *github) Type() string {
	return "TODO"
}

func (g *github) Root() string {
	return "TODO"
}

func (g *github) Flags() location.Flags {
	return 0
}

func (g *github) Ping(context.Context) error {
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

func NewImporter(context.Context, *connectors.Options, string, map[string]string) (importer.Importer, error) {
	return &github{}, nil
}
