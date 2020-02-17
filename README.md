# gopaniccheck

`go/analysis.Analyzer` and tooling for reporting usage of panics in Go code.

## Install

### Local Install

Release binaries are available in the [Releases](https://github.com/bflad/gopaniccheck/releases) section.

To instead use Go to install into your `$GOBIN` directory (e.g. `$GOPATH/bin`):

```console
$ go get github.com/bflad/gopaniccheck/cmd/gopaniccheck
```

### Docker Install

```console
$ docker pull bflad/gopaniccheck
```

### Homebrew Install

```console
$ brew install bflad/tap/gopaniccheck
```

## Usage

Additional information about usage and configuration options can be found by passing the `help` argument:

```console
$ gopaniccheck help
```

### Local Usage

Change into the directory of the Terraform Provider code and run:

```console
$ gopaniccheck ./...
```

It is also possible to run via [`go vet`](https://golang.org/cmd/vet/):

```console
$ go vet -vettool $(which gopaniccheck) ./...
```

### Docker Usage

Change into the directory of the Terraform Provider code and run:

```console
$ docker run -v $(pwd):/src bflad/gopaniccheck ./...
```

## Reporting Analyzers

Reporting analyzers are all included in the `gopaniccheck` tool. For additional information about each check, you can run `gopaniccheck help NAME`.

| Analyzer | Description |
|---|---|
| `panic` | reports usage of `panic()` calls |
| `logpanic` | reports usage of `log.Panic()` calls |
| `logpanicf` | reports usage of `log.Panicf()` calls |
| `logpanicln` | reports usage of `log.Panicln()` calls |

## Result Analyzers

Result analyzers return the usage of panics. These are used by the reporting analyzers above, but can be used directly if building custom reporting.

| Check | Description |
|---|---|
| `paniccallexpr` | returns `panic()` `*ast.CallExpr` |
| `logpaniccallexpr` | returns `log.Panic()` `*ast.CallExpr` |
| `logpanicselectorexpr` | returns `log.Panic` `*ast.SelectorExpr` |
| `logpanicfcallexpr` | returns `log.Panicf()` `*ast.CallExpr` |
| `logpanicfselectorexpr` | returns `log.Panicf` `*ast.SelectorExpr` |
| `logpaniclncallexpr` | returns `log.Panicln()` `*ast.CallExpr` |
| `logpaniclnselectorexpr` | returns `log.Panicln` `*ast.SelectorExpr` |

## Development and Testing

This project is built on the [`go/analysis`](https://godoc.org/golang.org/x/tools/go/analysis) framework and uses [Go Modules](https://github.com/golang/go/wiki/Modules) for dependency management.

Helpful tooling for development:

* [`astdump`](https://github.com/wingyplus/astdump): a tool for displaying the AST form of Go file
* [`ssadump`](https://godoc.org/golang.org/x/tools/cmd/ssadump): a tool for displaying and interpreting the SSA form of Go programs

### Adding an Analyzer

* Create new analyzer in `passes/`
* If the `Analyzer` reports issues, add to `AllReportAnalyzers` variable in `passes/analyzers.go`

### Implementing a Custom Lint Tool

The `go/analysis` framework and this codebase are designed for flexibility. You may wish to permanently disable certain default checks or even implement your own provider-specific checks. An example of how to incorporate all default reporting analyzers in a CLI command can be found in `cmd/gopaniccheck`. To permanently exclude checks, each desired `Analyzer` must be individually included, similar to how `AllReportAnalyzers()` is built in `passes/analyzers.go`.

### Updating Dependencies

```console
$ go get URL
$ go mod tidy
```

### Unit Testing

```console
$ go test ./...
```

### Local Install Testing

```console
$ go install ./cmd/gopaniccheck
```
