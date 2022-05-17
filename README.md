# go-vulnerable

This is an illustration of https://github.com/dependabot/dependabot-core/issues/4740

## false-positive

- The `false-positive` package imports a package that imports a vulnerable dependency in a test. 
- The Go tools have to download the `github.com/dsp-testing/go-vulnerable/vuln-too` package to check for usage, so it makes a `go.sum` entry.
- Since it found no imports except in tests, there is no entry in `go.mod`, thus it's not vulnerable. 

## vuln-too

- This is a good alert, since it imports for testing, the tests are vulnerable.

## vuln

- Interestingly, no alert is created for this situation which is vulnerable.

## safe

- This has no imports, so is safe. Sanity check!

## safe-too 

- Also safe, imports the safe package above.
