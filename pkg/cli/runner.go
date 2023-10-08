package cli

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-ci-env/v3/cienv"
	"github.com/suzuki-shunsuke/mkghtag/pkg/controller"
	"github.com/suzuki-shunsuke/mkghtag/pkg/log"
)

type Runner struct {
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
	LDFlags *LDFlags
	LogE    *logrus.Entry
}

type LDFlags struct {
	Version string
	Commit  string
	Date    string
}

func (l *LDFlags) VersionString() string {
	if l == nil {
		return ""
	}
	if l.Version == "" {
		if l.Date == "" {
			return ""
		}
		return fmt.Sprintf("(%s)", l.Date)
	}
	if l.Date == "" {
		return l.Version
	}
	return fmt.Sprintf("%s (%s)", l.Version, l.Date)
}

type Flags struct {
	Owner       string
	Repo        string
	SHA         string
	Msg         string
	LogLevel    string
	Tag         string
	Version     bool
	Help        bool
	LightWeight bool
}

func setFlagVars(fs *flag.FlagSet, flags *Flags) {
	fs.StringVar(&flags.Owner, "owner", "", "GitHub Repository owner")
	fs.StringVar(&flags.Repo, "repo", "", "GitHub Repository name or full name <owner>/<repo>")
	fs.StringVar(&flags.SHA, "sha", "", "Commit hash")

	fs.StringVar(&flags.Msg, "msg", "", "Tag message")
	fs.BoolVar(&flags.LightWeight, "light", false, "Create a lightweight tag")

	fs.StringVar(&flags.LogLevel, "log-level", "info", "Log Level")

	fs.BoolVar(&flags.Version, "version", false, "Show the mkghtag's version")
	fs.BoolVar(&flags.Help, "help", false, "Show the help message")
}

func (r *Runner) Run(ctx context.Context, args ...string) error { //nolint:funlen,cyclop
	flags := &Flags{}
	if len(args) == 0 {
		return errors.New("arguments are required")
	}
	fs := flag.NewFlagSet(args[0], flag.ContinueOnError)
	setFlagVars(fs, flags)
	buf := &bytes.Buffer{}
	fs.SetOutput(buf)
	if err := fs.Parse(args[1:]); err != nil {
		if _, err := io.Copy(r.Stderr, buf); err != nil {
			return fmt.Errorf("output the parse error to stderr: %w", err)
		}
		return errors.New("")
	}
	flags.Tag = fs.Arg(0)

	if flags.Help {
		fs.PrintDefaults()
		fmt.Fprintf(r.Stderr, `mkghtag - Create a GitHub Tag via GitHub API

https://github.com/suzuki-shunsuke/mkghtag

Usage:
  mkghtag \
    [-owner <GitHub Repository Owner>] \
    [-repo <GitHub Repository name or full name>] \
    [-sha <commit hash>] \
    [-msg <Tag message>] \
    [-log-level <log level|info>] \
    [-light] \
    <tag>

  mkghtag -version
  mkghtag -help

Example:
  # Create an annotated tag
  mkghtag \
    -owner suzuki-shunsuke \
    -repo mkghtag \
    -sha c03b46bf86599637e7cb18884b0ee525e340977f \
    -msg hello \
    -log-level debug

  # Create a lightweight tag with "-light" option
  mkghtag \
    -owner suzuki-shunsuke \
    -repo mkghtag \
    -sha c03b46bf86599637e7cb18884b0ee525e340977f \
    -light

Options:
%s`, buf.String())
		return nil
	}

	if flags.Version {
		if s := r.LDFlags.VersionString(); s != "" {
			fmt.Fprintln(r.Stdout, s)
		}
		return nil
	}

	if flags.Tag == "" {
		return errors.New("tag is required")
	}

	if owner, repo, ok := strings.Cut(flags.Repo, "/"); ok {
		flags.Owner = owner
		flags.Repo = repo
	}

	if pt := cienv.Get(nil); pt != nil {
		if flags.Owner == "" {
			flags.Owner = pt.RepoOwner()
		}
		if flags.Repo == "" {
			flags.Repo = pt.RepoName()
		}
		if flags.SHA == "" {
			flags.SHA = pt.SHA()
		}
	}

	logE := r.LogE.WithFields(logrus.Fields{
		"repo_owner": flags.Owner,
		"repo_name":  flags.Repo,
		"sha":        flags.SHA,
		"tag":        flags.Tag,
	})

	ctrl := controller.New(ctx)
	log.SetLevel(flags.LogLevel, logE)
	param := &controller.ParamRun{
		Owner:       flags.Owner,
		Repo:        flags.Repo,
		SHA:         flags.SHA,
		Msg:         flags.Msg,
		Tag:         flags.Tag,
		LightWeight: flags.LightWeight,
	}
	return ctrl.Run(ctx, logE, param) //nolint:wrapcheck
}
