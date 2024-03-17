package film

import (
    "net/http"

    "main/internal/api/user"
    "main/internal/service/film"
)

type Handler struct {
    serv film.IService
}

func (h *Handler) ServeHttp(w http.ResponseWriter, r *http.Request) {
    endpoint := r.URL.Path
    switch endpoint {
    case "/film/add":
        user.AdminMiddleware(h.CreateFilm)(w, r)
    case "/film/update":
        user.AdminMiddleware(h.UpdateFilm)(w, r)
    case "/film/delete":
        user.AdminMiddleware(h.DeleteFilm)(w, r)
    }
}

func NewHandler(serv film.IService) *Handler {
    return &Handler{serv: serv}
}
