package api

import (
    "main/internal/api/film"
    "main/internal/api/user"
    "main/internal/database"
    filmservice "main/internal/service/film"
    userservice "main/internal/service/user"
)

type API struct {
    User *user.Handler
    Film *film.Handler
}

func NewAPI(db database.QuerierWithTx) *API {
    return &API{
        User: user.NewHandler(userservice.NewService(db)),
        Film: film.NewHandler(filmservice.NewService(db)),
    }
}
