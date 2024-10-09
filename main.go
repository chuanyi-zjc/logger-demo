package main

import (
    "chuanyi-zjc/logdemo/logger"
    "time"
)

func main() {
    log := logger.NewLogger(logger.DEBUG)
    go func() {
        for {
            log.Debug("This is a debug message.")
            time.Sleep(100 * time.Millisecond)
        }
    }()
    go func() {
        for {
            log.Info("This is an info message.")
            time.Sleep(200 * time.Millisecond)
        }
    }()
    go func() {
        for {
            log.Warning("This is a warning message.")
            time.Sleep(300 * time.Millisecond)
        }
    }()
    go func() {
        for {
            log.Error("This is an error message.")
            time.Sleep(400 * time.Millisecond)
        }
    }()
    time.Sleep(5 * time.Second)
}
