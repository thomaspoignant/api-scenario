# Contribution Guidelines
We love to see contributions to the project and have tried to make it easy to 
do so. If you would like to contribute code to this project you can do so 
through GitHub by [forking the repository and sending a pull request](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request-from-a-fork).

We want to make contributing to this project as easy and transparent as possible, whether it's:

- Reporting a bug
- Discussing the current state of the code
- Submitting a fix
- Proposing new features
- Becoming a maintainer


# Setup your machine
`api-scenario` is written in [Go](https://golang.org/).

## Prerequisites:
- `make`
- [Go 1.14+](https://golang.org/doc/install)

## Build project
Clone `api-scenario` anywhere:
```shell script
$ git clone git@github.com:thomaspoignant/api-scenario.git
```
Install the build:
```shell script
$ make setup
```

A good way of making sure everything is all right is running the test suite:
```shell script
$ make test
```

# Pull Request Process

We Use [Github Flow](https://guides.github.com/introduction/flow/index.html), So All Code Changes Happen Through Pull Requests.

Pull requests are the best way to propose changes to the codebase. We actively welcome your pull requests:

1. Be sure to open a issue about your fix/change/feature.
2. Fork the repo and create your branch from master.
3. Add tests to your code.
4. Ensure the test suite passes.
5. Make sure your code lints (`make lint`).
6. Issue that pull request!
7. Verify that all CI checks passed on your PR.

# Licence
Any contributions you make will be under the [Unlicence](https://github.com/thomaspoignant/api-scenario/blob/master/LICENSE).

In short, when you submit code changes, your submissions are understood to be under the same Unlicence that covers the project. Feel free to contact the maintainers if that's a concern.

# Report bugs using Github's [issues](https://github.com/thomaspoignant/api-scenario/issues)

We use GitHub issues to track public bugs. Report a bug by [opening a new issue](https://github.com/thomaspoignant/api-scenario/issues); it's that easy!

When creating new issue choose your template a fill as much informations as you can.

# Contributors
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://github.com/thomaspoignant"><img src="https://avatars.githubusercontent.com/u/17908063?v=3" width="100px;" alt=""/><br /><sub><b>@thomaspoignant</b></sub></a><br /><a href="https://github.com/thomaspoignant/api-scenario/commits?author=thomaspoignant" title="Code">💻</a> <a href="https://github.com/thomaspoignant/api-scenario/commits?author=thomaspoignant" title="Documentation">📖</a> <a href="https://github.com/thomaspoignant/api-scenario/pulls?q=is%3Apr+reviewed-by%3Athomaspoignant" title="Reviewed Pull Requests">👀</a></td>
    <td align="center"><a href="https://github.com/aftouh"><img src="https://avatars.githubusercontent.com/u/7847154?v=3" width="100px;" alt=""/><br /><sub><b>@aftouh</b></sub></a><br /><a href="https://github.com/thomaspoignant/api-scenario/commits?author=aftouh" title="Code">💻</a> <a href="https://github.com/thomaspoignant/api-scenario/commits?author=aftouh" title="Documentation">📖</a> <a href="https://github.com/thomaspoignant/api-scenario/pulls?q=is%3Apr+reviewed-by%3Aaftouh" title="Reviewed Pull Requests">👀</a></td>
  </tr>
</table>
<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
  

You can see the [full list of contributors](https://github.com/thomaspoignant/api-scenario/graphs/contributors).


