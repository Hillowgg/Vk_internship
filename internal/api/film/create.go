package film

import (
    "encoding/json"
    "net/http"

    "main/internal/logs"
    "main/internal/service/film"
)

func (h *Handler) CreateFilm(w http.ResponseWriter, r *http.Request) {
    var f film.NewFilm
    err := json.NewDecoder(r.Body).Decode(&f)
    if err != nil {
        logs.Log.Errorw("Failed to decode film", "err", err)
        http.Error(w, "Failed to decode film", http.StatusBadRequest)
        return
    }
    _, err = h.serv.CreateFilm(r.Context(), &f)
    if err != nil {
        logs.Log.Errorw("Failed to create film", "film", f, "err", err)
        http.Error(w, "Failed to create film", http.StatusInternalServerError)
        return
    }
    logs.Log.Infow("Added film", "film", f)
    w.WriteHeader(http.StatusOK)
}
