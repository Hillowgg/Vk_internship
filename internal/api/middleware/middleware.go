package middleware

import (
    "net/http"
    "strings"

    "main/internal/logs"
    "main/internal/service/session"
)

type Middleware struct {
    session session.IService
}

func (mw *Middleware) UserMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        bearer := r.Header.Get("Authorization")
        token := strings.TrimPrefix(bearer, "Bearer ")
        if token == "" {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        ok := mw.session.IsUser(r.Context(), token)

        if !ok {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        logs.Log.Infow("User request", "token", token)
        next(w, r)
    })
}

func (mw *Middleware) AdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        bearer := r.Header.Get("Authorization")
        token := strings.TrimPrefix(bearer, "Bearer ")
        if token == "" {
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte("No token"))
            return
        }

        ok := mw.session.IsAdmin(r.Context(), token)

        if !ok {
            w.WriteHeader(http.StatusForbidden)
            w.Write([]byte("Invalid token"))
            return
        }
        logs.Log.Infow("Admin request", "token", token)
        next(w, r)
    })
}

func NewMiddleware(session session.IService) *Middleware {
    return &Middleware{session: session}
}
