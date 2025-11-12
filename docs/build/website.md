# Website Build and Deployment

The Website repository provides the user-facing web application for Gitopedia. It is built with Next.js and produces a static site that can be hosted on AWS S3 (served via CloudFront). Additionally, it includes a serverless search API (an AWS Lambda function) that enables full-text search queries on the static content.

## Next.js Static Site Generation

The website uses Next.js to generate static HTML pages for all Gitopedia articles:

- We configure Next.js for Static Export by setting `output: 'export'` in `next.config.js` and using the export command to produce HTML files[8].
- At build time, the site pulls the latest Markdown content from the Gitopedia repository. This can be done by including the Gitopedia repo as a submodule or by fetching content via the GitHub API in a build script. (For example, a script could use the GitHub REST API to fetch raw files from the Gitopedia main branch.)
- For each article Markdown, a corresponding page is generated. We use Next.js dynamic routes (e.g., a page template `pages/[slug].js`) and `getStaticPaths`/`getStaticProps` to create a page for each article.
- **Markdown to HTML**: The build process will parse the Markdown (using a library like `gray-matter` to get front matter and `remark` or `marked` to convert to HTML). The front matter data (title, etc.) can be passed as props to the page for metadata (e.g., HTML `<title>` tag).
- The article content is embedded in the static HTML. This means when a user opens an article URL, the content is just regular HTML (with proper headings, paragraphs, etc., as rendered at build time).
- We also generate index pages or landing pages as needed. For example, the homepage (`index.html`) might list all article titles (this list can be sourced from the Knowledgebase `meta.json` or directly from scanning the content files during build).
- **Styling and assets**: We can use CSS or a component library to style the pages. Since it's static, we might precompute any navigation menus or sidebars. Basic navigation (home, search, list of tags) can be included on each page if desired.

After running `npm run build && npm run export`, Next.js will output static files (HTML, CSS, JS) into an `/out` directory. These files can be uploaded to an S3 bucket configured for static website hosting. Each article will correspond to an HTML file (e.g., `machine-learning-basics.html` or `machine-learning-basics/index.html` depending on Next's export structure).

### Content Updates

Because content is static, whenever Gitopedia updates, the site must be rebuilt and redeployed to include new articles or changes. We plan to automate this via CI:

- A GitHub Action in the Website repo can be triggered (by a webhook or manual dispatch) whenever the Knowledgebase signals that new content is available.
- The action will pull the latest Gitopedia content (and optionally the latest knowledgebase meta if needed), run the Next.js build, then upload the `out/` folder to S3 (for example, using the AWS CLI or an deployment tool). We'll also configure CloudFront invalidation for updated files if using a CDN.

## Search Functionality Integration

Because the site is static, we cannot perform server-side search queries. Instead, we deploy a separate Search API using AWS Lambda. This Lambda uses the `index.sqlite` produced by Knowledgebase to execute full-text searches and return results to the frontend.

### Search Lambda (Backend)

- We implement the Lambda function (in Node.js or Python). The function will load the SQLite database file either from an included asset or by downloading it from S3 on cold start.
- On receiving a search request (e.g., an HTTP GET request with a query parameter `q`), the function executes a SQL query against the SQLite FTS index to find matching articles. For example:

```sql
SELECT id, title, snippet(article_index) as snippet
FROM article_index
WHERE article_index MATCH ? 
LIMIT 10;
```

This returns a list of article IDs/titles and a snippet of context.

- The Lambda then formats this result as JSON (e.g., a list of `{id, title, snippet}` objects) and returns it in the HTTP response.

We will expose this Lambda via API Gateway or AWS Lambda Function URL:

- For simplicity, an API Gateway HTTP API can be configured with a route `/search` that triggers the Lambda[9].
- We enable CORS on this API so that our static website (which might be on a different domain or CloudFront distribution) can call it from the browser. Specifically, we allow the website's origin in the API's CORS config[5].

### Frontend (Client-side)

- The website will have a Search page (e.g., `/search` route) or at least a search bar component.
- When a user submits a query, a client-side JavaScript function will send a fetch request to the Search API endpoint (for example: `fetch('https://api.gitopedia.org/search?q=machine learning')`).
- The response (JSON with results) is then used to display search results on the page: typically a list of titles (each linking to the respective article page) with a brief snippet of the context around the query term.
- We can implement the search page as a React component that maintains state for query and results. Initially, it might show an empty state or a prompt to enter a query. After a search, it renders a list of results.
- For better UX, we might also enable searching directly from any page (e.g., a small search box in the nav bar that redirects to the search page with the query).

Because this is all client-side, the search page itself can be a static page with an empty result list that populates after the JS runs the query. It does not need server-side rendering.

### Deployment of Search API

- The Lambda function code resides in the Website repository (e.g., in a `lambda/search.py` file). We will create an AWS SAM or Serverless Framework configuration to deploy it, or use AWS CLI commands in CI.
- We have two options for providing the SQLite index to the Lambda:

  1. **Embed on Deploy**: Whenever Knowledgebase produces a new `index.sqlite`, we package that file along with the Lambda code and deploy. This ensures the Lambda always has the latest index locally (fast queries, no network call needed). We must keep the package under the size limit (~50MB zipped)[10], but SQLite indexes for text are typically quite compact.

  2. **Load from S3 at runtime**: Upload `index.sqlite` to S3 (e.g., `knowledgebase/index.sqlite`). On cold start, the Lambda downloads this file to `/tmp` storage and uses it. The Lambda can cache the DB connection across invocations to speed up subsequent queries. This method allows updating the index without redeploying code; however, it introduces a slight cold-start latency (download) and requires the Lambda to have network access to S3.

- In early phases, we might implement method (1) for simplicity: include the index in the deployment artifact each time. As the system matures, method (2) or a hybrid (embed initial index, but allow refreshing from S3 if a newer version flag is detected) could be used to decouple content updates from code deploys.

## Putting it All Together

Once both the static site and search API are deployed:

- Users accessing the Gitopedia website will get fast-loading pages (served from CloudFront/S3). They can navigate through the knowledge base as through a normal website.
- When using the search feature, the browser will call the Search API. The results come back quickly from the Lambda (which is querying the pre-built index), and the user can click on any result to navigate to that article page.
- The integration with the other components ensures that whenever content changes:
  - Knowledgebase updates the index.
  - Website CI rebuilds and deploys the new pages.
  - Search API gets the new index (via redeploy or S3 update).
  - Thus, the website and search stay in sync with the content.

By structuring the website as a static Next.js app with a supplemental search API, we achieve a balance of performance and functionality:

- The bulk of content is delivered as static files (highly scalable and cost-efficient).
- Dynamic behavior (search) is handled by a purpose-built serverless function that can scale independently and only incurs cost when used.
- This architecture follows the Jamstack approach: pre-rendered content + on-demand serverless functions for interactivity.
