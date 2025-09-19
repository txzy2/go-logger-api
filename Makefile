docs-gen:
	swag init -g cmd/app/main.go -o docs

.PHONY: docs-gen