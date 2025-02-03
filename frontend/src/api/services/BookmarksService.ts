/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { BookmarkResponse } from '../models/BookmarkResponse';
import type { BookmarksResponse } from '../models/BookmarksResponse';
import type { CreateBookmarkRequest } from '../models/CreateBookmarkRequest';
import type { DeleteResponse } from '../models/DeleteResponse';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class BookmarksService {
    /**
     * Create a new bookmark
     * Creates a new bookmark and automatically fetches metadata from the URL
     * @param requestBody
     * @returns BookmarkResponse Bookmark created successfully
     * @throws ApiError
     */
    public static createBookmark(
        requestBody: CreateBookmarkRequest,
    ): CancelablePromise<BookmarkResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/bookmarks',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Invalid request body or URL`,
                500: `Internal server error`,
            },
        });
    }
    /**
     * List all bookmarks
     * Retrieves a list of all bookmarks
     * @returns BookmarksResponse List of bookmarks retrieved successfully
     * @throws ApiError
     */
    public static listBookmarks(): CancelablePromise<BookmarksResponse> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/bookmarks',
            errors: {
                500: `Internal server error`,
            },
        });
    }
    /**
     * Get a specific bookmark
     * Retrieves a specific bookmark by ID
     * @param id ID of the bookmark
     * @returns BookmarkResponse Bookmark retrieved successfully
     * @throws ApiError
     */
    public static getBookmark(
        id: number,
    ): CancelablePromise<BookmarkResponse> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/bookmarks/{id}',
            path: {
                'id': id,
            },
            errors: {
                404: `Bookmark not found`,
                500: `Internal server error`,
            },
        });
    }
    /**
     * Delete a bookmark
     * Deletes a specific bookmark by ID
     * @param id ID of the bookmark
     * @returns DeleteResponse Bookmark deleted successfully
     * @throws ApiError
     */
    public static deleteBookmark(
        id: number,
    ): CancelablePromise<DeleteResponse> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/bookmarks/{id}',
            path: {
                'id': id,
            },
            errors: {
                404: `Bookmark not found`,
                500: `Internal server error`,
            },
        });
    }
}
