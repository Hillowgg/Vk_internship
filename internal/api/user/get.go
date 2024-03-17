package user

import (
    "crypto/rand"
    "encoding/hex"
    "encoding/json"
    "errors"
    "net/http"
    "sync"

    "github.com/google/uuid"
    "main/internal/logs"
    "main/internal/service/user"
)

// func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
//     id := r.Context().Value("userId").(uuid.UUID)
//
//     user, err := h.service.GetUser(r.Context(), id)
//     if err != nil {
//         w.WriteHeader(http.StatusInternalServerError)
//         return
//     }
//     if user == nil {
//         w.WriteHeader(http.StatusNotFound)
//         return
//     }
//     w.WriteHeader(http.StatusOK)
//     err = json.NewEncoder(w).Encode(user)
//     if err != nil {
//         w.WriteHeader(http.StatusInternalServerError)
//         logs.Log.Errorw("Failed to encode user", "err", err)
//         return
//     }
// }

// Register,

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
    var newUser user.NewUser
    err := json.NewDecoder(r.Body).Decode(&newUser)
    newUser.IsAdmin = false

    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        logs.Log.Errorw("Failed to decode user", "err", err)
        return
    }
    id, err := h.serv.CreateUser(r.Context(), &newUser)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        logs.Log.Errorw("Failed to create user", "err", err)
        return
    }
    w.WriteHeader(http.StatusOK)
    err = json.NewEncoder(w).Encode(id)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        logs.Log.Errorw("Failed to encode user", "err", err)
        return
    }
}

var tokens = make(map[string]struct {
    id      uuid.UUID
    isAdmin bool
})
var mutex = &sync.RWMutex{}

func randomHex(n int) (string, error) {
    bytes := make([]byte, n)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return hex.EncodeToString(bytes), nil
}

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
    u, err := h.serv.CheckLoginCredentials(r.Context(), login.Login, login.Password)
    if errors.Is(err, user.WrongCredentials) {
        w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte("Wrong credentials"))
        return
    }
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        logs.Log.Errorw("Failed to check login credentials", "err", err)
        return
    }

    token, _ := randomHex(20)
    mutex.Lock()
    tokens[token] = struct {
        id      uuid.UUID
        isAdmin bool
    }{u.Id, u.IsAdmin}
    mutex.Unlock()
    w.WriteHeader(http.StatusOK)
    err = json.NewEncoder(w).Encode(token)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        logs.Log.Errorw("Failed to encode token", "err", err)
        return
    }
}
