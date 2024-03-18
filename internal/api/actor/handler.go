package actor

import (
    "net/http"

    "main/internal/api/middleware"
    "main/internal/service/actor"
)

type Handler struct {
    serv actor.IService
    mw   *middleware.Middleware
}

func NewHandler(serv actor.IService, mw *middleware.Middleware) *Handler {
    return &Handler{serv: serv, mw: mw}
}
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    endpoint := r.Method + " " + r.URL.Path

    switch endpoint {
    case "GET /actor/get":
        h.mw.UserMiddleware(h.GetActor)(w, r)
    case "GET /actor/get_with_films":
        h.mw.UserMiddleware(h.GetActorsWithFilms)(w, r)
    case "PUT /actor/create":
        h.mw.AdminMiddleware(h.CreateActor)(w, r)
    case "POST /actor/update":
        h.mw.AdminMiddleware(h.UpdateActor)(w, r)
    case "DELETE /actor/delete":
        h.mw.AdminMiddleware(h.DeleteActor)(w, r)

    default:
        w.WriteHeader(http.StatusNotFound)
    }
}
