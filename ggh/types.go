package ggh

type Config struct {
	MaxOrgs     int64  `json:"max_organizations"`
	MaxRepo     int64  `json:"max_repositories"`
	MaxIssues   int64  `json:"max_issues"`
	MaxPullR    int64  `json:"max_pull"`
	MaxIssLabel int64  `json:"max_issue_label"`
	Source      Source `json:"source"`
}

type Source struct {
	Name   string `json:"name"`
	Token  string `json:"token"`
	Client GGH
}

type User struct {
	ID       string
	Name     string
	URL      string
	Username string
}

type Organization struct {
}

type Repository struct {
}

type Issue struct {
}

type PullRequest struct {
}

type Comment struct {
}
