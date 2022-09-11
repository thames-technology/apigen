# cloudstd

Generate Protobuf service definitions conforming to popular design guides:

- [Buf Style Guide](https://docs.buf.build/best-practices/style-guide)
- [Uber Protobuf Style Guide V2](https://github.com/uber/prototool/blob/dev/style/README.md)
- [Google Cloud API Design Guide](https://cloud.google.com/apis/design)

## Installation

```sh
go install github.com/slavovojacek/cloudstd@latest
```

## Usage

### Modern, simple style

This style is based on [Buf Style Guide](https://docs.buf.build/best-practices/style-guide) and [Uber Protobuf Style Guide V2](https://github.com/uber/prototool/blob/dev/style/README.md).

```sh
proto> cloudstd proto --package "acme.shelf.v1" --resource "shelf,shelves"
proto> cloudstd proto --package "acme.book.v1" --resource "book" --parent "shelf,shelves"
```

The generated code can be found [in the docs](docs/examples/proto/modern/README.md).

### Google style

This style is based on the [Google Cloud API Design Guide](https://cloud.google.com/apis/design).

```sh
proto> cloudstd proto --package "acme.shelf.v1" --resource "shelf,shelves" --google
proto> cloudstd proto --package "acme.book.v1" --resource "book" --parent "shelf,shelves" --google
```

The generated code can be found [in the docs](docs/examples/proto/google/README.md).
