package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func main() {
	server := initServer()
	defer server.Close()

	transport := &http.Transport{
		// This setting, MaxConnsPerHost, is what controls the pool limit, and
		// thus controls the concurrency limit _for specific http hosts_. By
		// default, the maximum _simultaneous_ connections per host is unlimited.
		// The docs for MaxConnsPerHost make it sound like it regulates the overall
		// maximum connections that will ever be sent to that host, but it is really
		// talking about simultaneous connections.
		MaxConnsPerHost: 2,
		// Technically, 2 is the default. It is only included here for clarity.
		MaxIdleConnsPerHost: 2,
	}
	client := http.Client{
		Transport: transport,
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 20; i += 1 {
		wg.Add(1)
		go func(i int, client http.Client) {
			defer wg.Done()
			fmt.Println(i, "request queued")
			res, err := client.Get("http://127.0.0.1:8080/")
			if err != nil {
				fmt.Println(i, err.Error())
				return
			}

			// Without the next two lines, running the test will result in only
			// two requests "finishing" (they will print "done"), but no other
			// requests completing. This is due to the fact that http.Client's
			// internal pooling does not consider a request to be fully handled
			// until the response body has been read and closed.
			//
			// In actuality, only one of either line is necessary on at least
			// Go 1.22. Try commenting out either line to see this. But if the
			// response body is not closed, resources will be leaked. So in practice,
			// both lines are necessary.
			defer res.Body.Close()
			io.Copy(io.Discard, res.Body)

			fmt.Println(i, "done")
		}(i, client)
	}
	wg.Wait()
}
