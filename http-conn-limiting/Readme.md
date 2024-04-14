# HTTP Connection Limiting

This experiment is an illustration of simultaneous HTTP connection limiting.
The [limiting-concurrency](../limiting-concurrency) experiment is a more
general experiment that was conducted to solve the problem of this
experiment. However, I have since learned that the standard Go HTTP library
supports connection limiting natively through
[http.Transport](https://pkg.go.dev/net/http#Transport).

The experiment can be run by:

```sh
$ go run .
```

## Related Articles

+ https://www.sobyte.net/post/2022-03/go-http-client-connection-control/
  ([archive](https://web.archive.org/web/20240229112402/https://www.sobyte.net/post/2022-03/go-http-client-connection-control/))
+ https://blog.cubieserver.de/2022/http-connection-reuse-in-go-clients/
([archive](https://web.archive.org/web/20230518081337/https://blog.cubieserver.de/2022/http-connection-reuse-in-go-clients/))
+ https://www.loginradius.com/blog/engineering/tune-the-go-http-client-for-high-performance/
([archive](https://web.archive.org/web/20240223160724/https://www.loginradius.com/blog/engineering/tune-the-go-http-client-for-high-performance/))
