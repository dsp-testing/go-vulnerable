# go-vulnerable

This is an illustration of https://github.com/dependabot/dependabot-core/issues/4740

- The `false-positive` package imports a package that imports a vulnerable dependency in a test. 
- The Go tools have to download the `github.com/dsp-testing/go-vulnerable/vuln-too` package to check for usage, so it makes a `go.sum` entry.
- Since it found no imports except in tests, there is no entry in `go.mod`, thus it's not vulnerable. 
