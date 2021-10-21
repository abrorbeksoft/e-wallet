CURRENT_DIR=$(shell pwd)

start:
	 go run cmd/main.go

swag-gen:
		swag init -g api/main.go -o api/docs

.PHONY:compile