package actor

import (
    "io"
    "net/http"
    "time"

    "github.com/tidwall/gjson"
    "github.com/tidwall/sjson"
    "main/internal/database"
    "main/internal/service/actor"
)

func validateActor(bytes []byte) (*actor.NewActor, error) {
    json := gjson.ParseBytes(bytes)
    birthday, err := time.Parse("2006-01-02", json.Get("Birthday").String())
    if err != nil {
        return nil, err
    }
    var gender database.Gender
    err = gender.Scan(json.Get("Gender").String())
    if err != nil {
        return nil, err
    }
    a := &actor.NewActor{
        Name:     json.Get("Name").String(),
        Birthday: birthday,
        Gender:   gender,
    }
    return a, nil
}

func (h *Handler) CreateActor(w http.ResponseWriter, r *http.Request) {
    body, _ := io.ReadAll(r.Body)
    actor, err := validateActor(body)
    if err != nil {
        http.Error(w, "Invalid actor", http.StatusBadRequest)
        return
    }
    id, err := h.serv.CreateActor(r.Context(), actor)
    if err != nil {
        http.Error(w, "Failed to create actor", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    json, _ := sjson.SetBytes([]byte{}, "id", id)
    w.Write(json)
}
