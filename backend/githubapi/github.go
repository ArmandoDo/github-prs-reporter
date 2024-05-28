package githubapi

import (
	"context"

	"github.com/google/go-github/github"
)

// Github Client Object
type Client struct {
	auth *github.Client // Github Client
}

// Configuration for Github API
type Config struct {
	Direction string // Sort direction
	Page      int    // Page of PR list
	PerPage   int    // Number of PRs per page
	RepoName  string // Name of repository
	RepoOwner string // Owner of repository
	Sort      string // Sort list
	State     string // State of PR
}

// Github pull requests object
type PullRequests struct {
	Pulls []*github.PullRequest // Pull requests interface
}

// GetPullRequests returns the list of PRs per page. It uses the Github API
func (cli *Client) GetPullRequests(conf *Config) (*PullRequests, error) {
	// Get list of pull request using pagination
	prs, _, err := cli.auth.PullRequests.List(context.Background(), conf.RepoOwner, conf.RepoName, githubOptions(conf))
	// Check for errors
	if err != nil {
		return nil, err
	}

	return &PullRequests{Pulls: prs}, nil
}

// NewClient creates a Github Client
func NewClient() *Client {
	return &Client{auth: github.NewClient(nil)}
}
