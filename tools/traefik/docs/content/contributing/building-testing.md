---
title: "Traefik Building & Testing Documentation"
description: "Compile and test your own Traefik Proxy! Learn how to build your own Traefik binary from the sources, and read the technical documentation."
---

# Building and Testing

Compile and Test Your Own Traefik!
{: .subtitle }

You want to build your own Traefik binary from the sources?
Let's see how.

## Building

You need either [Docker](https://github.com/docker/docker "Link to website of Docker") and `make` (Method 1), or [Go](https://go.dev/ "Link to website of Go") (Method 2) in order to build Traefik.
For changes to its dependencies, the `dep` dependency management tool is required.

### Method 1: Using `Docker` and `Makefile`

Run make with the `binary` target.

```bash
make binary
```

This will create binaries for the Linux platform in the `dist` folder.

In case when you run build on CI, you may probably want to run docker in non-interactive mode. To achieve that define `DOCKER_NON_INTERACTIVE=true` environment variable.

```bash
$ make binary
docker build -t traefik-webui -f webui/Dockerfile webui
Sending build context to Docker daemon  2.686MB
Step 1/11 : FROM node:8.15.0
 ---> 1f6c34f7921c
[...]
Successfully built ce4ff439c06a
Successfully tagged traefik-webui:latest
[...]
docker build  -t "traefik-dev:4475--feature-documentation" -f build.Dockerfile .
Sending build context to Docker daemon    279MB
Step 1/10 : FROM golang:1.16-alpine
 ---> f4bfb3d22bda
[...]
Successfully built 5c3c1a911277
Successfully tagged traefik-dev:4475--feature-documentation
docker run  -e "TEST_CONTAINER=1" -v "/var/run/docker.sock:/var/run/docker.sock" -it -e OS_ARCH_ARG -e OS_PLATFORM_ARG -e TESTFLAGS -e VERBOSE -e VERSION -e CODENAME -e TESTDIRS -e CI -e CONTAINER=DOCKER		 -v "/home/ldez/sources/go/src/github.com/traefik/traefik/"dist":/go/src/github.com/traefik/traefik/"dist"" "traefik-dev:4475--feature-documentation" ./script/make.sh generate binary
---> Making bundle: generate (in .)
removed 'autogen/genstatic/gen.go'

---> Making bundle: binary (in .)

$ ls dist/
traefik*
```

The following targets can be executed outside Docker by setting the variable `IN_DOCKER` to an empty string (although be aware that some of the tests might fail in that context):

- `test-unit`
- `test-integration`
- `validate`
- `binary` (the webUI is still generated by using Docker)

ex:

```bash
IN_DOCKER= make test-unit
```

### Method 2: Using `go`

Requirements:

- `go` v1.16+
- environment variable `GO111MODULE=on`

!!! tip "Source Directory"

    It is recommended that you clone Traefik into the `~/go/src/github.com/traefik/traefik` directory.
    This is the official golang workspace hierarchy that will allow dependencies to be properly resolved.

!!! note "Environment"

    Set your `GOPATH` and `PATH` variable to be set to `~/go` via:

    ```bash
    export GOPATH=~/go
    export PATH=$PATH:$GOPATH/bin
    ```

    For convenience, add `GOPATH` and `PATH` to your `.bashrc` or `.bash_profile`

    Verify your environment is setup properly by running `$ go env`.
    Depending on your OS and environment, you should see an output similar to:

    ```bash
    GOARCH="amd64"
    GOBIN=""
    GOEXE=""
    GOHOSTARCH="amd64"
    GOHOSTOS="linux"
    GOOS="linux"
    GOPATH="/home/<yourusername>/go"
    GORACE=""
    ## ... and the list goes on
    ```

#### Build Traefik

Once you've set up your go environment and cloned the source repository, you can build Traefik.

```bash
# Generate UI static files
make clean-webui generate-webui

# required to merge non-code components into the final binary,
# such as the web dashboard/UI
go generate
```

```bash
# Standard go build
go build ./cmd/traefik
```

You will find the Traefik executable (`traefik`) in the `~/go/src/github.com/traefik/traefik` directory.

## Testing

### Method 1: `Docker` and `make`

Run unit tests using the `test-unit` target.
Run integration tests using the `test-integration` target.
Run all tests (unit and integration) using the `test` target.

```bash
$ make test-unit
docker build -t "traefik-dev:your-feature-branch" -f build.Dockerfile .
# […]
docker run --rm -it -e OS_ARCH_ARG -e OS_PLATFORM_ARG -e TESTFLAGS -v "/home/user/go/src/github/traefik/traefik/dist:/go/src/github.com/traefik/traefik/dist" "traefik-dev:your-feature-branch" ./script/make.sh generate test-unit
---> Making bundle: generate (in .)
removed 'gen.go'

---> Making bundle: test-unit (in .)
+ go test -cover -coverprofile=cover.out .
ok      github.com/traefik/traefik   0.005s  coverage: 4.1% of statements

Test success
```

For development purposes, you can specify which tests to run by using (only works the `test-integration` target):

```bash
# Run every tests in the MyTest suite
TESTFLAGS="-check.f MyTestSuite" make test-integration

# Run the test "MyTest" in the MyTest suite
TESTFLAGS="-check.f MyTestSuite.MyTest" make test-integration

# Run every tests starting with "My", in the MyTest suite
TESTFLAGS="-check.f MyTestSuite.My" make test-integration

# Run every tests ending with "Test", in the MyTest suite
TESTFLAGS="-check.f MyTestSuite.*Test" make test-integration
```

Check [gocheck](https://labix.org/gocheck "Link to website of gocheck") for more information.

### Method 2: `go`

Unit tests can be run from the cloned directory using `$ go test ./...` which should return `ok`, similar to:

```test
ok      _/home/user/go/src/github/traefik/traefik    0.004s
```

Integration tests must be run from the `integration/` directory and require the `-integration` switch: `$ cd integration && go test -integration ./...`.