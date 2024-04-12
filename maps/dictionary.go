package maps

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(keys string) (string, error) {
	val, ok := d[keys]

	if !ok {
		return "", ErrNotFound
	}

	return val, nil
}
