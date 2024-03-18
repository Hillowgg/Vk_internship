package user

import (
    "encoding/json"
    "errors"
    "net/http"

    "main/internal/logs"
    "main/internal/service/user"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
    var login struct {
        Login    string
        Password string
    }
    err := json.NewDecoder(r.Body).Decode(&login)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        logs.Log.Errorw("Failed to decode login", "err", err, "body", r.Body)
        return
    }
    u, err := h.userServ.CheckLoginCredentials(r.Context(), login.Login, login.Password)
    if errors.Is(err, user.WrongCredentials) {
        http.Error(w, "Wrong credentials", http.StatusUnauthorized)
        w.Write([]byte("Wrong credentials"))
        return
    }
    if err != nil {
        http.Error(w, "Failed to check login credentials", http.StatusInternalServerError)
        logs.Log.Errorw("Failed to check login credentials", "err", err)
        return
    }
    token, err := h.sessionServ.CreateSession(r.Context(), u.Id, u.IsAdmin)
    if err != nil {
        http.Error(w, "Failed to create session", http.StatusInternalServerError)
        logs.Log.Errorw("Failed to create session", "err", err)
        return
    }
    w.WriteHeader(http.StatusOK)
    err = json.NewEncoder(w).Encode(token)
    if err != nil {
        http.Error(w, "Failed to encode token", http.StatusInternalServerError)
        logs.Log.Errorw("Failed to encode token", "err", err)
        return
    }
}
