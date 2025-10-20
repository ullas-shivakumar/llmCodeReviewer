package api

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/ullas-shivakumar/llmCodeReviewer/internal/core"
	"github.com/ullas-shivakumar/llmCodeReviewer/internal/models"
)

func isItGithubPR(url string) bool {
	if url == "" {
		return false
	}
	trimmedurl := strings.TrimSpace(url)

	if strings.Contains(trimmedurl, " ") {
		return false
	}

	re := regexp.MustCompile(`^https:\/\/github\.com\/([^\/]+)\/([^\/]+)\/pull\/\d+$`)
	return re.MatchString(trimmedurl)
}

func reviewhHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, models.APIResponse{
			Status:  "error",
			Message: "method not allowed",
		})
		return
	}

	var req models.ReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, models.APIResponse{
			Status:  "error",
			Message: "invalid Json in Request",
		})
		return
	}

	if req.Input == "" {
		writeJSON(w, http.StatusBadRequest, models.APIResponse{
			Status:  "error",
			Message: "Input Cannot be Empty",
		})
		return
	}

	var (
		resp *models.ReviewResponse
		err  error
	)

	if isPR := isItGithubPR(req.Input); isPR {
		resp, err = core.ReviewGithubPR(req)
	} else {
		resp, err = core.ReviewCode(req)
	}

	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.APIResponse{
			Status:  "error",
			Message: "internal service error",
		})
	}

	writeJSON(w, http.StatusOK, models.APIResponse{
		Status:  "success",
		Message: "Code Reviewed Successfully",
		Data:    resp,
	})

}

func writeJSON(w http.ResponseWriter, status int, resp models.APIResponse) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}
