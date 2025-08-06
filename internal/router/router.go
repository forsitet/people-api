package router

import (
	"log"
	"net/http"
	"os"

	peopleV1 "people/shared/pkg/openapi/people/v1"
)

type Router struct {
	server *peopleV1.Server
}

func NewRouter(server *peopleV1.Server) *Router {
	return &Router{
		server: server,
	}
}

func (r *Router) SetupRoutes() http.Handler {
	mux := http.NewServeMux()
	r.setupAPIRoutes(mux)
	return mux
}

func (r *Router) setupAPIRoutes(mux *http.ServeMux) {
	mux.Handle("/api/v1/", r.server)

	mux.HandleFunc("/openapi.yaml", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "people.openapi.v1.bundle.yaml")
	})

	r.setupSwaggerRoutes(mux)

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/" {
			http.Redirect(w, req, "/swagger", http.StatusSeeOther)
			return
		}
		http.NotFound(w, req)
	})
}

func (r *Router) setupSwaggerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/swagger/", func(w http.ResponseWriter, req *http.Request) {
		content, err := os.ReadFile("swagger.html")
		if err != nil {
			http.Error(w, "Failed to read swagger.html", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		if _, err := w.Write(content); err != nil {
			log.Printf("Error writing response: %v", err)
		}
	})
}
