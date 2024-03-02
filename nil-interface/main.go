package main

import (
	"fmt"
	"github.com/gookit/goutil/dump"
	"ni/client"
)

func main() {
	c := client.NewClient("https://jsonplaceholder.typicode.com")

	/* These will fail because the check for nil will not work. */
	//posts := c.GetPosts(nil)
	//dump.P(posts)
	//
	//fmt.Println("-----")
	//
	//photos := c.GetPhotos(nil)
	//dump.P(photos)

	/* These will work because a value is passed through. */
	posts := c.GetPosts(&client.PostsParams{UserId: 2})
	dump.P(posts)

	fmt.Println("-----")

	photos := c.GetPhotos(&client.PhotosParams{AlbumId: 2})
	dump.P(photos)
}
