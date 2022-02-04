package reading_io_test

import (
	"reflect"
	reading_io "tdd/reading-io"
	"testing"
	"testing/fstest"
)

func TestNewBlogPost(t *testing.T) {
	const (
		firstBody = `Title: post 1
Description: description 1`
		secondBody = `Title: post 2
Description: description 2`
	)
	systemTestMap := fstest.MapFS{
		"hello word.md":  {Data: []byte("Title: post 1")},
		"hello-word2.md": {Data: []byte("Title: post 2")},
	}
	posts, err := reading_io.NewPostFromFS(systemTestMap)
	if err != nil {
		t.Fatalf("Error on newPost: %v", err)
	}
	assertPost(t, posts[0], reading_io.Post{Title: "post 1",
		Description: "description 1",
	})
}

func assertPost(t *testing.T, got reading_io.Post, want reading_io.Post) {
	t.Helper()
	if reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
