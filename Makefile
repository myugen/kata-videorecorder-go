download:
	@echo Download go.mod dependencies
	@go mod download

install-tools: download
	@echo Installing tools from tools.go
	$(eval TOOLS := $(shell go list -f '{{range .Imports}}{{.}} {{end}}' tools.go))
	@go install $(TOOLS)
