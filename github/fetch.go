package github

import (
	"context"

	"github.com/washingt0/ggh/ggh"
)

func (g *GitHub) GetViewer(cfg *ggh.Config) (*ggh.User, error) {
	var u struct {
		Viewer struct {
			ID    string
			Name  string
			URL   string
			Login string
		}
	}

	err := g.GraphQL.Query(context.Background(), &u, nil)
	if err != nil {
		return nil, err
	}

	return &ggh.User{
		ID:       u.Viewer.ID,
		Name:     u.Viewer.Name,
		URL:      u.Viewer.URL,
		Username: u.Viewer.Login,
	}, nil
}

func (g *GitHub) GetOrganization() ([]ggh.Organization, error) {
	return []ggh.Organization{}, nil
}

func (g *GitHub) GetRepositories() ([]ggh.Repository, error) {
	return []ggh.Repository{}, nil
}
func (g *GitHub) GetIssues() ([]ggh.Issue, error) {
	return []ggh.Issue{}, nil
}
func (g *GitHub) GetPullRequests() ([]ggh.PullRequest, error) {
	return []ggh.PullRequest{}, nil
}
