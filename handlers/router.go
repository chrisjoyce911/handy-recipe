package handlers

import (
	"context"
	"handy-recipe/middleware"
	"handy-recipe/store"
	"net/http"
	"text/template"
)

// Handler holds the store.
type Handler struct {
	Store store.Store
}

// NewHandler returns a new Handler with the given store.
func NewHandler(s store.Store) *Handler {
	return &Handler{Store: s}
}

// SetupRoutes initializes the router with all the routes, using the store.
func (h *Handler) SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	// Serve static files with logging.
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("static/")))
	mux.Handle("GET /static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Inject "static" as handler name into context.
		r = r.WithContext(context.WithValue(r.Context(), middleware.HandlerNameKey, "static"))
		middleware.LoggingMiddleware(staticFileHandler).ServeHTTP(w, r)
	}))

	dataFileHandler := http.StripPrefix("/data/", http.FileServer(http.Dir("data/")))
	mux.Handle("GET /data/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Inject "static" as handler name into context.
		r = r.WithContext(context.WithValue(r.Context(), middleware.HandlerNameKey, "data"))
		middleware.LoggingMiddleware(dataFileHandler).ServeHTTP(w, r)
	}))

	// Register dynamic routes using the new routing syntax.
	// Note: the method is specified (GET) and parameters are defined in curly braces.
	mux.HandleFunc("GET /", wrapHandler("home", h.home))
	mux.HandleFunc("GET /about", wrapHandler("about", h.about))
	mux.HandleFunc("GET /tags", wrapHandler("tags", h.tags))
	mux.HandleFunc("GET /recipes", wrapHandler("recipes", h.recipes))
	mux.HandleFunc("GET /contact", wrapHandler("contact", h.contact))
	// For individual recipe pages by slug.
	mux.HandleFunc("GET /recipe/{slug}", wrapHandler("recipe", h.recipe))
	// For tag pages, listing recipes by tag.
	mux.HandleFunc("GET /tag/{tag}", wrapHandler("tag", h.tag))

	return mux
}

// wrapHandler is a helper that injects the handler name into the context and applies the logging middleware.
func wrapHandler(name string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Inject the handler name into the context.
		r = r.WithContext(context.WithValue(r.Context(), middleware.HandlerNameKey, name))
		// Wrap the handler with the logging middleware.
		middleware.LoggingMiddleware(http.HandlerFunc(handler)).ServeHTTP(w, r)
	}
}

// renderTemplate renders an HTML template with the provided data.
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
