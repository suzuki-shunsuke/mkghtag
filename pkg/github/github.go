package github

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/google/go-github/v74/github"
	"github.com/suzuki-shunsuke/ghtkn/pkg/api"
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

func New(ctx context.Context, logger *slog.Logger, url, clientID string) (*Client, error) {
	if url != "" {
		client, err := github.NewClient(getHTTPClientForGitHub(ctx, getGitHubToken())).WithEnterpriseURLs(url, url)
		if err != nil {
			return nil, fmt.Errorf("create a GHES client: %w", err)
		}
		return client, nil
	}
	if clientID == "" {
		return github.NewClient(getHTTPClientForGitHub(ctx, getGitHubToken())), nil
	}
	tm := api.New(api.NewInput())
	token, err := tm.Get(ctx, logger, &api.InputGet{
		ClientID:   clientID,
		UseKeyring: true,
	})
	if err != nil {
		return nil, fmt.Errorf("get token using client id: %w", err)
	}
	return github.NewClient(getHTTPClientForGitHub(ctx, token.AccessToken)), nil
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
