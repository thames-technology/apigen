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

### Protobuf

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
  // UUID used to prevent duplicate requests.
  optional string request_id = 1;
  // The book to be created.
  Book book = 2;
}

// Response message for CreateBook method.
message CreateBookResponse {
  // The created book.
  Book book = 1;
}

// Request message for GetBook method.
message GetBookRequest {
  // ID of the book to retrieve.
  string book_id = 1;
}

// Response message for GetBook method.
message GetBookResponse {
  // The retrieved book.
  Book book = 1;
}

// Request message for ListBooks method.
message ListBooksRequest {
  // The maximum number of books to return.
  optional int32 page_size = 1;
  // The token to retrieve the next page of results.
  optional string page_token = 2;
  // ID of the author to list books for.
  optional string author_id = 3;
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
  // UUID used to prevent duplicate requests.
  optional string request_id = 1;
  // ID of the book to update.
  string book_id = 2;
  // The book to be updated.
  Book book = 3;
  // The field mask specifying the fields to update.
  google.protobuf.FieldMask update_mask = 4;
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
  // Additional fields for the book.
  // TODO: Add fields here.
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
  // UUID used to prevent duplicate requests.
  optional string request_id = 1;
  // The author to be created.
  Author author = 2;
}

// Response message for CreateAuthor method.
message CreateAuthorResponse {
  // The created author.
  Author author = 1;
}

// Request message for GetAuthor method.
message GetAuthorRequest {
  // ID of the author to retrieve.
  string author_id = 1;
}

// Response message for GetAuthor method.
message GetAuthorResponse {
  // The retrieved author.
  Author author = 1;
}

// Request message for ListAuthors method.
message ListAuthorsRequest {
  // The maximum number of authors to return.
  optional int32 page_size = 1;
  // The token to retrieve the next page of results.
  optional string page_token = 2;
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
  // UUID used to prevent duplicate requests.
  optional string request_id = 1;
  // ID of the author to update.
  string author_id = 2;
  // The author to be updated.
  Author author = 3;
  // The field mask specifying the fields to update.
  google.protobuf.FieldMask update_mask = 4;
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
  // Additional fields for the author.
  // TODO: Add fields here.
}
```

### TS-REST

Create `book` contract with `author` parent resource:

```sh
apigen ts-rest -r book -p author -w
```

This will generate the following standard API definitions in `contracts/bookContract.ts`:

```ts
import { initContract } from '@ts-rest/core';
import { z } from 'zod';

// ====================================================================================================================
// Schemas
// ====================================================================================================================

// Timestamps using ISO 8601 format
const Timestamp = z.string().datetime().describe('ISO 8601 format timestamp');

// Book schema
export const Book = z.object({
  id: z.string().ulid().describe('Unique identifier for the book'),
  createTime: Timestamp.describe('When the book was created'),
  updateTime: Timestamp.describe('When the book was last updated'),
  authorId: z.string().ulid().describe('ID of the author of the book'),
});

// ====================================================================================================================
// Request and response schemas
// ====================================================================================================================

export const CreateBookRequest = z.object({
  requestId: z.string().uuid().optional().describe('UUID used to prevent duplicate requests'),
  book: Book.omit({ id: true, createTime: true, updateTime: true }),
});

export const CreateBookResponse = z.object({
  book: Book,
});

export const GetBookRequest = z.object({
  bookId: z.string().ulid().describe('The unique identifier of the book'),
});

export const GetBookResponse = z.object({
  book: Book,
});

export const ListBooksRequest = z.object({
  pageSize: z.number().int().min(1).max(1000).default(100).describe('The maximum number of books to return'),
  pageToken: z.string().optional().describe('The token to retrieve the next page of results'),
  authorId: z.string().ulid().optional().describe('The unique identifier of the parent author resource'),
});

export const ListBooksResponse = z.object({
  results: z.array(Book).describe('The list of books'),
  nextPageToken: z.string().ulid().nullish().describe('The token to retrieve the next page of items'),
  totalResultsEstimate: z.number().int().min(0).describe('The estimated total number of results'),
});

export const UpdateBookRequest = z.object({
  requestId: z.string().uuid().optional().describe('UUID used to prevent duplicate requests'),
  bookId: z.string().ulid().describe('The unique identifier of the book to update'),
  book: Book.omit({ id: true, createTime: true, updateTime: true }).partial(),
});

export const UpdateBookResponse = z.object({
  book: Book,
});

export const DeleteBookRequest = z.object({
  bookId: z.string().ulid().describe('The unique identifier of the book to delete'),
});

export const DeleteBookResponse = z.null();

// ====================================================================================================================
// Contracts
// ====================================================================================================================

const c = initContract();

