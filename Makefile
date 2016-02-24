PACKAGES=$(shell go list ./...)

all: format
	@go build -v ./...

format:
	@echo "--> Running go fmt"
	@go fmt ./...

test:
	@echo "--> Running tests"
	@go test -v ./...
	@$(MAKE) vet

coverage:
	@echo "--> Running tests with coverage"
	@echo "" > coverage.txt
	for pkg in $(shell go list ./...); do \
		go test -coverprofile=.pkg.coverage -covermode=atomic -v $$pkg ; \
		cat .pkg.coverage >> coverage.txt ; \
	done
	@rm .pkg.coverage
	@$(MAKE) vet

vet:
	@go tool vet 2>/dev/null ; if [ $$? -eq 3 ]; then \
    	go get golang.org/x/tools/cmd/vet; \
    fi
	@echo "--> Running go tool vet $(VETARGS)"
	@find . -name "*.go" | grep -v "./Godeps/" | xargs go tool vet $(VETARGS); if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for reviewal."; \
	fi
