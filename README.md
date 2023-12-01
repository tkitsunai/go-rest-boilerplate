[![Go Report Card](https://goreportcard.com/badge/github.com/tkitsunai/go-rest-boilerplate)](https://goreportcard.com/report/github.com/tkitsunai/go-rest-boilerplate)

# go-rest-boilerplate

---

A greatest starting point for building REST APIs in Go.

We hope to assist you in swiftly constructing and developing robust Go applications.

While this project lacks most implementation code, it adheres to certain principles outlined in the directory and package structure. Moreover, the design of this project is based on the following concepts and will continuously evolve, requiring your adherence:

- Concepts from Clean Architecture
- Concepts from Domain Driven Design
- High testability
- Backend API for microservices

> [!IMPORTANT]
> Given the numerous elements essential to software architecture, a detailed explanation of each element is beyond the scope of this text.

## Features

- [x] Configuration using viper
- [x] CLI features using cobra
- [x] Structured logging using zerolog
- [x] Web framework using fiber
- [x] Testing using testify
- [x] CI using GitHub Actions workflow
- [x] Releaser using goreleaser
- [x] hot reloading using air

### Inline tool feature

- [ ] Code generate tool

## Running
before running, setup air
```bash
make setupAir
```
or
```bash
go install github.com/cosmtrek/air@latest
```

## Testing

```bash
make test
```

## Environment Variables

By default viper will look at $HOME/.go-rest-boilerplate.yaml for a config file.

### Folder Structure

```
go-rest-boilerplate/
├── cmd
│   └── api
│       └── server
└── internal
    ├── adapter
    ├── config
    ├── di
    ├── domain
    ├── driver
    ├── logger
    ├── port
    ├── rest
    │   ├── handlers
    │   └── routes
    └── usecase
```

TODO 英語文章にする

上記のディレクトリ・パッケージ群は基本的な構造であり、増やしても減らしても良い。ただし、プロジェクトが持つ基本設計方針に従っているべきである。

### cmd

TODO

### internal

internalパッケージは、外部に参照されるべきでないアプリケーション内のコードであることを明確にするためのパッケージです。
アプリケーション独自の全てのコードはこのパッケージに集約されるべきです。

特に重要なパッケージについての説明

- `adapter` package
- `domain` package
- `driver` package
- `port` package
- `usecase` package

## License

[MIT](./LICENSE)
