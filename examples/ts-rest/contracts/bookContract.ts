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
