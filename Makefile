PKGS := ./

check:
	./bin/golangci-lint run $(PKGS) --timeout 5m

doc:
	gomarkdoc --output api.md -u .

fmt:
	gofumpt -w $(PKGS)

test: check
	gotestsum --format short-verbose -- -race -v  $(PKGS)

tools:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest
	go install gotest.tools/gotestsum@latest
	go install mvdan.cc/gofumpt@latest
	go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest