package user

import (
    "net/http"
    "strings"
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

        next(w, r)
    })
}

func AdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        bearer := r.Header.Get("Authorization")
        token := strings.TrimPrefix(bearer, "Bearer ")
        if token == "" {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }

        mutex.RLock()
        user, ok := tokens[token]
        mutex.RUnlock()

        if !ok {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        if !user.isAdmin {
            w.WriteHeader(http.StatusForbidden)
            return
        }

        next(w, r)
    })
}
