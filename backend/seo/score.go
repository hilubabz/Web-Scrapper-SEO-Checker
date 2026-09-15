package seo

func CalculateScore(issues []SEOIssue) int {
	score := 100

	for _, issue := range issues {
		switch issue.Severity {
		case SeverityError:
			score -= 20
		case SeverityWarning:
			score -= 8
		case SeverityInfo:
			score -= 2
		}
	}

	if score < 0 {
		score = 0
	}

	return score
}