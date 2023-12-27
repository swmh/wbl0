package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) GetHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	b, err := s.service.Get(r.Context(), id)
	if err != nil {
		if s.service.IsNotFound(err) {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Write(b)
}
