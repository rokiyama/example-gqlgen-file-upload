.PHONY: serve gqlgen test

serve:
	go run server.go

gqlgen:
	go run github.com/99designs/gqlgen

test:
	go test -v ./...
