[![Go Report Card](https://goreportcard.com/badge/github.com/tkitsunai/go-rest-boilerplate)](https://goreportcard.com/report/github.com/tkitsunai/go-rest-boilerplate)

# go-rest-boilerplate

A greatest starting point for building RESTful APIs in Go.
The implementation follows Clean Architecture principles as described by Uncle Bob.

## Features

- Configuration using viper
- CLI features using cobra
- Structured logging using zerolog
- Web framework using fiber
- Testing using testify
- CI using GitHub Actions workflow
- Releaser using goreleaser
- hot reloading using air

## Running

### running using "air"

before running, setup air
```bash
make setupAir
```
or
```bash
go install github.com/cosmtrek/air@latest
```

### simple running

```bash
make run
```

## Testing

```bash
make test
```

## Environment Variables

By default viper will look at $HOME/.go-rest-boilerplate.yaml for a config file. Setting your config as Environment Variables is recommended as by 12-Factor App.

