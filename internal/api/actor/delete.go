package actor

import (
    "net/http"
    "strconv"
)

func (h *Handler) DeleteActor(w http.ResponseWriter, r *http.Request) {
    actorId, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil {
        http.Error(w, "Invalid actor id", http.StatusBadRequest)
        return
    }
    err = h.serv.DeleteActor(r.Context(), int32(actorId))
    if err != nil {
        http.Error(w, "Failed to delete actor", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}
