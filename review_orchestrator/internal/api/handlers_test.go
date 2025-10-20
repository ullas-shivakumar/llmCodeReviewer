package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsItGithubPR(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want bool
	}{
		// Valid cases
		{"valid basic", "https://github.com/some-company/example/pull/123", true},
		{"valid with extra spaces", "  https://github.com/org-name/repo-name/pull/1  ", true},

		// Invalid cases
		{"space in middle of URL", "  https://github.com/org-name/repo-name /pull/1", false},
		{"missing PR number", "  https://github.com/org-name/repo-name/pull/", false},
		{"wrong path (issues instead of pull)", "  https://github.com/org-name/repo-name/issues/1  ", false},
		{"http instead of https", "  http://github.com/org-name/repo-name/pull/1  ", false},
		{"wrong domain", "  https://something.com/org-name/repo-name/pull/1  ", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isItGithubPR(tt.url)
			assert.Equal(t, tt.want, got, "URL: %q", tt.url)
		})
	}
}
