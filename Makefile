SOURCES := ./...
BINARY := goscores
BINDIR := out
COVERAGE_REPORT := coverage.out

FLAGS := CGO_ENABLED=0

.DEFAULT_GOAL := all-noclean


tooling:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest


coverage: test-coverage
	@echo "Generating coverage report"
	go tool cover -html=$(COVERAGE_REPORT)


check-coverage: test-coverage
	@echo "Get total coverage"
	$(eval COVERED=$(shell go tool cover -func $(COVERAGE_REPORT) | grep total | awk '{print substr($$3, 1, length($$3)-1)}'))
	@if [ "80.0" != "$(word 1, $(sort 80.0 $(COVERED)))" ]; then \
		echo "Test coverage (${COVERED}%) is less than 80%";\
		exit 1;\
	fi


doc:
	@echo "Start godoc"
	godoc -http=:6060


all: clean fmt lint test build


all-noclean: fmt lint test build


clean:
	@echo "Removing $(BINDIR) and coverage report"
	@rm -rf $(BINDIR) $(COVERAGE_REPORT)
	@echo "Cleaning build and test cache"
	@go clean -cache -testcache


build: 
	@echo "Building binaries"
	mkdir -p ${BINDIR}
	${FLAGS} go build -buildvcs=true -o ${BINDIR} ${SOURCES}


generate:
	@echo "Generating source files"
	${FLAGS} go generate ${SOURCES}


fmt:
	@echo "Running gofmt"
	${FLAGS} gofmt -w .


install:
	@echo "Installing ${BINARY} to $(shell go env GOPATH)/bin"
	${FLAGS} go install  -buildvcs=true ${SOURCES}


lint:
	@echo "Running lint"
	golangci-lint run -v ${SOURCES}


run: 
	@echo "Running ${BINARY}"
	${FLAGS} go run -buildvcs=true ${SOURCES}


bench:
	@echo "Running benchs"
	${FLAGS} CGO_ENABLED=1 go test -v -bench=. -benchmem -race ${SOURCES}


test:
	@echo "Running tests"
	${FLAGS} CGO_ENABLED=1 go test -v -race ${SOURCES}


test-coverage:
	@echo "Running tests"
	${FLAGS} CGO_ENABLED=1 go test -v -cover -coverprofile=$(COVERAGE_REPORT) -race ${SOURCES}


tidy:
	@echo "Running mod tidy"
	${FLAGS} go mod tidy


--internal-update:
	@echo "Updating all dependencies"
	${FLAGS} go get -d -u -t ${SOURCES}


update: --internal-update vendor tidy


vendor:
	@echo "Running mod vendor"
	${FLAGS} go mod vendor

