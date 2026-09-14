package crawler

type AuditResult struct {
	StartURL     string        `json:"startUrl"`
	PagesCrawled int           `json:"pagesCrawled"`
	Pages        []crawlResult `json:"pages"`
	OverallScore int           `json:"overallScore"`
	TotalIssues  int           `json:"totalIssues"`
}