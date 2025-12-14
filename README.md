# mkghtag

[![DeepWiki](https://img.shields.io/badge/DeepWiki-suzuki--shunsuke%2Fmkghtag-blue.svg?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACwAAAAyCAYAAAAnWDnqAAAAAXNSR0IArs4c6QAAA05JREFUaEPtmUtyEzEQhtWTQyQLHNak2AB7ZnyXZMEjXMGeK/AIi+QuHrMnbChYY7MIh8g01fJoopFb0uhhEqqcbWTp06/uv1saEDv4O3n3dV60RfP947Mm9/SQc0ICFQgzfc4CYZoTPAswgSJCCUJUnAAoRHOAUOcATwbmVLWdGoH//PB8mnKqScAhsD0kYP3j/Yt5LPQe2KvcXmGvRHcDnpxfL2zOYJ1mFwrryWTz0advv1Ut4CJgf5uhDuDj5eUcAUoahrdY/56ebRWeraTjMt/00Sh3UDtjgHtQNHwcRGOC98BJEAEymycmYcWwOprTgcB6VZ5JK5TAJ+fXGLBm3FDAmn6oPPjR4rKCAoJCal2eAiQp2x0vxTPB3ALO2CRkwmDy5WohzBDwSEFKRwPbknEggCPB/imwrycgxX2NzoMCHhPkDwqYMr9tRcP5qNrMZHkVnOjRMWwLCcr8ohBVb1OMjxLwGCvjTikrsBOiA6fNyCrm8V1rP93iVPpwaE+gO0SsWmPiXB+jikdf6SizrT5qKasx5j8ABbHpFTx+vFXp9EnYQmLx02h1QTTrl6eDqxLnGjporxl3NL3agEvXdT0WmEost648sQOYAeJS9Q7bfUVoMGnjo4AZdUMQku50McDcMWcBPvr0SzbTAFDfvJqwLzgxwATnCgnp4wDl6Aa+Ax283gghmj+vj7feE2KBBRMW3FzOpLOADl0Isb5587h/U4gGvkt5v60Z1VLG8BhYjbzRwyQZemwAd6cCR5/XFWLYZRIMpX39AR0tjaGGiGzLVyhse5C9RKC6ai42ppWPKiBagOvaYk8lO7DajerabOZP46Lby5wKjw1HCRx7p9sVMOWGzb/vA1hwiWc6jm3MvQDTogQkiqIhJV0nBQBTU+3okKCFDy9WwferkHjtxib7t3xIUQtHxnIwtx4mpg26/HfwVNVDb4oI9RHmx5WGelRVlrtiw43zboCLaxv46AZeB3IlTkwouebTr1y2NjSpHz68WNFjHvupy3q8TFn3Hos2IAk4Ju5dCo8B3wP7VPr/FGaKiG+T+v+TQqIrOqMTL1VdWV1DdmcbO8KXBz6esmYWYKPwDL5b5FA1a0hwapHiom0r/cKaoqr+27/XcrS5UwSMbQAAAABJRU5ErkJggg==)](https://deepwiki.com/suzuki-shunsuke/mkghtag)
[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/suzuki-shunsuke/mkghtag/badge)](https://scorecard.dev/viewer/?uri=github.com/suzuki-shunsuke/mkghtag) | [Install](INSTALL.md)

CLI to create GitHub Tags via API

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
