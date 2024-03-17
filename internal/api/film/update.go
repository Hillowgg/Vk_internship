package film

import (
    "io"
    "net/http"

    "github.com/tidwall/gjson"
    "main/internal/logs"
)

func (h *Handler) UpdateFilm(w http.ResponseWriter, r *http.Request) {
    film := make(map[string]any)
    body, _ := io.ReadAll(r.Body)
    json := gjson.ParseBytes(body)
    if !json.Get("Id").Exists() {
        http.Error(w, "id must be present", http.StatusBadRequest)
        logs.Log.Error("Invalid request")
        return
    }
    keys := []string{"Title", "Description", "ReleaseDate", "Rating"}
    film["Id"] = int32(json.Get("Id").Int())
    for _, key := range keys {
        if json.Get(key).Exists() {
            film[key] = json.Get(key).Value()
        }
    }
    if len(film) < 2 {
        http.Error(w, "At least 1 field and id must be present", http.StatusBadRequest)
        logs.Log.Error("Invalid request")
        return
    }

    err := h.serv.UpdateFilm(r.Context(), film)
    if err != nil {
        http.Error(w, "Failed to update film", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}
