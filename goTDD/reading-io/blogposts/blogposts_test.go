package blogpost_test

import (
	blogpost "projectsBook/goTDD/reading-io/blogposts"
	"reflect"
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
	posts, err := blogpost.NewPostFromFS(systemTestMap)
	if err != nil {
		t.Fatalf("Error on newPost: %v", err)
	}
	assertPost(t, posts[0], blogpost.Post{Title: "post 1",
		Description: "description 1",
		})
}

func assertPost(t *testing.T, got blogpost.Post, want blogpost.Post)  {
	t.Helper()
	if reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
