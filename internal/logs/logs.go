package logs

import (
    "os"

    "go.uber.org/zap"
)

var Log *zap.SugaredLogger

func init() {
    var nonSugar *zap.Logger
    var err error

    switch os.Getenv("log") {
    case "DEV":
        nonSugar, err = zap.NewDevelopment()
    default:
        nonSugar, err = zap.NewProduction()
    }
    if err != nil {
        panic(err)
    }
    Log = nonSugar.Sugar()
}
