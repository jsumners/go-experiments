package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type QueryParams interface {
	Values() url.Values
}

type Client struct {
	BaseUrl url.URL

	http *http.Client
}

func NewClient(baseUrl string) *Client {
	result := &Client{
		http: &http.Client{},
	}
	parsedUrl, _ := url.Parse(baseUrl)
	result.BaseUrl = *parsedUrl
	return result
}

func (c *Client) get(path string, target any, params QueryParams) error {
	destUrl := c.BaseUrl
	destUrl.Path = destUrl.Path + path
	if params != nil {
		destUrl.RawQuery = params.Values().Encode()
	}

	req := &http.Request{
		Method:     "GET",
		URL:        &destUrl,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Host:       destUrl.Host,
	}
	fmt.Printf("%s %s\n", req.Method, req.URL.String())
	res, err := c.http.Do(req)

	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	if res.StatusCode > 399 {
		return fmt.Errorf("request to '%s' failed with '%s'", req.URL.String(), res.Status)
	}

	err = json.NewDecoder(res.Body).Decode(target)
	if err != nil {
		return fmt.Errorf("could not parse response: %w", err)
	}

	return nil
}

func (c *Client) GetPosts(params *PostsParams) []Post {
	var result []Post
	err := c.get("/posts", &result, params)
	if err != nil {
		panic(err)
	}
	return result
}

func (c *Client) GetPhotos(params *PhotosParams) []Photo {
	var result []Photo
	err := c.get("/photos", &result, params)
	if err != nil {
		panic(err)
	}
	return result
}
