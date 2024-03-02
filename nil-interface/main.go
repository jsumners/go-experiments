package main

import (
	"fmt"
	"github.com/gookit/goutil/dump"
	"ni/client"
)

func main() {
	c := client.NewClient("https://jsonplaceholder.typicode.com")

	posts := c.GetPosts(nil)
	dump.P(posts)

	fmt.Println("-----")

	photos := c.GetPhotos(nil)
	dump.P(photos)

	fmt.Println("-----")

	posts = c.GetPosts(&client.PostsParams{UserId: 2})
	dump.P(posts)

	fmt.Println("-----")

	photos = c.GetPhotos(&client.PhotosParams{AlbumId: 2})
	dump.P(photos)
}
