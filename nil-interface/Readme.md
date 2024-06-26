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

## Explorations

1. The baseline experiment showing they problem is commit https://github.com/jsumners/go-experiments/tree/3326ede13ec28504b3e3d53eb9203a0662c2c38a/nil-interface
2. The issue being solved by reflection is commit https://github.com/jsumners/go-experiments/tree/3791e8cfb48f82a393512e006187f017f3c0c29d

## Benchmark

[./main_test.go](./main_test.go) has a benchmark that shows two ways of solving
the issue along with a benchmark of each. It shouldn't be any surprise that
using reflection to solve the issue is a less performant method, but it is
likely easier since the alternative method needs to know about all possible
implementations of the interface.
