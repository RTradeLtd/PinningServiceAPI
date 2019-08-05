.PHONY: gen
gen: clean gen-go-gin-server gen-go-server

.PHONY: clean
clean:
	rm -rf go
	rm -rf go-gin
	mkdir go
	mkdir go-gin

.PHONY: gen-go-gin-server
gen-go-gin-server:
	java -jar openapi-generator-cli.jar generate -i openapi_spec.yml  -g go-gin-server -o go-gin

.PHONY: gen-go-server
gen-go-server:
	java -jar openapi-generator-cli.jar generate -i openapi_spec.yml  -g go-gin-server -o go
