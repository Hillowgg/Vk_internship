package film

import (
    "errors"
    "io"
    "net/http"
    "time"

    "github.com/tidwall/gjson"
    "github.com/tidwall/sjson"
    "main/internal/logs"
    "main/internal/service/film"
)

func validateNewFilm(bytes []byte) (*film.NewFilm, []int32, error) {
    json := gjson.ParseBytes(bytes)
    date, err := time.Parse("2006-01-02", json.Get("release_date").String())
    if err != nil {
        return nil, nil, err
    }
    f := &film.NewFilm{
        Title:       json.Get("title").String(),
        Description: json.Get("description").String(),
        ReleaseDate: date,
        Rating:      int8(json.Get("rating").Int()),
    }
    actors := json.Get("actors").Array()
    actorsIds := make([]int32, 0, len(actors))
    for _, actor := range actors {
        actorsIds = append(actorsIds, int32(actor.Int()))
    }
    if len(f.Title) < 1 || len(f.Title) > 150 || len(f.Description) < 1 ||
        len(f.Description) > 1000 || f.Rating < 1 || f.Rating > 10 {
        logs.Log.Errorw("Invalid film data", "film", f)
        return nil, nil, errors.New("invalid film data")
    }
    return f, actorsIds, nil
}

func (h *Handler) CreateFilm(w http.ResponseWriter, r *http.Request) {
    res, _ := io.ReadAll(r.Body)
    f, actors, err := validateNewFilm(res)
    if err != nil {
        logs.Log.Errorw("Failed to validate film", "film", f, "err", err)
        http.Error(w, "Failed to validate film", http.StatusBadRequest)
        return
    }
    id, err := h.serv.CreateFilmWithActors(r.Context(), f, actors)
    if err != nil {
        logs.Log.Errorw("Failed to create film", "film", f, "err", err)
        http.Error(w, "Failed to create film", http.StatusInternalServerError)
        return
    }
    logs.Log.Infow("Added film", "film", f)
    w.WriteHeader(http.StatusOK)
    json, _ := sjson.SetBytes([]byte{}, "id", id)
    w.Write(json)
}
