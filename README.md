<p align="center">
  <img src="https://raw.githubusercontent.com/thames-technology/apigen/main/.github/assets/apigen-cover.png" alt="API Gen Logo" />
</p>

<p align="right">
  <i>If you use this repo, star it ✨</i>
</p>

---

<h2 align="center">Generate best-practice Protobuf APIs following design patterns</h2>

<p align="center">
  Inspired by <a href="https://www.oreilly.com/library/view/api-design-patterns/9781617295850/" target="_blank">API Design Patterns</a>
</p>

---

## Install

Homebrew:

```sh
brew install thames-technology/tap/apigen
```

Go:

```sh
go install github.com/thames-technology/apigen
```

## Getting started

Create standard `BookService` definition with `author` parent resource in `proto/bookservice/v1alpha1/service.proto`:

```sh
apigen proto -r book -p author -pkg bookservice.v1alpha1 -w
```

Create standard `AuthorService` definition without a parent resource in `proto/authorservice/v1alpha1/service.proto`:

```sh
apigen proto -r author -pkg authorservice.v1alpha1 -w
```
