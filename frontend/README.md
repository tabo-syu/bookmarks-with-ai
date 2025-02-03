# Bookmarks Frontend

A modern Next.js application for managing bookmarks with a clean, responsive interface.

## Features

- Modern React components with TypeScript
- Real-time bookmark management
- Automatic metadata display (title, description, favicon)
- Responsive design with Tailwind CSS
- React Query for efficient data fetching and caching
- Auto-generated API client from OpenAPI spec

## Prerequisites

- Node.js 18 or higher
- npm 9 or higher
- Backend service running (see ../backend)

## Setup

1. Install dependencies:
```bash
npm install
```

2. Generate API client code:
```bash
npm run generate-api
```

3. Start the development server:
```bash
npm run dev
```

The application will be available at http://localhost:3000

## Project Structure

- `src/app`: Next.js app directory
  - `page.tsx`: Main application page
  - `layout.tsx`: Root layout with providers
  - `providers.tsx`: React Query setup
  - `globals.css`: Global styles

- `src/components`: React components
  - `BookmarkList.tsx`: Displays all bookmarks
  - `CreateBookmark.tsx`: Form for adding new bookmarks

- `src/api`: Generated API client code
  - `services`: API service functions
  - `models`: TypeScript interfaces
  - `core`: API client configuration

- `scripts`: Utility scripts
  - `generate-api.js`: OpenAPI client generator

## Development

### Available Scripts

- `npm run dev`: Start development server
- `npm run build`: Build for production
- `npm run start`: Start production server
- `npm run lint`: Run ESLint
- `npm run generate-api`: Update API client code

### API Integration

The application uses a generated API client from the OpenAPI specification. When the backend API changes:

1. Copy the updated openapi.yaml to the frontend directory
2. Run `npm run generate-api`
3. TypeScript will help identify any breaking changes

### Styling

- Tailwind CSS for utility-first styling
- Custom components using @headlessui/react
- Heroicons for icons
- Responsive design patterns

### State Management

- React Query for server state
- React hooks for local state
- Optimistic updates for better UX
- Proper error handling and loading states

## Features

### Bookmark List
- Display all bookmarks with metadata
- Sort by creation date
- Delete functionality
- Error handling for failed requests
- Loading states
- Empty state handling

### Create Bookmark
- Modal form for new bookmarks
- URL validation
- Automatic metadata fetching
- Success/error feedback
- Loading state during creation

## Error Handling

- Network error handling
- Invalid URL handling
- Loading states
- Error messages
- Form validation
- Retry mechanisms

## Performance

- React Query caching
- Optimistic updates
- Lazy loading
- Proper TypeScript types
- Efficient re-renders

## Contributing

1. Ensure the backend is running
2. Make your changes
3. Run ESLint: `npm run lint`
4. Test the changes
5. Submit a pull request

## Deployment

1. Build the application:
```bash
npm run build
```

2. Start the production server:
```bash
npm run start
```

The application can be deployed to any platform that supports Next.js applications (Vercel, Netlify, etc.).
