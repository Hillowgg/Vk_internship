package api

import (
    "main/internal/api/actor"
    "main/internal/api/film"
    middleware "main/internal/api/middleware"
    "main/internal/api/user"
    "main/internal/cache"
    "main/internal/database"
    actorservice "main/internal/service/actor"
    filmservice "main/internal/service/film"
    "main/internal/service/session"
    userservice "main/internal/service/user"
)

type API struct {
    User  *user.Handler
    Film  *film.Handler
    Actor *actor.Handler
}

func NewAPI(db database.QuerierWithTx, cache cache.ICache) *API {
    mw := middleware.NewMiddleware(session.NewService(cache))
    return &API{
        User:  user.NewHandler(userservice.NewService(db), session.NewService(cache), mw),
        Film:  film.NewHandler(filmservice.NewService(db), mw),
        Actor: actor.NewHandler(actorservice.NewService(db), mw),
    }
}
