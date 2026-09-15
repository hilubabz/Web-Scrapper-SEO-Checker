# Frontend - SEO Web Analyzer

The frontend application provides a sleek, interactive dashboard for users to input URLs, trigger SEO audits, and review the results in a highly readable format. It is built using **Next.js (React, TypeScript)**.

## Architecture & Implementation

The frontend is structured using the Next.js App Router:

- **`app/`**: Contains the main routing and global layouts.
  - `page.tsx`: The main landing page and search form. It manages the state for the audit (loading, results, errors) and passes data to the presentation components.
  - `globals.css`: Contains the design system tokens, CSS variables for the glassmorphic UI, and custom utility classes.
- **`components/`**: Reusable React UI components.
  - `Dashboard.tsx`: Displays high-level metrics like the overall score and total pages crawled.
  - `PageDetails.tsx`: Renders detailed metrics for an individual page.
  - `IssuesList.tsx`: Renders the specific SEO issues found on a page, categorized by severity.
- **`lib/api.ts`**: The API client that handles the `POST` request to the backend `/api/audits` endpoint.

## Design & Interactivity

The UI focuses on a premium user experience:
- **Animations**: Uses `framer-motion` for smooth entrance animations, list staggering, and subtle hover interactions on cards.
- **Icons**: Uses `lucide-react` to visually enhance the data presentation and severity badges.
- **Styling**: Relies on a modern dark-mode aesthetic with gradients, blur filters (glassmorphism), and distinct status colors for SEO scores.

## Running the Frontend

Ensure you have Node.js and `pnpm` installed.

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```
2. Install dependencies (if you haven't already):
   ```bash
   pnpm install
   ```
3. Start the development server:
   ```bash
   pnpm dev
   ```

The application will be available at `http://localhost:3000`. You can configure the backend API URL by setting `NEXT_PUBLIC_API_URL` in your `.env.local` file.
