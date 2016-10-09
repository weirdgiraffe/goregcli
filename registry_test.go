package main

import (
	"fmt"
	"testing"
)

var testRegistyUrls = []struct {
	url         string
	valid       bool
	expectedURL string
}{
	{"http://someurl.com", true, "http://someurl.com"},
	{"https://someurl.com", true, "https://someurl.com"},
	{"someurl.com", true, "https://someurl.com"},
	{"someurl", true, "https://someurl"},
	{"unknowschema://someurl", false, ""},
}

func TestRegistryCreationn(t *testing.T) {
	for _, tt := range testRegistyUrls {
		r, err := NewRegistry(tt.url)
		if err != nil {
			if tt.valid == true {
				t.Errorf("Failed for %v : %v", tt.url, err)
			}
			continue
		}
		if r == nil {
			t.Fatalf("Failed for %v : Registry is nil", tt.url)
		}
		if r.url != tt.expectedURL {
			t.Errorf("Failed for %v : '%v' != '%v'",
				tt.url,
				r.url,
				tt.expectedURL)
		}
	}
}

func TestRegistryImageList(t *testing.T) {
	r, err := NewRegistry("http://localhost:5000")
	if err != nil {
		t.Fatalf("Failed for create registry : %v", err)
	}

	l, err := r.getImageList()
	if err != nil {
		t.Errorf("Failed to get images list : %v", err)
	}
	if len(l) == 0 {
		t.Errorf("Empty repositories list")
	}
	fmt.Println(l)
}
