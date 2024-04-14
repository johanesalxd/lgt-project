package readingfiles_test

import (
	"testing"
	"testing/fstest"

	readingfiles "github.com/johanesalxd/lgt-project/reading_files"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello_world.md":   {Data: []byte("hi")},
		"hello_world_2.md": {Data: []byte("hola")},
	}

	posts := readingfiles.NewPostsFromFS(fs)

	got := len(posts)
	want := len(fs)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
