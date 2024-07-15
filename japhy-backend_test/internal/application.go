package internal

import (
	charmLog "github.com/charmbracelet/log"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	logger *charmLog.Logger
	gormDB *gorm.DB
}

func NewApp(logger *charmLog.Logger, gormDB *gorm.DB) *App {
	if gormDB == nil {
		panic("gorm nil")
	}
	return &App{
		logger: logger,
		gormDB: gormDB,
	}
}

func (a *App) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/breed", a.createBreed).Methods("POST")
	r.HandleFunc("/breed", a.getBreed).Methods("GET")
	r.HandleFunc("/breed/{id}", a.updateBreed).Methods("PUT")
	r.HandleFunc("/breed/{id}", a.deleteBreed).Methods("DELETE")
}
