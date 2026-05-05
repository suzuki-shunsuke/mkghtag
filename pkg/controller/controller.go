package controller

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/suzuki-shunsuke/mkghtag/pkg/github"
)

type Controller struct {
	gh GitHub
}

type GitHub interface {
	CreateRef(ctx context.Context, owner string, repo string, ref github.CreateRef) (*github.Reference, *github.Response, error)
	CreateTag(ctx context.Context, owner string, repo string, tag github.CreateTag) (*github.Tag, *github.Response, error)
}

func New(ctx context.Context, url string) (*Controller, error) {
	gh, err := github.New(ctx, url)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}
	return &Controller{
		gh: gh.Git,
	}, nil
}

type ParamRun struct {
	Owner       string
	Repo        string
	SHA         string
	Msg         string
	Tag         string
	LightWeight bool
}

func (c *Controller) Run(ctx context.Context, logger *slog.Logger, param *ParamRun) error {
	if param.Owner == "" {
		return errors.New("owner is required")
	}
	if param.Repo == "" {
		return errors.New("repo is required")
	}
	if param.SHA == "" {
		return errors.New("sha is required")
	}
	if param.Tag == "" {
		return errors.New("tag is required")
	}

	logger.Info("creating a reference")
	_, _, err := c.gh.CreateRef(ctx, param.Owner, param.Repo, github.CreateRef{
		Ref: "refs/tags/" + param.Tag,
		SHA: param.SHA,
	})
	if err != nil {
		return fmt.Errorf("create a reference: %w", err)
	}
	if param.LightWeight {
		return nil
	}
	logger.Info("creating a tag")
	_, _, err = c.gh.CreateTag(ctx, param.Owner, param.Repo, github.CreateTag{
		Tag:     param.Tag,
		Message: param.Msg,
		Object:  param.SHA,
		Type:    "commit",
	})
	if err != nil {
		return fmt.Errorf("create a tag: %w", err)
	}
	return nil
}