export const BookServiceContract = c.router({
  createBook: {
    method: 'POST',
    path: '/books',
    body: CreateBookRequest,
    responses: {
      201: CreateBookResponse,
    },
    summary: 'Create a book',
  },
  getBook: {
    method: 'GET',
    path: '/books/:bookId',
    pathParams: GetBookRequest,
    responses: {
      200: GetBookResponse,
    },
    summary: 'Get a book by ID',
  },
  listBooks: {
    method: 'GET',
    path: '/books',
    query: ListBooksRequest,
    responses: {
      200: ListBooksResponse,
    },
    summary: 'List all books',
  },
  updateBook: {
    method: 'PATCH',
    path: '/books/:bookId',
    pathParams: UpdateBookRequest.pick({ bookId: true }),
    body: UpdateBookRequest.omit({ bookId: true }),
    responses: {
      200: UpdateBookResponse,
    },
    summary: 'Update a book',
  },
  deleteBook: {
    method: 'DELETE',
    path: '/books/:bookId',
    pathParams: DeleteBookRequest,
    body: z.null(), // No body, but we need to specify it as null
    responses: {
      204: DeleteBookResponse,
    },
    summary: 'Delete a book',
  },
});
```

Create `author` contract withouth a parent resource and using uuid as id:

```sh
apigen ts-rest -r author --id uuid -w
```

This will generate the following standard API definitions in `contracts/authorContract.ts`:

```ts
import { initContract } from '@ts-rest/core';
import { z } from 'zod';

// ====================================================================================================================
// Schemas
// ====================================================================================================================

// Timestamps using ISO 8601 format
const Timestamp = z.string().datetime().describe('ISO 8601 format timestamp');

// Author schema
export const Author = z.object({
  id: z.string().uuid().describe('Unique identifier for the author'),
  createTime: Timestamp.describe('When the author was created'),
  updateTime: Timestamp.describe('When the author was last updated'),
});

// ====================================================================================================================
// Request and response schemas
// ====================================================================================================================

export const CreateAuthorRequest = z.object({
  requestId: z.string().uuid().optional().describe('UUID used to prevent duplicate requests'),
  author: Author.omit({ id: true, createTime: true, updateTime: true }),
});

export const CreateAuthorResponse = z.object({
  author: Author,
});

export const GetAuthorRequest = z.object({
  authorId: z.string().uuid().describe('The unique identifier of the author'),
});

export const GetAuthorResponse = z.object({
  author: Author,
});

export const ListAuthorsRequest = z.object({
  pageSize: z.number().int().min(1).max(1000).default(100).describe('The maximum number of authors to return'),
  pageToken: z.string().optional().describe('The token to retrieve the next page of results'),
});

export const ListAuthorsResponse = z.object({
  results: z.array(Author).describe('The list of authors'),
  nextPageToken: z.string().uuid().nullish().describe('The token to retrieve the next page of items'),
  totalResultsEstimate: z.number().int().min(0).describe('The estimated total number of results'),
});

export const UpdateAuthorRequest = z.object({
  requestId: z.string().uuid().optional().describe('UUID used to prevent duplicate requests'),
  authorId: z.string().uuid().describe('The unique identifier of the author to update'),
  author: Author.omit({ id: true, createTime: true, updateTime: true }).partial(),
});

export const UpdateAuthorResponse = z.object({
  author: Author,
});

export const DeleteAuthorRequest = z.object({
  authorId: z.string().uuid().describe('The unique identifier of the author to delete'),
});

export const DeleteAuthorResponse = z.null();

// ====================================================================================================================
// Contracts
// ====================================================================================================================

const c = initContract();

export const AuthorServiceContract = c.router({
  createAuthor: {
    method: 'POST',
    path: '/authors',
    body: CreateAuthorRequest,
    responses: {
      201: CreateAuthorResponse,
    },
    summary: 'Create a author',
  },
  getAuthor: {
    method: 'GET',
    path: '/authors/:authorId',
    pathParams: GetAuthorRequest,
    responses: {
      200: GetAuthorResponse,
    },
    summary: 'Get a author by ID',
  },
  listAuthors: {
    method: 'GET',
    path: '/authors',
    query: ListAuthorsRequest,
    responses: {
      200: ListAuthorsResponse,
    },
    summary: 'List all authors',
  },
  updateAuthor: {
    method: 'PATCH',
    path: '/authors/:authorId',
    pathParams: UpdateAuthorRequest.pick({ authorId: true }),
    body: UpdateAuthorRequest.omit({ authorId: true }),
    responses: {
      200: UpdateAuthorResponse,
    },
    summary: 'Update a author',
  },
  deleteAuthor: {
    method: 'DELETE',
    path: '/authors/:authorId',
    pathParams: DeleteAuthorRequest,
    body: z.null(), // No body, but we need to specify it as null
    responses: {
      204: DeleteAuthorResponse,
    },
    summary: 'Delete a author',
  },
});
```
