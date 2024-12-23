# biscuit-dnd-go

WIP TUI character creation tool for 5e using the [5e-database](https://github.com/5e-bits/5e-database) that powers [dnd5eapi](https://www.dnd5eapi.co/). Name is also WIP.

## Local Development Setup

### Database setup

If you are not on an M series mac, you can run the database with just:

```sh
make db-up
```

This will pull the 5e-databases image, build it, and expose it on port `27017`.

If you are on an M series mac you will need to clone the database repository and build the image yourself:

1. Clone the database repository - https://github.com/5e-bits/5e-database
   - It's suggested to clone this one directory up from the biscuit-dnd-go directory.
2. In the `docker-compose.yaml` file, comment out the `image` block, and uncomment the `build` block.
   - If you cloned the database repository somewhere other than `../5e-database`, you'll need to update the `build` block to point to the right place.
3. Run `make db-up` to spin up the mongo database.

### Testing

This project uses the [testify](https://github.com/stretchr/testify) package for tests. To run all tests, either run:

```sh
make test
```

or:

```sh
go test ./...
```

To run only unit tests that don't require the mongodb:

```sh
make test-unit
```

or

```sh
go test -short ./...
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
