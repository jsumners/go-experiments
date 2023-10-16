# http-iterator

The purpose of this experiment is to determine an ergonomic API for iterating
over paged results from a REST API.

In the future, we may be able to use https://bitfieldconsulting.com/golang/iterators
but that is not feasible at this time (2023-10). Thus, the implementation
herein is based upon:

+ https://github.com/googleapis/google-cloud-go/wiki/Iterator-Guidelines
+ https://go.dev/play/p/aFj6AvaBmIg
