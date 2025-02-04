package app

import (
	"backend/internal/database"
	"backend/internal/services"
	"backend/internal/transport/rest"
	"context"
	"net/http"
)

type App struct {
	pool database.DBPool
	port string
}

func NewApp(pool database.DBPool, port string) *App {
	return &App{pool: pool, port: port}
}

func (a *App) Run() error {
	db := database.NewPGXDatabase(a.pool)
	err := db.CreateTableQuery(context.Background())
	if err != nil {
		return err
	}
	service := services.NewService(db)
	handler := rest.NewHandler(service)
	http.HandleFunc("/api/status", handler.GetStatusesData)
	http.HandleFunc("/api/ping", handler.UpdateStatusData)
	err = http.ListenAndServe(":"+a.port, nil)
	return err
}
