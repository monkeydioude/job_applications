package internal

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// updateBreed, through a PUT method, handles the modification of a breed.
// It accepts a JSON payload and an `id` as URL value.
// Specifying `id` in the payload does nothing, as it is auto-incremental.
// Using an already existing `name` won't update the row as `name` is a unique key.
func (a *App) updateBreed(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var breed Breed
	if err := a.gormDB.First(&breed, id).Error; err != nil {
		a.NotFound(w, errors.New("breed not found"))
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&breed); err != nil {
		a.InternalServerError(w, err)
		return
	}
	breed.ID = &id
	if err := a.gormDB.Model(&breed).Updates(breed).Error; err != nil {
		http.Error(w, "Failed to update breed", http.StatusInternalServerError)
		return
	}
	a.Ok(w, breed)
}
