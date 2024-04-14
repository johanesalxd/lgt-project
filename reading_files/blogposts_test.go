package readingfiles_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	readingfiles "github.com/johanesalxd/lgt-project/reading_files"
)

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("always fail")
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("error handling test", func(t *testing.T) {
		_, err := readingfiles.NewPostsFromFS(StubFailingFS{})

		if err == nil {
			t.Fatal(err)
		}
	})
	t.Run("content and file checking", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello_world.md":   {Data: []byte("Title: Post 1")},
			"hello_world_2.md": {Data: []byte("Title: Post 2")},
		}

		posts, err := readingfiles.NewPostsFromFS(fs)
		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d want %d", len(posts), len(fs))
		}

		assertPost(t, posts[0], readingfiles.Post{Title: "Post 1"})
	})
}

func assertPost(t *testing.T, got, want readingfiles.Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v want %+v", got, want)
	}
}
