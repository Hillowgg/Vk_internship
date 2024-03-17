package film

import (
    "encoding/json"
    "net/http"
)

func (h *Handler) DeleteFilm(w http.ResponseWriter, r *http.Request) {
    var filmId int32
    err := json.NewDecoder(r.Body).Decode(&filmId)
    if err != nil {
        http.Error(w, "Failed to decode request", http.StatusBadRequest)
        return
    }
    err = h.serv.DeleteFilm(r.Context(), filmId)
    if err != nil {
        http.Error(w, "Failed to delete film", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}
