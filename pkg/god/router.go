package god

import (
	"log/slog"
	"net/http"
	"os"
	"sync"
)

type Router struct {
	mu     sync.RWMutex
	routes map[string]map[string][]HandlerFunc
	// middleware []HandlerFunc
	log *slog.Logger
}

// Default creates a new default Router instance.
func Default() *Router {
	return &Router{
		routes: make(map[string]map[string][]HandlerFunc),
		log:    slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
	}
}

// TODO: Use => adds middleware to the router.

// Handle registers a new route with a method and path.
func (r *Router) Handle(method, path string, handlers ...HandlerFunc) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.routes[method][path]; !ok {
		r.routes[method] = make(map[string][]HandlerFunc)
	}

	r.routes[method][path] = append(r.routes[method][path], handlers...)
}

// ServeHTTP implements the http.Handler interface.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path

	if handlers, ok := r.routes[method][path]; ok {
		c := NewContext(w, req)

		// TODO: Implement Params parsing and storing

		c.fullPath = path
		c.handlers = handlers
		c.Next()
	} else {
		http.NotFound(w, req)
	}
}

// GET registers a GET route.
func (r *Router) GET(path string, handler HandlerFunc) {
	r.Handle(http.MethodGet, path, handler)
}

// POST registers a POST route.
func (r *Router) POST(path string, handler HandlerFunc) {
	r.Handle(http.MethodPost, path, handler)
}

// PUT registers a PUT route.
func (r *Router) PUT(path string, handler HandlerFunc) {
	r.Handle(http.MethodPut, path, handler)
}

// DELETE registers a DELETE route.
func (r *Router) DELETE(path string, handler HandlerFunc) {
	r.Handle(http.MethodDelete, path, handler)
}

// TODO: LoggerMiddleware

// Run starts the HTTP server.
func (r *Router) Run(addr string) error {
	for method, rs := range r.routes {
		for path := range rs {
			r.log.Info("Registered route", slog.String("method", method), slog.String("path", path))
		}
	}
	return http.ListenAndServe(addr, r)
}
