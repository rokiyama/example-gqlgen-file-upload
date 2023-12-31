package main

import (
	"os"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/rokiyama/example-gqlgen-file-upload/graph"
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

func TestMultipleUpload(t *testing.T) {
	var resp struct {
		MultipleUpload string
	}
	f1, err := os.Open("a.txt")
	require.NoError(t, err)
	f2, err := os.Open("b.txt")
	require.NoError(t, err)
	files := []*os.File{f1, f2}

	err = c.Post(`mutation ($files: [Upload!]!) {
		multipleUpload(files: $files)
	}`,
		&resp,
		client.Var("files", files),
		client.WithFiles(),
	)
	require.NoError(t, err)

	require.Equal(t, "success", resp.MultipleUpload)
}
