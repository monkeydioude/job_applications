package internal

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// deleteBreed, through a DELETE method, handles the deletion of a breed.
// It accepts an `id` as URL value.
func (a *App) deleteBreed(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// Decode the `id` URL value from /v1/breed/{id}
	id, err := strconv.Atoi(params["id"])
	if _, ok := params["id"]; !ok {
		a.NotFound(w, err)
		return
	}
	// Actual deletion D:
	if err := a.gormDB.Delete(&Breed{}, id).Error; err != nil {
		a.InternalServerError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
