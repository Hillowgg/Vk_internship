package user

import (
    "net/http"

    "main/internal/logs"
    "main/internal/service/user"
)

type Handler struct {
    serv user.IService
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    endpoint := r.URL.Path
    logs.Log.Infow("Request", "method", r.Method, "endpoint", endpoint)
    switch endpoint {
    case "/user/register":
        h.Register(w, r)
    case "/user/login":
        h.Login(w, r)
    }
}

func NewHandler(service user.IService) *Handler {
    return &Handler{serv: service}
}
