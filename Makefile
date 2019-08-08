.PHONY: careful
careful:
	@echo "careful or you'll wipe all the code"

.PHONY: gen
gen: clean gen-go-gin-server

.PHONY: clean
clean:
	rm -rf go
	
.PHONY: gen-go-gin-server
gen-go-gin-server:
	java -jar openapi-generator-cli.jar generate -i openapi_spec.yml  -g go-gin-server --package-name api -o generated


# run standard go tooling for better readability
.PHONY: tidy
tidy: imports fmt
	go vet ./...
	golint ./...

# automatically add missing imports
.PHONY: imports
imports:
	find . -type f -name '*.go' -exec goimports -w {} \;

# format code and simplify if possible
.PHONY: fmt
fmt:
	find . -type f -name '*.go' -exec gofmt -s -w {} \;
