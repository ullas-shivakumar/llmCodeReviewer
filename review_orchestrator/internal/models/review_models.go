package models

type ReviewRequest struct {
	Input    string
	Language string
}

type Issue struct {
	Type    string
	Message string
}

type Suggestion struct {
	Type    string
	Message string
}

type ReviewResult struct {
	summary     string
	Language    string
	issues      []Issue
	suggestions []Suggestion
	confidence  float64
}
type ReviewResponse struct {
	Cached bool
	Result ReviewResult
}
