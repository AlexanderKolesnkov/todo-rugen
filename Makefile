.PHONY: run default

all: default

run:
	go run ./cmd/web -addr=$$TODO_ADDR

default:
	go run ./cmd/web