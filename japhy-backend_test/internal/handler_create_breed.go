package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// createBreed, through a POST method, handles the creation of a breed.
// It accepts a JSON payload.
// Specifying `id` in the payload does nothing, as it is auto-incremental.
// Using an already existing `name` won't succeed as `name` is a unique key.
func (a *App) createBreed(w http.ResponseWriter, r *http.Request) {
	// Decoding the payload
	var breed Breed
	if err := json.NewDecoder(r.Body).Decode(&breed); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	breed.ID = nil

	// Gently verify if the breed exists
	var breedExist Breed
	a.gormDB.
		Model(&Breed{}).
		Where("name = ?", breed.Name).
		First(&breedExist)
	if breedExist.ID != nil {
		a.BadRequest(w, fmt.Errorf("A `%s` breed already exists", breed.Name))
		return
	}
	// Save the new breed to the database
	if err := a.gormDB.Create(&breed).Error; err != nil {
		a.InternalServerError(w, err)
		return
	}
	a.Created(w, breed)
}
