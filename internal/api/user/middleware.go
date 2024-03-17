package user

import (
    "net/http"
    "strings"

    "main/internal/logs"
)

func UserMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        bearer := r.Header.Get("Authorization")
        token := strings.TrimPrefix(bearer, "Bearer ")
        if token == "" {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        mutex.RLock()
        _, ok := tokens[token]
        mutex.RUnlock()

        if !ok {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        logs.Log.Infow("User request", "token", token)
        next(w, r)
    })
}

func AdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        bearer := r.Header.Get("Authorization")
        token := strings.TrimPrefix(bearer, "Bearer ")
        if token == "" {
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte("No token"))
            return
        }

        mutex.RLock()
        user, ok := tokens[token]
        mutex.RUnlock()

        if !ok {
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte("Invalid token"))
            return
        }
        if !user.isAdmin {
            w.WriteHeader(http.StatusForbidden)
            w.Write([]byte("Not admin"))
            return
        }
        logs.Log.Infow("Admin request", "user", user)
        next(w, r)
    })
}
