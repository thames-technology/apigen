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
