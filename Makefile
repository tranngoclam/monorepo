COLOR := "\e[1;36m%s\e[0m\n"

.PHONY: golint
golint:
	@for module in $$(go list -f '{{.Dir}}' -m | xargs); do (printf $(COLOR) $$module && cd $$module && golangci-lint run ./... --fix); done

.PHONY: gogenerate
gogenerate:
	@for module in $$(go list -f '{{.Dir}}' -m | xargs); do (printf $(COLOR) $$module && go generate -C $$module); done

.PHONY: gomoddownload
gomoddownload:
	@for module in $$(go list -f '{{.Dir}}' -m | xargs); do (printf $(COLOR) $$module && go mod download -C $$module); done

.PHONY: gomodtidy
gomodtidy:
	@for module in $$(go list -f '{{.Dir}}' -m | xargs); do (printf $(COLOR) $$module && go mod tidy -C $$module); done

.PHONY: gotest
gotest:
	@for module in $$(go list -f '{{.Dir}}' -m | xargs); do (printf $(COLOR) $$module && go test -v -race $$module/...); done

.PHONY: goworksync
goworksync:
	@go work sync
