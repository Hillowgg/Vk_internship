package film

import (
    "net/http"

    "main/internal/api/user"
    "main/internal/logs"
    "main/internal/service/film"
)

type Handler struct {
    serv film.IService
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    endpoint := r.URL.Path
    logs.Log.Infow("Request", "method", r.Method, "endpoint", endpoint)
    switch endpoint {
    case "/film/add":
        user.AdminMiddleware(h.CreateFilm)(w, r)
    case "/film/update":
        user.AdminMiddleware(h.UpdateFilm)(w, r)
    case "/film/delete":
        user.AdminMiddleware(h.DeleteFilm)(w, r)
    default:
        w.WriteHeader(http.StatusNotFound)
    }
}

func NewHandler(serv film.IService) *Handler {
    return &Handler{serv: serv}
}
