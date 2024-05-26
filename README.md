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

Create `BookService` with `author` parent resource:

```sh
apigen proto -r book -p author -pkg bookservice.v1alpha1 -w
```

This will generate the following standard API definitions in `proto/bookservice/v1alpha1/service.proto`:

```proto
syntax = "proto3";

package bookservice.v1alpha1;

import "google/protobuf/empty.proto";
import "google/protobuf/field_mask.proto";
import "google/protobuf/timestamp.proto";

// =============================================================================
// Service definition
// =============================================================================

// The BookService defines the standard methods for managing books.
service BookService {
  // Creates a new book in the library.
  rpc CreateBook(CreateBookRequest) returns (CreateBookResponse);
  // Retrieves a book by its ID.
  rpc GetBook(GetBookRequest) returns (GetBookResponse);
  // Lists books with pagination support.
  rpc ListBooks(ListBooksRequest) returns (ListBooksResponse);
  // Updates an existing book.
  rpc UpdateBook(UpdateBookRequest) returns (UpdateBookResponse);
  // Deletes a book by its ID.
  rpc DeleteBook(DeleteBookRequest) returns (google.protobuf.Empty);
}

// =============================================================================
// Service request and response messages
// =============================================================================

// Request message for CreateBook method.
message CreateBookRequest {
  // The book to be created.
  Book book = 1;
}

// Response message for CreateBook method.
message CreateBookResponse {
  // The created book.
  Book book = 1;
}

// Request message for GetBook method.
message GetBookRequest {
  // ID of the book to retrieve.
  string id = 1;
}

// Response message for GetBook method.
message GetBookResponse {
  // The retrieved book.
  Book book = 1;
}

// Request message for ListBooks method.
message ListBooksRequest {
  // The maximum number of books to return.
  int32 page_size = 1;
  // The token to retrieve the next page of results.
  string page_token = 2;
  // The ID of the parent resource.
  string parent_id = 3;
}

// Response message for ListBooks method.
message ListBooksResponse {
  // The list of books.
  repeated Book results = 1;
  // Token to retrieve the next page of results.
  string next_page_token = 2;
  // Estimated total number of results.
  int32 total_results_estimate = 3;
}

// Request message for UpdateBook method.
message UpdateBookRequest {
  // The book to be updated.
  Book book = 1;
  // The field mask specifying the fields to update.
  google.protobuf.FieldMask update_mask = 2;
}

// Response message for UpdateBook method.
message UpdateBookResponse {
  // The updated book.
  Book book = 1;
}

// Request message for DeleteBook method.
message DeleteBookRequest {
  // ID of the book to delete.
  string id = 1;
}

// =============================================================================
// Resource messages
// =============================================================================

// The Book message represents a book in the library.
message Book {
  // Unique identifier for the book.
  string id = 1;
  // When the book was created.
  google.protobuf.Timestamp create_time = 2;
  // When the book was last updated.
  google.protobuf.Timestamp update_time = 3;
  // ID of the author of the book.
  string author_id = 4;
}
```

Create `AuthorService` definition without a parent resource:

```sh
apigen proto -r author -pkg authorservice.v1alpha1 -w
```

This will generate the following standard API definitions in `proto/authorservice/v1alpha1/service.proto`:

```proto
syntax = "proto3";

package authorservice.v1alpha1;

import "google/protobuf/empty.proto";
import "google/protobuf/field_mask.proto";
import "google/protobuf/timestamp.proto";

// =============================================================================
// Service definition
// =============================================================================

// The AuthorService defines the standard methods for managing authors.
service AuthorService {
  // Creates a new author in the library.
  rpc CreateAuthor(CreateAuthorRequest) returns (CreateAuthorResponse);
  // Retrieves a author by its ID.
  rpc GetAuthor(GetAuthorRequest) returns (GetAuthorResponse);
  // Lists authors with pagination support.
  rpc ListAuthors(ListAuthorsRequest) returns (ListAuthorsResponse);
  // Updates an existing author.
  rpc UpdateAuthor(UpdateAuthorRequest) returns (UpdateAuthorResponse);
  // Deletes a author by its ID.
  rpc DeleteAuthor(DeleteAuthorRequest) returns (google.protobuf.Empty);
}

// =============================================================================
// Service request and response messages
// =============================================================================

// Request message for CreateAuthor method.
message CreateAuthorRequest {
  // The author to be created.
  Author author = 1;
}

// Response message for CreateAuthor method.
message CreateAuthorResponse {
  // The created author.
  Author author = 1;
}

// Request message for GetAuthor method.
message GetAuthorRequest {
  // ID of the author to retrieve.
  string id = 1;
}

// Response message for GetAuthor method.
message GetAuthorResponse {
  // The retrieved author.
  Author author = 1;
}

// Request message for ListAuthors method.
message ListAuthorsRequest {
  // The maximum number of authors to return.
  int32 page_size = 1;
  // The token to retrieve the next page of results.
  string page_token = 2;
  // The ID of the parent resource.
  string parent_id = 3;
}

// Response message for ListAuthors method.
message ListAuthorsResponse {
  // The list of authors.
  repeated Author results = 1;
  // Token to retrieve the next page of results.
  string next_page_token = 2;
  // Estimated total number of results.
  int32 total_results_estimate = 3;
}

// Request message for UpdateAuthor method.
message UpdateAuthorRequest {
  // The author to be updated.
  Author author = 1;
  // The field mask specifying the fields to update.
  google.protobuf.FieldMask update_mask = 2;
}

// Response message for UpdateAuthor method.
message UpdateAuthorResponse {
  // The updated author.
  Author author = 1;
}

// Request message for DeleteAuthor method.
message DeleteAuthorRequest {
  // ID of the author to delete.
  string id = 1;
}

// =============================================================================
// Resource messages
// =============================================================================

// The Author message represents a author in the library.
message Author {
  // Unique identifier for the author.
  string id = 1;
  // When the author was created.
  google.protobuf.Timestamp create_time = 2;
  // When the author was last updated.
  google.protobuf.Timestamp update_time = 3;
}
```
