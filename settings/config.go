package settings

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/washingt0/ggh/ggh"
	"github.com/washingt0/ggh/github"
	"golang.org/x/oauth2"
)

const (
	serviceGitHub = "github"
)

func NewFromFile(path string) (*ggh.Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg ggh.Config

	err = json.Unmarshal(data, &cfg)

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.Source.Token},
	)

	httpClient := oauth2.NewClient(context.Background(), src)

	switch cfg.Source.Name {

	case serviceGitHub:
		cfg.Source.Client = github.New(httpClient)
	}

	return &cfg, nil
}
