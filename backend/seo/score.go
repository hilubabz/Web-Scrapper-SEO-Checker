package seo

func CalculateScore(issues []SEOIssue) int {
	score := 100

	for _, issue := range issues {
		switch issue.Severity {
		case SeverityError:
			score -= 15

		case SeverityWarning:
			score -= 5

		case SeverityInfo:
			score -= 1
		}
	}

	if score < 0 {
		score = 0
	}

	return score
}