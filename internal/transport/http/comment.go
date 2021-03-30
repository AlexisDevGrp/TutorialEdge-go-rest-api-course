package http

import (
	"encoding/json"
	"fmt"
	"github.com/TutorialEdge/go-rest-api-course/internal/comment"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// GetComment - retrieves a comment by ID
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		SendErrorResponse(w, "Error to parse the ID of the comment to numeric", err)
		return
	}
	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		SendErrorResponse(w, "Error Retrieving the ID of the comment", err)
		return
	}
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}

}
// GetAllComments - retrieves all comments
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	comments, err := h.Service.GetAllComments()
	if err != nil {
		SendErrorResponse(w, "Error Retrieving all the comments", err)
		return
	}
	if err := json.NewEncoder(w).Encode(comments); err != nil {
		panic(err)
	}

}
// DeleteComment - Delete a comment by ID
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		SendErrorResponse(w, "Error to parse the ID of the comment to numeric", err)
		return
	}
	err = h.Service.DeleteComment(uint(commentID))
	if err != nil {
		SendErrorResponse(w, "Error deleting the comment via ID", err)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Comment successfully deleted"}); err != nil {
		panic(err)

	}
}
// PostComment - adds a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		SendErrorResponse(w, "Failed to decode new Comment via JSON Body", err)
		return
	}
	comment, err := h.Service.PostComment(comment)
	if err != nil {
		fmt.Fprintf(w, "%+v", comment)
	}
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}
// UpdateComment - updates a comment by ID
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		SendErrorResponse(w, "Failed to decode new Comment via JSON Body", err)
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		SendErrorResponse(w, "Error to parse the ID of the comment to numeric", err)
		return
	}
	comment, err = h.Service.UpdateComment(uint(commentID), comment)
	if err != nil {
		SendErrorResponse(w, "Error updating a comment", err)
		return
	}
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}


