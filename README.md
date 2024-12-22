# biscuit-dnd-go

WIP TUI character creation tool for 5e using the [5e-database](https://github.com/5e-bits/5e-database) that powers [dnd5eapi](https://www.dnd5eapi.co/). Name is also WIP.

## Local Development Setup

### Database setup

If you are not on an M series mac, you can run the database with just:

```sh
docker run -p 27017:27017 ghcr.io/5e-bits/5e-database:latest
```

If you are on an M series mac or if you prefer to use the makefile to build & run the image then you will need to do the following:

1. Clone the database repository - https://github.com/5e-bits/5e-database
   - It's suggested to clone this one directory up from the biscuit-dnd-go directory.
2. Run `make startdb` to spin up the mongo database.
   - This will build the [5e-database Dockerfile](https://github.com/5e-bits/5e-database/blob/main/Dockerfile) and start it, exposing it on port 27017.
   - If you cloned the database repository somewhere other than `../5e-database`, you'll need to pass the path to the make command as `make startdb PATH=<path/to/5e-database>`

### Testing

This project uses the [testify](https://github.com/stretchr/testify) package for tests. To run all tests, either run:

```sh
make test
```

or:

```sh
go test ./...
```

### Linting

This project uses [golangci-lint](https://golangci-lint.run/) to run linters. If you are on a mac and have [homebrew](https://brew.sh/) installed, you can easily install golangci-lint by running `make lint` and entering `y` when asked if you would like to install. Otherwise, see their [install directions](https://golangci-lint.run/welcome/install/#local-installation).

Golangci-lint also offers integrations with many common editors. Check out their [docs](https://golangci-lint.run/welcome/integrations/) for more on how to set those integrations up.

### Build & Run

If you want to output a build, run

```sh
make build
```

to output the binary to `bin/main`. Of course you can also use `go build` directly if you wish to add any additional options.

To quickly run the program, run

```sh
make run
```
