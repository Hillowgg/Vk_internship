package user

import (
    "encoding/json"
    "errors"
    "io"
    "net/http"
    "regexp"

    "github.com/tidwall/gjson"
    "main/internal/logs"
    "main/internal/service/user"
)

func validateUser(bytes []byte) (*user.NewUser, error) {
    newUser := &user.NewUser{
        Nickname: gjson.GetBytes(bytes, "login").String(),
        Email:    gjson.GetBytes(bytes, "email").String(),
        Password: gjson.GetBytes(bytes, "password").String(),
    }

    if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(newUser.Nickname) ||
        len(newUser.Nickname) < 3 ||
        len(newUser.Nickname) > 20 {
        return nil, errors.New("invalid nickname")
    }
    if !regexp.MustCompile(`^[a-zA-Z0-9]*@[a-zA-Z0-9]*\.[a-zA-Z0-9]*$`).MatchString(newUser.Email) {
        return nil, errors.New("invalid email")
    }
    if len(newUser.Password) < 8 {
        return nil, errors.New("invalid password")
    }
    if newUser.Nickname == "" || newUser.Email == "" || newUser.Password == "" {
        return nil, errors.New("empty fields")
    }
    return newUser, nil
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
    body, _ := io.ReadAll(r.Body)
    newUser, err := validateUser(body)

    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        logs.Log.Errorw("Failed to decode user", "err", err)
        return
    }
    id, err := h.userServ.CreateUser(r.Context(), newUser)
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
