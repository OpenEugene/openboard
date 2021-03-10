# Go and the Go Command Walk-through

## Introduction

- [The Go Project](https://golang.org/project)

## Installation

- [Download and Install](https://golang.org/doc/install)

## Operation

### The Basics

- [Go Tour](https://tour.golang.org)
- [Tutorial: Get started with Go](https://golang.org/doc/tutorial/getting-started)
- [Tutorial: Create a Go Module](https://golang.org/doc/tutorial/create-module)

### A Little More

- [Writing Web Applications](https://golang.org/doc/articles/wiki)
- [How to Write Go Code](https://golang.org/doc/code)

### Specifics

- [Command go](https://golang.org/cmd/go)

#### > `go mod`

Used to manage the module namespace of your project, as well as the list of
modules that your module is dependent on. 

Start your project:

```sh
go mod init github.com/OpenEugene/openboard
```

#### > `go get` and `mod` again

Used to obtain code from a given module to be added to your own module.

Add a dependency:

```sh
go get github.com/codemodus/sigmon/v2@latest
```

If you've modified your project imports and wish to update your list of modules,
`go mod` comes back into the picture.

Ensure dependencies are recorded properly:

```sh
go mod tidy
```

#### > `go test`

Used to run tests using the standard library test framework. Files with the
suffix `_test.go` will be picked up for handling.

Run tests:

```sh
go test # can be run with -v to see more info
```

#### > `go run`

Used to build and run a temporary executable of your application.

Build and run (temporary):

```sh
go run .
```

#### > `go build`

Used to build an executable of your application.

Build and run (persistent):

```sh
go build -o ./build/myproject
./myproject
```

#### > `go install`

Used to install an executable directly from a repository.

Install from source and run:

```sh
go install github.com/OpenEugene/openboard/back/cmd/openbsrv@latest
openbsrv
```

You will need to be sure that the install location is in your PATH environment
variable.

> Executables are installed in the directory named by the GOBIN environment
> variable, which defaults to $GOPATH/bin or $HOME/go/bin if the GOPATH
> environment variable is not set. Executables in $GOROOT are installed in
> $GOROOT/bin or $GOTOOLDIR instead of $GOBIN.
