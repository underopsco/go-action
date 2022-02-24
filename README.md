# Go Action

<a href="https://pkg.go.dev/github.com/crqra/go-action?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>

> GitHub Actions Toolkit for Go

### Goals

The goal is simple: provide a nice developer experience to people who want to author GitHub Actions
using Go while exploring how that experience can be improved.

### Features

- Functions for every [workflow command][workflow commands] (ex.: `action.GetInput()`, `action.SetInput()`)
- Easy access to the [GitHub Context][github context] (ex.: `action.Context.Token`, `action.Context.Ref`)
- Ready-to-use GitHub REST client (`action.REST`)
- Typed [events][events] (`action.GetEvent()`)
- _Experimental._ Automatic bindings of inputs and outputs (see [this example][example])

[workflow commands]: https://docs.github.com/en/actions/using-workflows/workflow-commands-for-github-actions
[github context]: https://docs.github.com/en/actions/learn-github-actions/environment-variables#default-environment-variables
[events]: https://docs.github.com/en/actions/using-workflows/events-that-trigger-workflows
[example]: https://github.com/crqra/go-action/blob/main/examples/hello_world_bind/main.go

_**This project is under active development and stability isn't guaranteed, use at your own discretion.**_

## Installation

```bash
$ go get github.com/crqra/go-action
```

## Examples

Check the [examples](examples) directory for a list of examples. You can also see
them being used in the [examples.yml][example workflow] workflow ([check runs][check runs]).

[example workflow]: .github/workflows/examples.yml
[check runs]: https://github.com/crqra/go-action/actions/workflows/examples.yml

## Documentation

Documentation is available on [pkg.go.dev](https://pkg.go.dev/github.com/crqra/go-action).

## Go Action in the Wild

- [conventional-commits-action](https://github.com/crqra/conventional-commits-action): Validate a Pull Request title and commit messages against Conventional Commits guidelines

## Acknowledgements

These are packages we either use or projects we are inspired by, and to whom we want to express
our gratitude:

- [google/go-github](https://github.com/google/go-github): Go library for accessing the GitHub API
- [sethvargo/go-githubactions](https://github.com/sethvargo/go-githubactions): Go SDK for GitHub Actions

## Contributing

We appreciate every contribution, thanks for considering it!

- [Open an issue][issues] if you have a problem or found a bug
- [Open a Pull Request][pulls] if you have a suggestion, improvement or bug fix

[issues]: https://github.com/crqra/go-action/issues
[pulls]: https://github.com/crqra/go-action/pulls

## License

This project is released under the [MIT License](LICENSE).