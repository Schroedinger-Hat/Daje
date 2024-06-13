CURRENT_DIR := $(shell pwd)

# The version in this build has the form `<tag>|<commit_sha>`
build:
	@echo "PROD: Building daje..."
	go build -ldflags "-X github.com/Schrodinger-Hat/Daje/constants.Version=$(shell git describe --exact-match --tags $(git rev-parse HEAD) 2>/dev/null || git rev-parse --short HEAD)" -o ./bin/daje .

# For debugging reason the version in this build has the form `<current_branch>+<commit_sha>+<number_of_changes_not_committed>`
build-test-dev:
	@echo "TEST-DEV: Building daje for testing pourpose..."
	rm -rf $(CURRENT_DIR)/testdata
	go build -ldflags "-X github.com/Schrodinger-Hat/Daje/constants.DajeConfigBaseDir=$(CURRENT_DIR)/testdata -X github.com/Schrodinger-Hat/Daje/constants.Version=$(shell git branch --show-current)+$(shell git rev-parse --short origin/main)+$(shell git status --porcelain | wc -l)" -o ./bin/daje .

