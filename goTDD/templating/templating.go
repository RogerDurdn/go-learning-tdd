package templating

import (
	"fmt"
	"io"
	blogpost "projectsBook/goTDD/reading-io/blogposts"
)

func main() {
}

func Render(writer io.Writer, post blogpost.Post) error  {
	_, err := fmt.Fprintf(writer,"<h1>%s</h1><p>%s</p>",
		post.Title, post.Description)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(writer, "Tags: <ul>")
	if err != nil {
		return err
	}
	for _, tag := range post.Tags {
		_, err = fmt.Fprintf(writer, "<li>%s</li>", tag)
	}
	_, err = fmt.Fprintf(writer, "</ul>")
	return nil
}
