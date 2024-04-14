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
		post, err := getPost(inputFS, f.Name())
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(inputFS fs.FS, inputName string) (Post, error) {
	postFile, err := inputFS.Open(inputName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	return newPost(postFile)
}

func newPost(input io.Reader) (Post, error) {
	postData, err := io.ReadAll(input)
	if err != nil {
		return Post{}, err
	}

	post := Post{Title: string(postData[7:])}
	
	return post, nil
}
