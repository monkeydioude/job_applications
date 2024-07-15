package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
)

func TestICanDeleteABreedByID(t *testing.T) {
	gormDB, sqlMock, logger := setupTest()

	sqlMock.ExpectBegin()
	sqlMock.
		ExpectExec("^DELETE FROM `breeds` WHERE `breeds`.`id` = \\?$").
		WithArgs(3).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	app := NewApp(logger, gormDB)

	// Mock a ResponseWriter
	rr := httptest.NewRecorder()
	// Mock the gorilla router
	router := mux.NewRouter()
	router.HandleFunc("/breed/{id}", app.deleteBreed).Methods("DELETE")
	// Mock a request to /breed with an id
	req, err := http.NewRequest("DELETE", "/breed/3", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Actual requesting of the route
	router.ServeHTTP(rr, req)

	if db, err := gormDB.DB(); err != nil {
		db.Close()
	}
	// Status, are you ok?
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}

	// Check the db was indeed requested using the expected parameters
	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Fatalf("could not meet expectations: %s", err)
	}
}
