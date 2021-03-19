package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Handler - stores pointer to our comments service
type Handler struct{
	Router *mux.Router
}

// NewHandler - returs a pointer to a handler
func NewHandler() *Handler {
	return &Handler{}
}
// Setup Routes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting Up routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,  "I am alive!")
	})
}
