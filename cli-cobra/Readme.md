The purpose of this experiment is to determine a good baseline pattern for
writing CLI applications with [Cobra](https://github.com/spf13/cobra).

At a minimum, we should be able to:

1. Handle errors at any point of the command handling without stop-the-world
things like `os.exit` or something drastic like `panic`.
2. Read a configuration file that applies to the command being executed.
3. Avoid usage of globals as much as possible.

## Running

All of the following command lines are valid invocations of this experiment:

```sh
$ go run main.go help
$ go run main.go config dump -c ./sample.config.yaml
$ go run main.go config generate
$ go run main.go auth -c ./sample.config.yaml
$ CLI_CONFIG_FILE=./sample.config.yaml go run main.go auth
$ CLI_AUTH_KEY=123456 go run main.go auth
```
