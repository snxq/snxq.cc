package server

import (
	"log"
	"net/http"
	"strings"
)

// Server .
type Server struct {
	dao Dao
}

// New return server with dao
func New(d Dao) *Server {
	return &Server{d}
}

// ServeHTTP impl http.Handler
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("path: %s", r.URL.Path)

	path := r.URL.Path
	if path == "/" {
		s.homeHandler(w, r)
	} else if strings.HasPrefix(path, "/articles/") {
		s.articleHandler(w, r)
	} else if path == "/notfound" {
		notFoundHandler(w, r)
	} else {
		notFoundHandler(w, r)
	}
}

func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	file := "templates/index.html"
	data, err := s.dao.ArticleQuery(r.Context(), nil)
	if err != nil {
		log.Printf("query article failed: %+v", err)
		return
	}
	templateHandle(file, w, data)
}

func (s *Server) articleHandler(w http.ResponseWriter, r *http.Request) {
	file := "templates/article.html"
	paths := strings.Split(strings.TrimPrefix(r.URL.Path, "/articles/"), "/")
	if len(paths) == 0 {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	data, err := s.dao.ArticleGet(r.Context(),
		map[string]interface{}{"id": paths[0]})
	if err != nil {
		log.Printf("get article failed: %+v", err)
		return
	}
	templateHandle(file, w, data)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	// TODO: fill body
}
