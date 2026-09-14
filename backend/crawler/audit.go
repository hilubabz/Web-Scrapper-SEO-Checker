package crawler

func BuildAuditResult(
	startURL string,
	pages []crawlResult,
) AuditResult {
	totalScore := 0
	totalIssues := 0

	for _, page := range pages {
		totalScore += page.Score
		totalIssues += len(page.Issues)
	}

	overallScore := 0

	if len(pages) > 0 {
		overallScore = totalScore / len(pages)
	}

	return AuditResult{
		StartURL:     startURL,
		PagesCrawled: len(pages),
		Pages:        pages,
		OverallScore: overallScore,
		TotalIssues:  totalIssues,
	}
}