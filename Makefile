CURRENT_DIR := $(shell pwd)
BRANCH := $(shell git branch --show-current)
DAJE_TEST_NAME := daje-test-$(BRANCH)

clean:
	@echo "\n----- Start CLEAN -----\n"
	@echo "# Remove all the daje metadata and builds..."
	rm -r $(CURRENT_DIR)/bin &2>/dev/null
	@echo "\n----- End CLEAN -----"

# The version in this build has the form `<tag>|<commit_sha>`
build: clean
	@echo "\n----- Start BUILD -----\n"
	@echo "# Building daje for production environment..."
	go build -ldflags "-X github.com/Schroedinger-Hat/Daje/constants.Version=$(shell git describe --exact-match --tags $(git rev-parse HEAD) 2>/dev/null || git rev-parse --short HEAD)" -o ./bin/daje .
	@echo "\n----- End BUILD -----"

# For debugging reason the version in this build has the form `<current_branch>+<commit_sha>+<number_of_changes_not_committed>`
build-test-dev: clean
	@echo "\n----- Start BUILD-TEST-DEV -----\n"
	@echo "# Building daje for development environment..."
	go build -ldflags "-X github.com/Schroedinger-Hat/Daje/constants.ConfigBasepath=$(CURRENT_DIR)/testdata -X github.com/Schroedinger-Hat/Daje/constants.Version=$(BRANCH)+$(shell git rev-parse --short origin/main)+$(shell git status --porcelain | wc -l | tr -d ' ')" -o ./bin/$(DAJE_TEST_NAME) .
	@echo "\n----- End BUILD-TEST-DEV -----"

checkhealth: build-test-dev
	@echo "\n----- Start CHECKHEALTH -----\n"
	@echo "# run checkhealth command..."
	./bin/$(DAJE_TEST_NAME) checkhealth
	@echo "\n----- End CHECKHEALTH -----"
