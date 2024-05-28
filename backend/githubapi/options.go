package githubapi

import "github.com/google/go-github/github"

// githubOptions returns an object with the Github options
func githubOptions(conf *Config) *github.PullRequestListOptions {
	return &github.PullRequestListOptions{
		State:     conf.State,
		Sort:      conf.Sort,
		Direction: conf.Direction,
		ListOptions: github.ListOptions{
			Page:    conf.Page,
			PerPage: conf.PerPage,
		},
	}
}
