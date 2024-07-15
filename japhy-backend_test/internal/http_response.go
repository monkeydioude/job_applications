package internal

import (
	"encoding/json"
	"net/http"
)

func HttpResponse(w http.ResponseWriter, statusCode int, encode any) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(encode)
}

func (a *App) InternalServerError(w http.ResponseWriter, err error) {
	HttpResponse(w, http.StatusInternalServerError, err.Error())
	if a.logger != nil {
		a.logger.Error(err.Error())
	}
}

func (a *App) BadRequest(w http.ResponseWriter, err error) {
	HttpResponse(w, http.StatusBadRequest, err.Error())
	if a.logger != nil {
		a.logger.Error(err.Error())
	}
}

func (a *App) NotFound(w http.ResponseWriter, err error) {
	HttpResponse(w, http.StatusNotFound, err.Error())
	if a.logger != nil {
		a.logger.Error(err.Error())
	}
}

func (a *App) Ok(w http.ResponseWriter, encode any) {
	w.Header().Set("Content-Type", "application/json")
	HttpResponse(w, http.StatusOK, encode)
}

func (a *App) Created(w http.ResponseWriter, encode any) {
	w.Header().Set("Content-Type", "application/json")
	HttpResponse(w, http.StatusCreated, encode)
}
