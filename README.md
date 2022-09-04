# cloudstd

Generate Protobuf service definitions conforming to popular style guides:

- [Buf Style Guide](https://docs.buf.build/best-practices/style-guide)
- [Uber Protobuf Style Guide V2](https://github.com/uber/prototool/blob/dev/style/README.md)
- [Google Cloud API Design Guide](https://cloud.google.com/apis/design)

## Installation

```sh
go install github.com/slavovojacek/cloudstd@latest
```

## Usage

### Default (modern, simple) styleguide

This style is based on [Buf Style Guide](https://docs.buf.build/best-practices/style-guide) and [Uber Protobuf Style Guide V2](https://github.com/uber/prototool/blob/dev/style/README.md).

```sh
proto> cloudstd proto --package "acme.shelf.v1" --resource "shelf,shelves"
proto> cloudstd proto --package "acme.book.v1" --resource "book" --parent "shelf,shelves"
```

See [docs](docs/examples/proto/modern/README.md) for working example.

Sample `buf.yaml` configuration:

```yaml
version: v1
lint:
  use:
    - DEFAULT
    - COMMENTS
    - UNARY_RPC
    - PACKAGE_NO_IMPORT_CYCLE
```

### Google styleguide

This style is based on the [Google Cloud API Design Guide](https://cloud.google.com/apis/design).

```sh
proto> cloudstd proto --package "acme.shelf.v1" --resource "shelf,shelves" --google
proto> cloudstd proto --package "acme.book.v1" --resource "book" --parent "shelf,shelves" --google
```

See [docs](docs/examples/proto/google/README.md) for working example.

Sample `buf.yaml` configuration:

```yaml
version: v1
lint:
  use:
    - DEFAULT
    - COMMENTS
    - UNARY_RPC
    - PACKAGE_NO_IMPORT_CYCLE
  except:
    # Set to comply with the Google Cloud API style guide
    - RPC_REQUEST_RESPONSE_UNIQUE
    - RPC_RESPONSE_STANDARD_NAME
```
