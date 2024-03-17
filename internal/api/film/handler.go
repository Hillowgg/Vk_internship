package film

import (
    "net/http"

    "main/internal/api/middleware"
    "main/internal/logs"
    "main/internal/service/film"
)

type Handler struct {
    serv film.IService
    mw   *middleware.Middleware
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    endpoint := r.URL.Path
    logs.Log.Infow("Request", "method", r.Method, "endpoint", endpoint)
    switch endpoint {
    case "/film/search":
        h.mw.UserMiddleware(h.SearchByActorAndTitle)(w, r)
    case "/film/add":
        h.mw.AdminMiddleware(h.CreateFilm)(w, r)
    case "/film/update":
        h.mw.AdminMiddleware(h.UpdateFilm)(w, r)
    case "/film/delete":
        h.mw.AdminMiddleware(h.DeleteFilm)(w, r)
    default:
        w.WriteHeader(http.StatusNotFound)
    }
}

func NewHandler(serv film.IService, mw *middleware.Middleware) *Handler {
    return &Handler{serv: serv, mw: mw}
}
