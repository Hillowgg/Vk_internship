package app

import (
    "context"
    "net/http"

    "github.com/jackc/pgx/v5/pgxpool"
    "main/internal/api"
    "main/internal/database"
    "main/internal/logs"
)

type App struct {
    api *api.API
    db  database.QuerierWithTx
}

func NewApp() (*App, error) {
    app := &App{}
    app.initDataBase()
    app.initAPI()
    return app, nil
}

func (a *App) initDataBase() {
    ctx := context.Background()
    conn, err := pgxpool.New(ctx, "postgres://postgres:postgres@localhost:5432/internship")
    if err != nil {
        logs.Log.Fatalw("Failed to connect to database", "err", err)
    }
    logs.Log.Infow("Connected to database")
    a.db = database.New(conn)
}

func (a *App) initAPI() {
    a.api = api.NewAPI(a.db)
}

func (a *App) Run() error {
    mux := http.NewServeMux()
    mux.Handle("/user/", a.api.User)
    mux.Handle("/film/", a.api.Film)
    logs.Log.Infow("Starting server", "port", 8080)
    err := http.ListenAndServe(":8080", mux)
    return err
}
