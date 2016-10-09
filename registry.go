package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const (
	acceptHeader = "application/vnd.docker.distribution.manifest.v2+json"
)

// DebugHTTPRequests write dump of all http requests to stdout
var DebugHTTPRequests = true

// DebugHTTPResponses write dump of all http resposes to stdout
var DebugHTTPResponses = true

func debugDumpReq(req *http.Request) {
	if DebugHTTPRequests == false {
		return
	}
	b, err := httputil.DumpRequest(req, true)
	if err == nil {
		fmt.Println(string(b))
	} else {
		fmt.Println(err)
	}
}

func debugDumpRes(res *http.Response) {
	if DebugHTTPResponses == false {
		return
	}
	b, err := httputil.DumpResponse(res, true)
	if err == nil {
		fmt.Println(string(b))
	} else {
		fmt.Println(err)
	}
}

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

func (r *Registry) getImageListResponse() (io.ReadCloser, error) {
	requestURL := fmt.Sprintf("%s/v2/_catalog", r.url)
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", acceptHeader)

	debugDumpReq(req)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to get image list: %v", err)
	}
	debugDumpRes(res)
	return res.Body, nil
}

func (r *Registry) decodeImageListResponse(body io.Reader) ([]Image, error) {
	resjson := make(map[string]interface{})
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&resjson)
	if err != nil {
		return []Image{}, fmt.Errorf(
			"Failed to unmarshal ImageList Response: %v", err)
	}
	if irepos, ok := resjson["repositories"].([]interface{}); ok {
		ret := []Image{}
		for _, repo := range irepos {
			ret = append(ret, Image{Name: repo.(string)})
		}
		return ret, nil
	}
	return []Image{}, fmt.Errorf(
		"Malformed json in ImageList Response: %v", resjson)
}

func (r *Registry) getImageList() ([]Image, error) {
	resBody, err := r.getImageListResponse()
	if err != nil {
		return []Image{}, err
	}
	defer resBody.Close()
	return r.decodeImageListResponse(resBody)
}

// PrintImages prints images like docker images
func (r *Registry) PrintImages() error {
	images, err := r.getImageList()
	if err != nil {
		return err
	}
	for _, i := range images {
		tags, err := i.getTagList()
		if err != nil {
			return fmt.Errorf("Failed '%v' : %v", i.Name, err)
		}
		for _, t := range tags {
			fmt.Printf("%v\t%v\t%v\n", i.Name, t.Name, t.Digest)
		}
	}
	return nil
}
