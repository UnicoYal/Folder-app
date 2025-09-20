COVERAGE_FILE = coverage.out
EXCLUDE_FILE = exclude_from_coverage.txt
COVERAGE_HTML = coverage.html
BUILDNUMBER ?= $(shell date +'%Y%m%d.%H%M')
COMMIT_HASH ?= $(shell git log -1 --pretty='format:%h')

build: 
	go build -v -o bin/folder \
       	-ldflags "-X github.com/folder-app/version.commitHash=$(COMMIT_HASH) \
                  -X github.com/folder-app/version.version=$(BUILDNUMBER)" \
    ./cmd/main.go

all: test coverage report

test:
	go test -coverprofile=$(COVERAGE_FILE) ./...

coverage: test
	./filter_coverage.sh $(COVERAGE_FILE) $(EXCLUDE_FILE)

report:
	go tool cover -func=$(COVERAGE_FILE)

html:
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	open $(COVERAGE_HTML)

clean:
	rm -f $(COVERAGE_FILE) $(COVERAGE_HTML)

EASYJSON = easyjson

easyjson:
	@while read -r file; do \
		echo "Processing $$file..."; \
		$(EASYJSON) -all $$file || exit 1; \
	done < easyjson.txt

GOLANGCI_LINT_VERSION ?= v1.64.8
GOBIN ?= $(shell go env GOPATH)/bin
GOLANGCI_LINT := $(GOBIN)/golangci-lint
LINT_CFG ?= .golangci.pipeline.yml
LINT_TIMEOUT ?= 30m

.PHONY: lint-install lint
lint-install:
	@echo ">> installing golangci-lint $(GOLANGCI_LINT_VERSION)"
	@GOBIN=$(GOBIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

lint: $(GOLANGCI_LINT)
	@echo ">> running golangci-lint"
	$(GOLANGCI_LINT) run -c $(LINT_CFG) --timeout $(LINT_TIMEOUT) ./...


MOCKERY_PKG ?= github.com/vektra/mockery/v3
MOCKERY_VER ?= v3.0.0
MOCKERY_CFG ?= .mockery.yaml

.PHONY: generate_mocks 
generate_mocks:
	@echo ">> generating mocks via $(MOCKERY_CFG)"
	@if command -v mockery >/dev/null 2>&1; then \
		mockery --config $(MOCKERY_CFG); \
	else \
		echo "mockery not found in PATH, running via go run ($(MOCKERY_VER))"; \
		go run $(MOCKERY_PKG)@$(MOCKERY_VER) --config $(MOCKERY_CFG); \
	fi

