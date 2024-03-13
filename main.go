package main

import (
    "log"
    "os"
)

var (
    InfoLog  *log.Logger
    WarnLog  *log.Logger
    ErrorLog *log.Logger
    FatalLog *log.Logger
)

func init() {
    flags := log.LstdFlags | log.Lshortfile
    InfoLog = log.New(os.Stdout, "[INFO] ", flags)
    WarnLog = log.New(os.Stdout, "[WARN] ", flags)
    ErrorLog = log.New(os.Stderr, "[ERROR] ", flags)
    FatalLog = log.New(os.Stderr, "[FATAL] ", flags)
}

func main() {

}
