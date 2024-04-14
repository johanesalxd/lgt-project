package readingfiles

import (
	"io"
	"io/fs"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func NewPostsFromFS(inputFS fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(inputFS, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post

	for _, f := range dir {
		post, err := getPost(inputFS, f)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(inputFS fs.FS, f fs.DirEntry) (Post, error) {
	postFile, err := inputFS.Open(f.Name())
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	postData, err := io.ReadAll(postFile)
	if err != nil {
		return Post{}, err
	}

	post := Post{Title: string(postData[7:])}

	return post, nil
}
