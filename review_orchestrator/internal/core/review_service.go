package core

import "github.com/ullas-shivakumar/llmCodeReviewer/internal/models"

func ReviewCode(models.ReviewRequest) (*models.ReviewResponse, error) {
	// TODO: to implement the logic for reviews
	return &models.ReviewResponse{
		Cached: false,
		Result: models.ReviewResult{},
	}, nil
}

func ReviewGithubPR(models.ReviewRequest) (*models.ReviewResponse, error) {
	// TODO: to implement the logic for reviews
	return &models.ReviewResponse{
		Cached: false,
		Result: models.ReviewResult{},
	}, nil
}
