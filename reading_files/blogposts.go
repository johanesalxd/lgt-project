package readingfiles

import (
	"bufio"
	"io"
	"io/fs"
	"strings"
)

const (
	titleSep = "Title: "
	descSep  = "Description: "
	tagsSep  = "Tags: "
)

type Post struct {
	Title, Description string
	Tags               []string
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
	scanner := bufio.NewScanner(input)

	readLine := func(tag string) string {
		scanner.Scan()

		return strings.TrimPrefix(scanner.Text(), tag)
	}

	return Post{
		Title:       readLine(titleSep),
		Description: readLine(descSep),
		Tags:        strings.Split(readLine(tagsSep), ", "),
	}, nil
}
