package internal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestICanGetABreedByItsMaleWeight(t *testing.T) {
	gormDB, sqlMock, logger := setupTest()
	rows := sqlMock.
		NewRows([]string{"id", "species", "pet_size", "name", "average_male_adult_weight", "average_female_adult_weight"}).
		AddRow(1, "dog", "small", "affenpinscher", 6000, 5000)
	sqlMock.
		ExpectQuery("^SELECT \\* FROM `breeds` WHERE average_male_adult_weight <= \\?$").
		WithArgs("6000").
		WillReturnRows(rows)

	app := NewApp(logger, gormDB)
	// Mock a request to /breed with a query string
	req, err := http.NewRequest("GET", "/v1/breed?average_male_adult_weight[lte]=6000", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Mock a ResponseWriter
	rr := httptest.NewRecorder()
	// Mock the gorilla router
	router := mux.NewRouter()
	router.HandleFunc("/v1/breed", app.getBreed).Methods("GET")
	// Actual requesting of the route
	router.ServeHTTP(rr, req)

	if db, err := gormDB.DB(); err != nil {
		db.Close()
	}
	// Status, are you ok?
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Validation of the response
	var breeds []Breed
	if err := json.NewDecoder(rr.Body).Decode(&breeds); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}
	if len(breeds) != 1 || breeds[0].Name != "affenpinscher" {
		t.Errorf("handler returned unexpected body: got %v", breeds)
	}
	// Check the db was indeed requested using the expected parameters
	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Fatalf("could not meet expectations: %s", err)
	}
}
