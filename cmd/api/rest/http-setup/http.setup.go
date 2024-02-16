package httpsetup

import (
	"encoding/json"
	"net/http"
	"strings"
)

// middleware type represent middleware function signature
type Middleware func(http.HandlerFunc) http.HandlerFunc

// router represent custom router that support multiple handlers for a single path
type Router struct {
	Mux         *http.ServeMux
	middlewares []Middleware
	prefix      string
}

// NewRouter creates a new Router instance
func NewRouter() *Router {
	return &Router{
		Mux:    http.NewServeMux(),
		prefix: "",
	}
}

// Use adds middleware to be used with all registered handlers
func (r *Router) Use(middleware ...Middleware) {
	r.middlewares = append(r.middlewares, middleware...)
}

// Handle registers the given handler with the router for specific path
func (r *Router) Handle(path string, handler http.HandlerFunc) {
	newPath := strings.Split(path, " ")[0] + " " + r.prefix + strings.Split(path, " ")[1]
	r.Mux.HandleFunc(newPath, r.applyMiddleware(handler))
}

// applyMiddleware applies all middleware to the handler
func (r *Router) applyMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	finalHandler := handler
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		finalHandler = r.middlewares[i](finalHandler)
	}
	return finalHandler
}

// Group creates a new router with a prefix
func (r *Router) Group(prefix string) *Router {
	return &Router{
		Mux:    r.Mux,
		prefix: r.prefix + prefix,
	}
}

type ApiError struct {
	Error string `json:"error"`
}

type DataResponse struct {
	TotalItems int           `json:"totalItems"`
	Items      []interface{} `json:"items"`
}

type StandardResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	res := &StandardResponse{
		Code: status,
		Data: v,
	}

	return json.NewEncoder(w).Encode(res)
}
