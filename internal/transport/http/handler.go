package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TutorialEdge/go-rest-api-course/internal/comment"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

// Handler - stores pointer to our comments service
type Handler struct{
	Router *mux.Router
	Service *comment.Service
}
// Response - an object to store responses from our API
type Response struct {
	Message string
	Error string
}
// NewHandler - returs a pointer to a handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}
// LoggingMiddleWARE - adds middleware around endpoints
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(
			log.Fields{
				"Method": r.Method,
				"Path": r.URL.Path,
		}).Info("Handled request")
		next.ServeHTTP(w, r)
	})
}
// BasicAuth -a handy middleware function that will provide basic auth around specific endpoints
func BasicAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("Basic auth endpoint hit")
		user, pass, ok := r.BasicAuth()
		if user == "admin" && pass == "password" && ok {
		   original(w,r)
		} else {

			SendErrorResponse(w, "not authorized", errors.New("not authorized"))
			return
		}
		original(w,r)
	}
}
func validateToken(accessToken string) bool {
	var mySigningKey = []byte("missionimpossible")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there has been an error")
		}
		return mySigningKey, nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}

// JWTAuth - a decorator function for jwt validation for endpoints
func JWAuth(original func(w http.ResponseWriter, r *http.Request))  func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("jwt authentification hit")
		authHeader := r.Header["Authorization"]
		if authHeader == nil {

			SendErrorResponse(w, "not authorized", errors.New("not authorized"))
			return
		}
		// Bearer jwt-token
		authHeadersParts := strings.Split(authHeader[0], " ")
		if len(authHeadersParts) != 2 || strings.ToLower(authHeadersParts[0]) != "bearer" {

			SendErrorResponse(w, "not authorized", errors.New("not authorized"))
			return
		}
		if  validateToken(authHeadersParts[1]) {
			original(w, r)
		} else {

			SendErrorResponse(w, "not authorized", errors.New("not authorized"))
			return
		}
	}
}

// Setup Routes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	log.Info("Setting Up routes")
	h.Router = mux.NewRouter()
	h.Router.Use(LoggingMiddleware)
	h.Router.HandleFunc("/api/comments", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment", JWAuth(h.PostComment)).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", JWAuth(h.UpdateComment)).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", JWAuth(h.DeleteComment)).Methods("DELETE")
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "I am alive"}); err != nil {
			panic(err)

		}
	})
}

func SendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		log.Error(err)
	}
}