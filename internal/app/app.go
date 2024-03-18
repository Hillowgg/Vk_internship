package app

import (
    "context"
    "net/http"
    "os"

    "github.com/jackc/pgx/v5/pgxpool"
    "main/internal/api"
    "main/internal/cache"
    "main/internal/database"
    "main/internal/logs"
)

type App struct {
    api   *api.API
    db    database.QuerierWithTx
    cache cache.ICache
}

func NewApp() (*App, error) {
    app := &App{}
    app.initDataBase()
    app.initCache()
    app.initAPI()
    return app, nil
}

func (a *App) initDataBase() {
    ctx := context.Background()
    conn, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
    if err != nil {
        logs.Log.Fatalw("Failed to connect to database", "err", err)
    }
    err = conn.Ping(context.Background())
    if err != nil {
        logs.Log.Fatalw("Failed to ping database", "err", err)
    }
    logs.Log.Infow("Connected to database")
    a.db = database.New(conn)
}

func (a *App) initCache() {
    a.cache = cache.NewCache()
    logs.Log.Infow("Connected to cache")
}

func (a *App) initAPI() {
    a.api = api.NewAPI(a.db, a.cache)
}

func (a *App) Run() error {
    mux := http.NewServeMux()
    mux.Handle("/user/", a.api.User)
    mux.Handle("/film/", a.api.Film)
    mux.Handle("/actor/", a.api.Actor)
    host := ":" + os.Getenv("PORT")
    logs.Log.Infow("Starting server", "host", host)
    err := http.ListenAndServe(host, mux)
    return err
}
