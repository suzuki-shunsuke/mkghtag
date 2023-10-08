package controller

import (
	"context"

	"github.com/suzuki-shunsuke/mkghtag/pkg/github"
)

type Git interface {
	ListTags(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.RepositoryTag, *github.Response, error)
	GetCommitSHA1(ctx context.Context, owner, repo, ref, lastSHA string) (string, *github.Response, error)
}
