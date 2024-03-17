package film

import (
    "encoding/json"
    "net/http"
    "strings"

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

func validateSortBy(sortBy string) string {
    t := []string{"title", "release_date", "rating"}
    sortBy = strings.ToLower(sortBy)
    for _, v := range t {
        if sortBy == v {
            return v
        }
    }
    return "rating"
}

func validateSortType(sortType string) string {
    sortType = strings.ToLower(sortType)
    if sortType == "asc" {
        return sortType
    }
    return "desc"
}

func (h *Handler) GetFilms(w http.ResponseWriter, r *http.Request) {
    sortBy := validateSortBy(r.URL.Query().Get("sortBy"))
    sortType := validateSortType(r.URL.Query().Get("sortType"))

    logs.Log.Infow("Getting films", "sortBy", sortBy, "sortType", sortType)
    films, err := h.serv.GetFilms(r.Context(), sortBy, sortType)
    if err != nil {
        http.Error(w, "Failed to get films", http.StatusInternalServerError)
        logs.Log.Errorw("Failed to get films", "sortBy", sortBy, "sortType", sortType, "err", err)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(films)
}
