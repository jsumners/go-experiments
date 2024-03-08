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

## Notes

Commit https://github.com/jsumners/go-experiments/pull/1/commits/f48cf27d5383238b435a3a2c3480c2ac5b287300
shows an attempt at utilizing dependency injection for the various commands
the application provides. The issue with that commit is that the http client
will be `nil` when the `auth` command is attempted. This is because the closure
is established over the `nil` pointer, and the initialization is unable to
change that pointer to the real instance.

Commit https://github.com/jsumners/go-experiments/pull/1/commits/b5f22484b372ec8af6b5659a4c67fe05c2d5d5b9
shows a way to fix that problem. It creates a `CliApp` object that all of the
application's dependencies can be hung off of. Thereby allowing the application
initialization to update the fields on that object so that every closure can
utilize those fields instead of direct objects.

It's a little clunky in that each command's `New` method accepts a (potentially)
big object without a clear statement of what will be used out of it. But the
actual `run` functions can still reference the things that are needed. So if
we want to unit test any `run` functions, we can test them with the actual
dependency instead of the big `CliApp` container object.
