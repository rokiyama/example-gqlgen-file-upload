package main

import (
	"os"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/rokiyama/example-gqlgen2/graph"
	"github.com/stretchr/testify/require"
)

var (
	c = client.New(handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: &graph.Resolver{}})))
)

func TestTodo(t *testing.T) {
	var resp struct {
		CreateTodo struct{ Text string }
	}
	c.MustPost(`mutation { createTodo(input:{text:"Hello",userId:"user1"}) { text } }`, &resp)

	require.Equal(t, "Hello", resp.CreateTodo.Text)
}

func TestUpload(t *testing.T) {
	var resp struct {
		SingleUpload string
	}
	f, err := os.Open("a.txt")
	require.NoError(t, err)

	err = c.Post(`mutation ($file: Upload!) {
		singleUpload(file: $file)
	}`, &resp, client.Var("file", f), client.WithFiles())
	require.NoError(t, err)

	require.Equal(t, "success", resp.SingleUpload)
}
