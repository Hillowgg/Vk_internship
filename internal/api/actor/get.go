package actor

import (
    "encoding/json"
    "net/http"
    "strconv"
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
