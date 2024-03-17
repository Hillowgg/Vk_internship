package api

import (
    "main/internal/api/film"
    middleware "main/internal/api/middleware"
    "main/internal/api/user"
    "main/internal/cache"
    "main/internal/database"
    filmservice "main/internal/service/film"
    "main/internal/service/session"
    userservice "main/internal/service/user"
)

type API struct {
    User *user.Handler
    Film *film.Handler
}

func NewAPI(db database.QuerierWithTx, cache cache.ICache) *API {
    mw := middleware.NewMiddleware(session.NewService(cache))
    return &API{
        User: user.NewHandler(userservice.NewService(db), session.NewService(cache), mw),
        Film: film.NewHandler(filmservice.NewService(db), mw),
    }
}
