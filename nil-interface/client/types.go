package client

import (
	"fmt"
	"net/url"
)

type Post struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Photo struct {
	Id           int    `json:"id"`
	AlbumId      int    `json:"albumId"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	ThumbnailUrl string `json:"thumbnailUrl"`
}

type PostsParams struct {
	UserId int `json:"userId"`
}

func (pp *PostsParams) Values() url.Values {
	result := url.Values{}
	result.Set("userId", fmt.Sprintf("%d", pp.UserId))
	return result
}

type PhotosParams struct {
	AlbumId int `json:"albumId"`
}

func (pp *PhotosParams) Values() url.Values {
	result := url.Values{}
	result.Set("albumId", fmt.Sprintf("%d", pp.AlbumId))
	return result
}
