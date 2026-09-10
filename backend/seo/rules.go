package seo

import "strings"

type Severity string

const (
	SeverityError   Severity = "error"
	SeverityWarning Severity = "warning"
	SeverityInfo    Severity = "info"
)

type SEOIssue struct {
	Rule     string
	Severity Severity
	Message  string
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
			Message:  "Page is missing a title.",
		})
	}

	if result.TitleLength > 60 {
		issues = append(issues, SEOIssue{
			Rule:     "title-too-long",
			Severity: SeverityWarning,
			Message:  "Title is longer than 60 characters.",
		})
	}

	if result.TitleLength > 0 && result.TitleLength < 30 {
		issues = append(issues, SEOIssue{
			Rule:     "title-too-short",
			Severity: SeverityWarning,
			Message:  "Title is shorter than 30 characters.",
		})
	}

	if !result.HasMetaDescription {
		issues = append(issues, SEOIssue{
			Rule:     "missing-meta-description",
			Severity: SeverityWarning,
			Message:  "Page is missing a meta description.",
		})
	}

	if result.H1Count == 0 {
		issues = append(issues, SEOIssue{
			Rule:     "missing-h1",
			Severity: SeverityError,
			Message:  "Page is missing an H1 heading.",
		})
	}

	if result.H1Count > 1 {
		issues = append(issues, SEOIssue{
			Rule:     "multiple-h1",
			Severity: SeverityWarning,
			Message:  "Page contains multiple H1 headings.",
		})
	}

	if result.ImagesWithoutAlt > 0 {
		issues = append(issues, SEOIssue{
			Rule:     "missing-image-alt",
			Severity: SeverityWarning,
			Message:  "Some images are missing alt text.",
		})
	}

	if !result.HasCanonical {
		issues = append(issues, SEOIssue{
			Rule:     "missing-canonical",
			Severity: SeverityWarning,
			Message:  "Page is missing a canonical URL.",
		})
	}

	if result.MetaDescriptionLen > 0 &&
		result.MetaDescriptionLen < 70 {
		issues = append(issues, SEOIssue{
			Rule:     "meta-description-too-short",
			Severity: SeverityWarning,
			Message:  "Meta description is shorter than 70 characters.",
		})
	}

	if result.MetaDescriptionLen > 160 {
		issues = append(issues, SEOIssue{
			Rule:     "meta-description-too-long",
			Severity: SeverityWarning,
			Message:  "Meta description is longer than 160 characters.",
		})
	}

	if !strings.HasPrefix(result.URL.String(), "https://") {
		issues = append(issues, SEOIssue{
			Rule:     "not-https",
			Severity: SeverityError,
			Message:  "Page is not using HTTPS.",
		})
	}

	return issues
}