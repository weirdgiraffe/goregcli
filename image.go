package main

import "fmt"

// Image structure represents image in Registry
type Image struct {
	Name string
}

func (i *Image) getTagList() ([]Tag, error) {
	return []Tag{}, fmt.Errorf("Not implemented")
}
