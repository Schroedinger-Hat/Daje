CURRENT_DIR := $(shell pwd)
BRANCH := $(shell git branch --show-current)
DAJE_TEST_NAME := daje-test-$(BRANCH)

# Build Flags
FLAG_PREFIX := -X github.com/Schroedinger-Hat/Daje/constants
# TEST-DEV
FLAG_TEST-DEV_CONFIGBASEPATH := $(FLAG_PREFIX).ConfigBasepath=$(CURRENT_DIR)/testdata
FLAG_TEST-DEV_VERSION := $(FLAG_PREFIX).Version=$(BRANCH)+$(shell git rev-parse --short origin/main)+$(shell git status --porcelain | wc -l | tr -d ' ')
# PROD
FLAG_VERSION := $(FLAG_PREFIX).Version=$(shell git describe --exact-match --tags $(git rev-parse HEAD) 2>/dev/null || git rev-parse --short HEAD)

clean:
	@echo "\n----- Start CLEAN -----\n"
	@echo "# Remove all the daje metadata and builds..."
	rm -r $(CURRENT_DIR)/bin &2>/dev/null
	@echo "\n----- End CLEAN -----"

test:
	@echo "\n----- Start BUILD -----\n"
	@echo "# Running unit test..."
	go test ./...
	@echo "\n----- End BUILD -----"

# The version in this build has the form `<tag>|<commit_sha>`
build: test clean
	@echo "\n----- Start BUILD -----\n"
	@echo "# Building daje for production environment..."
	go build -ldflags "$(FLAG_VERSION)" -o ./bin/daje .
	@echo "\n----- End BUILD -----"

# For debugging reason the version in this build has the form `<current_branch>+<commit_sha>+<number_of_changes_not_committed>`
build-test-dev: test clean
	@echo "\n----- Start BUILD-TEST-DEV -----\n"
	@echo "# Building daje for development environment..."
	go build -ldflags "$(FLAG_TEST-DEV_CONFIGBASEPATH) $(FLAG_TEST-DEV_VERSION)" -o ./bin/$(DAJE_TEST_NAME) .
	@echo "\n----- End BUILD-TEST-DEV -----"

checkhealth: build-test-dev
	@echo "\n----- Start CHECKHEALTH -----\n"
	@echo "# run checkhealth command..."
	./bin/$(DAJE_TEST_NAME) checkhealth
	@echo "\n----- End CHECKHEALTH -----"
