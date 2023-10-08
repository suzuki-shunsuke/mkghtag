# mkghtag

CLI to create GitHub Tags via API

## Install

- Homebrew
- aqua
- [GitHub Releases](https://github.com/suzuki-shunsuke/mkghtag/releases)

## Usage

mkghtag creates an annotated tag by default.
You can create a lightweight tag with `-light` option.

```
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
  -help
    	Show the help message
  -light
    	Create a lightweight tag
  -log-level string
    	Log Level (default "info")
  -msg string
    	Tag message
  -owner string
    	GitHub Repository owner
  -repo string
    	GitHub Repository name or full name <owner>/<repo>
  -sha string
    	Commit hash
  -version
    	Show the mkghtag's version
```

## Complement parameters based on CI specific environment variables

mkghtag complements parameters with [go-ci-env](https://github.com/suzuki-shunsuke/go-ci-env).

## LICENSE

[MIT](LICENSE)
