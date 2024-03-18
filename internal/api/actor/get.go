package actor

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/tidwall/sjson"
)

func (h *Handler) GetActor(w http.ResponseWriter, r *http.Request) {
    actorId, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil {
        http.Error(w, "Invalid actor id", http.StatusBadRequest)
        return
    }
    actor, err := h.serv.GetActor(r.Context(), int32(actorId))
    if err != nil {
        http.Error(w, "Failed to get actor", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(actor)
}

func (h *Handler) GetActorsWithFilms(w http.ResponseWriter, r *http.Request) {
    actors, err := h.serv.GetActorsWithFilms(r.Context())
    if err != nil {
        http.Error(w, "Failed to get actors with films", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    err = json.NewEncoder(w).Encode(actors)
    var js []byte
    for k, v := range actors {
        js, _ = sjson.SetBytes(js, ".-1", map[string]interface{}{"actor": k, "films": v})
    }
    w.Write(js)
}
