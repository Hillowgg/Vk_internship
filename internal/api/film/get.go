package film

import (
    "encoding/json"
    "net/http"

    "main/internal/logs"
)

func (h *Handler) SearchByActorAndTitle(w http.ResponseWriter, r *http.Request) {
    actor := r.URL.Query().Get("actor")
    title := r.URL.Query().Get("title")
    if actor == "" || title == "" {
        http.Error(w, "Invalid query", http.StatusBadRequest)
        return
    }
    film, err := h.serv.SearchFilmByActor(r.Context(), title, actor)
    if err != nil {
        http.Error(w, "Failed to search", http.StatusInternalServerError)
        logs.Log.Errorw("Failed to search", "actor", actor, "title", title, "err", err)
        return
    }
    if film == nil {
        http.Error(w, "Film not found", http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(film)
}
