package templating_test

import (
	"bytes"
	blogpost "projectsBook/goTDD/reading-io/blogposts"
	"projectsBook/goTDD/templating"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogpost.Post{
			Title: "Hello",
			Body: "this is a body",
			Description: "This is description",
			Tags: []string{"go", "test"},
		}
	)

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := &bytes.Buffer{}
		err := templating.Render(buf, aPost)
		if err != nil {
			t.Fatal(err)
		}
		got := buf.String()
		want := `<h1>Hello</h1><p>This is description</p>Tags: <ul><li>go</li><li>test</li></ul>`
		if got != want {
			t.Errorf("got\n %s, want\n %s", got, want)
		}
	})

}
