COLOR := "\e[1;36m%s\e[0m\n"

.PHONY: go-lint
go-lint:
	@for module in $$(go list -f '{{.Dir}}' -m | xargs); do (printf $(COLOR) $$module && cd $$module && golangci-lint run ./... --fix); done

.PHONY: go-generate
go-generate:
	@for module in $$(go list -f '{{.Dir}}' -m | xargs); do (printf $(COLOR) $$module && go generate -C $$module); done

.PHONY: go-mod-download
go-mod-download:
	@for module in $$(go list -f '{{.Dir}}' -m | xargs); do (printf $(COLOR) $$module && go mod download -C $$module); done

.PHONY: go-mod-tidy
go-mod-tidy:
	@for module in $$(go list -f '{{.Dir}}' -m | xargs); do (printf $(COLOR) $$module && go mod tidy -C $$module); done

.PHONY: go-test
go-test:
	@for module in $$(go list -f '{{.Dir}}' -m | xargs); do (printf $(COLOR) $$module && go test -v -race $$module/...); done

.PHONY: go-work-sync
go-work-sync:
	@go work sync
