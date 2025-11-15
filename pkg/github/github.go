package github

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/v78/github"
	"golang.org/x/oauth2"
)

type (
	ListOptions   = github.ListOptions
	Reference     = github.Reference
	Response      = github.Response
	RepositoryTag = github.RepositoryTag
	Client        = github.Client
	GitObject     = github.GitObject
	Commit        = github.Commit
	Tag           = github.Tag
)

func New(ctx context.Context, url string) (*Client, error) {
	if url != "" {
		client, err := github.NewClient(getHTTPClientForGitHub(ctx, getGitHubToken())).WithEnterpriseURLs(url, url)
		if err != nil {
			return nil, fmt.Errorf("create a GHES client: %w", err)
		}
		return client, nil
	}
	return github.NewClient(getHTTPClientForGitHub(ctx, getGitHubToken())), nil
}

func getGitHubToken() string {
	return os.Getenv("GITHUB_TOKEN")
}

func getHTTPClientForGitHub(ctx context.Context, token string) *http.Client {
	if token == "" {
		return http.DefaultClient
	}
	return oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	))
}
