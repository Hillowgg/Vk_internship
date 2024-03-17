package actor

import (
    "io"
    "net/http"

    "github.com/tidwall/gjson"
)

func (h *Handler) UpdateActor(w http.ResponseWriter, r *http.Request) {
    body, _ := io.ReadAll(r.Body)
    actorId := int32(gjson.GetBytes(body, "Id").Int())
    keys := [...]string{"Name", "Birthday", "Gender"}
    actor := make(map[string]any)
    actor["Id"] = actorId
    for _, key := range keys {
        actor[key] = gjson.GetBytes(body, key).Value()
    }
    err := h.serv.UpdateActor(r.Context(), actor)
    if err != nil {
        http.Error(w, "Failed to update actor", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}
