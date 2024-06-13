# Development documentation

## Git Workflow

We use the [**Forking Workflow**](https://www.atlassian.com/git/tutorials/comparing-workflows/forking-workflow), so first of all you need a fork of the project and only then you can start contribute to the project.

## Branches

The main branch is the only one that are going to be used, no development. The reason is that we don't need it at this moment, we will consider adding more complexity later on if necessary.

## Requirements

- [go](https://go.dev/) version **1.18^**
- [gmake](https://www.gnu.org/software/make/)

## Setup environment

```
go mod tidy
```

## Run test with testdata

In the root folder, run:

```
make build-test-dev
```

This generates all the binary files necessary for daje, now you should find `daje` inside the `bin` directory. To check if everything is gone ok you can run:

```
./bin/daje checkhealt
```

## Version

Read this [https://go.dev/doc/modules/version-numbers](https://go.dev/doc/modules/version-numbers).

Every pull request merged to main will create a new version, this mean that every commit to the main branch will create a new version.

