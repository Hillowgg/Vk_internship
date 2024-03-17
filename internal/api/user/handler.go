package user

import (
    "net/http"

    "main/internal/service/user"
)

type Handler struct {
    serv user.IService
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    endpoint := r.URL.Path
    switch endpoint {
    case "/user/register":
        h.Register(w, r)
    }
}

func NewHandler(service user.IService) *Handler {
    return &Handler{serv: service}
}
