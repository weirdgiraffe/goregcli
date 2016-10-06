package main

import (
	"fmt"
	"net/url"
)

// Registry represents the whole Docker Registry
type Registry struct {
	url string
}

func fixScheme(urlscheme string) (string, error) {
	switch {
	case urlscheme == "http":
		return urlscheme, nil
	case urlscheme == "https":
		return urlscheme, nil
	case urlscheme == "":
		return "https", nil
	}
	return "", fmt.Errorf("Unsupported url schema '%v'", urlscheme)
}

// NewRegistry will allocate Registry (will not do any requests)
func NewRegistry(registryURL string) (*Registry, error) {
	u, err := url.Parse(registryURL)
	if err != nil {
		return nil, err
	}
	scheme, err := fixScheme(u.Scheme)
	if err != nil {
		return nil, err
	}
	u.Scheme = scheme
	if u.IsAbs() == false {
		return nil, fmt.Errorf("Not an absolute url '%v'", registryURL)
	}
	return &Registry{u.String()}, nil
}
