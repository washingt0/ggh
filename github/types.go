package github

import (
	"net/http"

	"github.com/shurcooL/githubv4"
)

type GitHub struct {
	GraphQL *githubv4.Client
}

type Issues struct {
	Nodes []struct {
		Title string
		ID    string
	}
}

type Repositories struct {
	Nodes []struct {
		Name   string
		Issues Issues `graphql:"issues(first:10, filterBy: {assignee: \"washingt0\"})"`
	}
}

type Organizations struct {
	Nodes []struct {
		Name         string
		Repositories Repositories `graphql:"repositories(first:30)"`
	}
}

func New(http *http.Client) *GitHub {
	return &GitHub{GraphQL: githubv4.NewClient(http)}
}
