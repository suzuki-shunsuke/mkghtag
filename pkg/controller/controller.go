package controller

import (
	"context"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/mkghtag/pkg/github"
)

type Controller struct {
	gh GitHub
}

type GitHub interface {
	CreateRef(ctx context.Context, owner string, repo string, ref *github.Reference) (*github.Reference, *github.Response, error)
	CreateTag(ctx context.Context, owner string, repo string, tag *github.Tag) (*github.Tag, *github.Response, error)
}

func New(ctx context.Context, url string) (*Controller, error) {
	gh, err := github.New(ctx, url)
	if err != nil {
		return nil, err
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

func stringP(s string) *string {
	return &s
}

func (c *Controller) Run(ctx context.Context, logE *logrus.Entry, param *ParamRun) error {
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

	logE.Info("creating a reference")
	_, _, err := c.gh.CreateRef(ctx, param.Owner, param.Repo, &github.Reference{
		Ref: stringP("refs/tags/" + param.Tag),
		Object: &github.GitObject{
			SHA: stringP(param.SHA),
		},
	})
	if err != nil {
		return fmt.Errorf("create a reference: %w", err)
	}
	if param.LightWeight {
		return nil
	}
	logE.Info("creating a tag")
	_, _, err = c.gh.CreateTag(ctx, param.Owner, param.Repo, &github.Tag{
		Tag:     stringP(param.Tag),
		SHA:     stringP(param.SHA),
		Message: stringP(param.Msg),
		Object: &github.GitObject{
			Type: stringP("commit"),
			SHA:  stringP(param.SHA),
		},
	})
	if err != nil {
		return fmt.Errorf("create a tag: %w", err)
	}
	return nil
}
