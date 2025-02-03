'use client';

import { BookmarkList } from '@/components/BookmarkList';
import { CreateBookmark } from '@/components/CreateBookmark';

export default function Home() {
  return (
    <main className="container mx-auto px-4 py-8">
      <div className="max-w-4xl mx-auto">
        <div className="flex justify-between items-center mb-8">
          <h1 className="text-3xl font-bold text-gray-900">My Bookmarks</h1>
          <CreateBookmark />
        </div>
        <BookmarkList />
      </div>
    </main>
  );
}
