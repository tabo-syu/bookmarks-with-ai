'use client';

import type { Bookmark } from '@/api/models/Bookmark';
import { BookmarksService } from '@/api/services/BookmarksService';
import { TrashIcon } from '@heroicons/react/24/outline';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';

export function BookmarkList() {
  const queryClient = useQueryClient();

  const { data: bookmarksResponse, isLoading, error } = useQuery({
    queryKey: ['bookmarks'],
    queryFn: () => BookmarksService.listBookmarks(),
  });

  const deleteMutation = useMutation({
    mutationFn: (id: number) => BookmarksService.deleteBookmark(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['bookmarks'] });
    },
  });

  if (isLoading) return <div className="text-center">Loading bookmarks...</div>;
  if (error) return <div className="text-red-500">Error loading bookmarks</div>;
  if (!bookmarksResponse?.bookmarks?.length) return <div className="text-center">No bookmarks yet</div>;

  return (
    <div className="space-y-4">
      {bookmarksResponse.bookmarks.map((bookmark: Bookmark) => (
        <div
          key={bookmark.id}
          className="bg-white shadow rounded-lg p-4 flex items-start justify-between"
        >
          <div className="flex-1">
            <div className="flex items-center gap-2">
              {bookmark.favicon_url && (
                <img
                  src={bookmark.favicon_url}
                  alt=""
                  className="w-4 h-4"
                  onError={(e) => {
                    e.currentTarget.style.display = 'none';
                  }}
                />
              )}
              <a
                href={bookmark.url}
                target="_blank"
                rel="noopener noreferrer"
                className="text-lg font-medium text-blue-600 hover:text-blue-800"
              >
                {bookmark.title || bookmark.url}
              </a>
            </div>
            {bookmark.description && (
              <p className="mt-1 text-gray-600">{bookmark.description}</p>
            )}
            <div className="mt-2 text-sm text-gray-500">
              Added on {bookmark.created_at ? new Date(bookmark.created_at).toLocaleDateString() : 'Unknown date'}
            </div>
          </div>
          <button
            onClick={() => bookmark.id && deleteMutation.mutate(bookmark.id)}
            className="ml-4 p-2 text-gray-400 hover:text-red-500 rounded-full hover:bg-gray-100"
            title="Delete bookmark"
          >
            <TrashIcon className="w-5 h-5" />
          </button>
        </div>
      ))}
    </div>
  );
}