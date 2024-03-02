The purpose of this experiment is to explore utilizing interfaces for an
optional parameter (nil-able parameter). The issue explored is around nil
interface values:

+ https://go.dev/tour/methods/12
+ https://go.dev/tour/methods/13
+ https://web.archive.org/web/20231003131005/https://trstringer.com/go-nil-interface-and-interface-with-nil-concrete-value/
+ https://web.archive.org/web/20230603184210/https://mangatmodi.medium.com/go-check-nil-interface-the-right-way-d142776edef1
+ https://www.pixelstech.net/article/1554553347-Be-careful-about-nil-check-on-interface-in-GoLang

We are exploring this issue via a simple HTTP client with methods that can
accept query parameters, but do not require query parameters. Each method has
a distinct set of possible query parameters, but building the URL from the
query parameters is a common operation.
