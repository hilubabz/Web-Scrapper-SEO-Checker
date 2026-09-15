export type Severity = "error" | "warning" | "info";

export interface SEOIssue {
  rule: string;
  severity: Severity;
  message: string;
}

export interface AnalysisResult {
  url: string;
  statusCode: number;
  title: string;
  titleLength: number;
  metaDescription: string;
  hasMetaDescription: boolean;
  h1Count: number;
  h2Count: number;
  imageCount: number;
  imagesWithoutAlt: number;
  linkCount: number;
  canonicalUrl: string;
  hasCanonical: boolean;
  robotsMeta: string;
  hasRobotsMeta: boolean;
  internalLinkCount: number;
  externalLinkCount: number;
  metaDescriptionLen: number;
}

export interface CrawlResult {
  page: AnalysisResult;
  depth: number;
  links: string[];
  issues: SEOIssue[];
  score: number;
}

export interface AuditResult {
  startUrl: string;
  pagesCrawled: number;
  pages: CrawlResult[];
  overallScore: number;
  totalIssues: number;
}

export async function runAudit(
  url: string,
  maxDepth: number = 3,
  maxPages: number = 100,
): Promise<AuditResult> {
  const response = await fetch(
    `${process.env.NEXT_PUBLIC_API_URL}/api/audits`,
    {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ url, maxDepth, maxPages }),
    },
  );

  if (!response.ok) {
    const errorData = await response.json().catch(() => null);
    throw new Error(errorData?.error || "Failed to run audit");
  }

  return response.json();
}
