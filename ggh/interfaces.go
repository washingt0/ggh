package ggh

type GGH interface {
	GetViewer(*Config) (*User, error)
	GetOrganization() ([]Organization, error)
	GetRepositories() ([]Repository, error)
	GetIssues() ([]Issue, error)
	GetPullRequests() ([]PullRequest, error)
}
