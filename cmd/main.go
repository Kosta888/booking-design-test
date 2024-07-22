package main

import (
	"applicationDesignTest/internal/app"
	"net/http"
	"os"
)

func main() {
    logLevel := "info"
    if lvl, ok := os.LookupEnv("LOG_LEVEL"); ok {
        logLevel = lvl
    }

    // Инициализация приложения с уровнем логирования
    application := app.NewApp(logLevel)
    
    // Запуск HTTP сервера
    application.Logger.Info("Starting server on :8080")
    http.ListenAndServe("0.0.0.0:8080", application.Router)
}
