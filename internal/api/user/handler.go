package user

import (
    "net/http"

    "main/internal/api/middleware"
    "main/internal/logs"
    "main/internal/service/session"
    "main/internal/service/user"
)

type Handler struct {
    userServ    user.IService
    sessionServ session.IService
    mw          *middleware.Middleware
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    endpoint := r.Method + " " + r.URL.Path
    logs.Log.Infow("Request", "method", r.Method, "endpoint", endpoint)
    switch endpoint {
    case "POST /user/register":
        h.Register(w, r)
    case "GET /user/login":
        h.Login(w, r)
    }
}

func NewHandler(user user.IService, session session.IService, mw *middleware.Middleware) *Handler {
    return &Handler{userServ: user, sessionServ: session, mw: mw}
}
