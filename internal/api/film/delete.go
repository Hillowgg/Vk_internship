package film

import (
    "io"
    "net/http"

    "github.com/tidwall/gjson"
)

func (h *Handler) DeleteFilm(w http.ResponseWriter, r *http.Request) {
    body, _ := io.ReadAll(r.Body)
    filmId := int32(gjson.GetBytes(body, "Id").Int())
    err := h.serv.DeleteFilm(r.Context(), filmId)
    if err != nil {
        http.Error(w, "Failed to delete film", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}
