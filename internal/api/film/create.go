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

func validateNewFilm(bytes []byte) (*film.NewFilm, error) {
    json := gjson.ParseBytes(bytes)
    date, err := time.Parse("2006-01-02", json.Get("ReleaseDate").String())
    if err != nil {
        return nil, err
    }
    f := &film.NewFilm{
        Title:       json.Get("Title").String(),
        Description: json.Get("Description").String(),
        ReleaseDate: date,
        Rating:      int8(json.Get("Rating").Int()),
    }
    if f.Title == "" || f.Description == "" || f.Rating < 1 || f.Rating > 10 {
        logs.Log.Errorw("Invalid film data", "film", f)
        return nil, errors.New("invalid film data")
    }
    return f, nil
}

func (h *Handler) CreateFilm(w http.ResponseWriter, r *http.Request) {
    res, _ := io.ReadAll(r.Body)
    f, err := validateNewFilm(res)
    if err != nil {
        logs.Log.Errorw("Failed to validate film", "film", f, "err", err)
        http.Error(w, "Failed to validate film", http.StatusBadRequest)
        return
    }
    id, err := h.serv.CreateFilm(r.Context(), f)
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
