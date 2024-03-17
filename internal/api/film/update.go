package film

import (
    "encoding/json"
    "net/http"

    "main/internal/database"
)

func (h *Handler) UpdateFilm(w http.ResponseWriter, r *http.Request) {
    var film database.OptUpdateFilm
    err := json.NewDecoder(r.Body).Decode(&film)
    if err != nil {
        http.Error(w, "Failed to decode request", http.StatusBadRequest)
        return
    }
    err = h.serv.UpdateFilm(r.Context(), &film)
    if err != nil {
        http.Error(w, "Failed to update film", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}
