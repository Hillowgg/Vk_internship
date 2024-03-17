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
    endpoint := r.URL.Path
    switch endpoint {
    case "/actor/get":
        h.mw.UserMiddleware(h.GetActor)(w, r)
    case "/actor/create":
        h.mw.AdminMiddleware(h.CreateActor)(w, r)
    case "/actor/update":
        h.mw.AdminMiddleware(h.UpdateActor)(w, r)
    case "/actor/delete":
        h.mw.AdminMiddleware(h.DeleteActor)(w, r)

    default:
        w.WriteHeader(http.StatusNotFound)
    }
}
