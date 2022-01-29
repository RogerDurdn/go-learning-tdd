package blogpost_test

import (
	blogpost "projectsBook/goTDD/reading-io/blogposts"
	"testing"
	"testing/fstest"
)

func TestNewBlogPost(t *testing.T) {
	fsystemTest := fstest.MapFS{
		"hello word.md":  {Data: []byte("hi")},
		"hello-word2.md": {Data: []byte("hola")},
	}
	posts := blogpost.NewPostFromFS(fsystemTest)
	got := len(posts)
	want := len(fsystemTest)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
