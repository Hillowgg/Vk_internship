package main

import (
    "main/internal/app"
    "main/internal/logs"
)

func main() {
    a, err := app.NewApp()
    if err != nil {
        logs.Log.Fatalw("Failed to create app", "err", err)
    }
    err = a.Run()
    if err != nil {
        logs.Log.Fatalw("Failed to run app", "err", err)
    }
}
