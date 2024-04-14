package readingfiles

import (
	"io/fs"
	"testing/fstest"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func NewPostsFromFS(inputFs fstest.MapFS) []Post {
	var posts []Post

	dir, _ := fs.ReadDir(inputFs, ".")

	for range dir {
		posts = append(posts, Post{})
	}

	return posts
}
