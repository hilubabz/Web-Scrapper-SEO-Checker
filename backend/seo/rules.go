package seo

import "strings"

type Severity string

const (
	SeverityError   Severity = "error"
	SeverityWarning Severity = "warning"
	SeverityInfo    Severity = "info"
)

type SEOIssue struct {
	Rule     string `json:"rule"`
	Severity Severity `json:"severity"`
	Message  string `json:"message"`
}

func CheckRules(result *AnalysisResult) []SEOIssue {
	issues := make([]SEOIssue, 0)

	if result.StatusCode >= 400 {
		issues = append(issues, SEOIssue{
			Rule:     "http-error",
			Severity: SeverityError,
			Message:  "Page returned an HTTP error status.",
		})
	}

	if result.Title == "" {
		issues = append(issues, SEOIssue{
			Rule:     "missing-title",
			Severity: SeverityError,
			Message:  "Page is missing a title tag.",
		})
	} else {
		if result.TitleLength > 60 {
			issues = append(issues, SEOIssue{
				Rule:     "title-too-long",
				Severity: SeverityWarning,
				Message:  "Title is longer than 60 characters. Optimal length is 50-60 characters.",
			})
		} else if result.TitleLength < 30 {
			issues = append(issues, SEOIssue{
				Rule:     "title-too-short",
				Severity: SeverityWarning,
				Message:  "Title is shorter than 30 characters. Expand it to better describe the page.",
			})
		}
	}

	if !result.HasMetaDescription {
		issues = append(issues, SEOIssue{
			Rule:     "missing-meta-description",
			Severity: SeverityError,
			Message:  "Page is missing a meta description. This significantly impacts click-through rates.",
		})
	} else {
		if result.MetaDescriptionLen < 70 {
			issues = append(issues, SEOIssue{
				Rule:     "meta-description-too-short",
				Severity: SeverityWarning,
				Message:  "Meta description is shorter than 70 characters. Aim for 120-155 characters.",
			})
		} else if result.MetaDescriptionLen > 160 {
			issues = append(issues, SEOIssue{
				Rule:     "meta-description-too-long",
				Severity: SeverityWarning,
				Message:  "Meta description is longer than 160 characters. It may be truncated in search results.",
			})
		}
	}

	if result.H1Count == 0 {
		issues = append(issues, SEOIssue{
			Rule:     "missing-h1",
			Severity: SeverityError,
			Message:  "Page is missing an H1 heading. An H1 tag is critical for SEO structure.",
		})
	} else if result.H1Count > 1 {
		issues = append(issues, SEOIssue{
			Rule:     "multiple-h1",
			Severity: SeverityWarning,
			Message:  "Page contains multiple H1 headings. It's best practice to have exactly one H1 per page.",
		})
	}

	if result.ImageCount > 0 {
		if result.ImagesWithoutAlt == result.ImageCount {
			issues = append(issues, SEOIssue{
				Rule:     "all-images-missing-alt",
				Severity: SeverityError,
				Message:  "All images on the page are missing alt text, hurting accessibility and image SEO.",
			})
		} else if result.ImagesWithoutAlt > 0 {
			issues = append(issues, SEOIssue{
				Rule:     "missing-image-alt",
				Severity: SeverityWarning,
				Message:  "Some images are missing alt text.",
			})
		}
	}

	if !result.HasCanonical {
		issues = append(issues, SEOIssue{
			Rule:     "missing-canonical",
			Severity: SeverityWarning,
			Message:  "Page is missing a canonical URL tag, which prevents duplicate content issues.",
		})
	}

	if result.HasRobotsMeta {
		robotsLower := strings.ToLower(result.RobotsMeta)
		if strings.Contains(robotsLower, "noindex") {
			issues = append(issues, SEOIssue{
				Rule:     "noindex-found",
				Severity: SeverityError,
				Message:  "Page is explicitly set to 'noindex', meaning search engines will not index it.",
			})
		}
	}

	if !strings.HasPrefix(result.URL, "https://") {
		issues = append(issues, SEOIssue{
			Rule:     "not-https",
			Severity: SeverityError,
			Message:  "Page is not served over HTTPS. Security is a major ranking factor.",
		})
	}

	if result.ExternalLinkCount > 100 {
		issues = append(issues, SEOIssue{
			Rule:     "too-many-external-links",
			Severity: SeverityWarning,
			Message:  "Page has a very high number of external links (over 100).",
		})
	}

	if result.InternalLinkCount == 0 {
		issues = append(issues, SEOIssue{
			Rule:     "no-internal-links",
			Severity: SeverityWarning,
			Message:  "Page has no internal links, which creates an orphan page and hurts site crawlability.",
		})
	}

	return issues
}