package main

import "errors"

var ErrNotFound = errors.New("Could not find the word")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
