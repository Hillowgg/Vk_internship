package film

import (
    "net/http"

    "github.com/tidwall/sjson"
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
    var json []byte
    json, _ = sjson.SetBytes(json, "id", film.Id)
    json, _ = sjson.SetBytes(json, "title", film.Title)
    json, _ = sjson.SetBytes(json, "description", film.Description)
    json, _ = sjson.SetBytes(json, "releaseDate", film.ReleaseDate)
    json, _ = sjson.SetBytes(json, "rating", film.Rating)
    w.WriteHeader(http.StatusOK)
    w.Write(json)
}
