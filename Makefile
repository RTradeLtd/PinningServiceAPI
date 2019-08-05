.PHONY: careful
careful:
	@echo "careful or you'll wipe all the code"

.PHONY: gen
gen: clean gen-go-gin-server

.PHONY: clean
clean:
	rm -rf go-gin
	mkdir go-gin

.PHONY: gen-go-gin-server
gen-go-gin-server:
	java -jar openapi-generator-cli.jar generate -i openapi_spec.yml  -g go-gin-server