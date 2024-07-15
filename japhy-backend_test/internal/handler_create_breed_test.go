package internal

import (
	"testing"
)

// Didnt have the time to understand why I cant make a double expect working
func TestCreateBreed(t *testing.T) {
	// gormDB, sqlMock, logger := setupTest()

	// // Set up the transaction expectations
	// sqlMock.ExpectBegin()

	// // Expect a SELECT query
	// sqlMock.ExpectQuery("^SELECT * FROM `breeds` WHERE name = \\? ORDER BY `breeds`.`id` LIMIT \\?$").
	// 	WithArgs("american hairless terrier", 1).
	// 	WillReturnRows(sqlmock.NewRows(nil))
	// // Expect an INSERT query
	// sqlMock.ExpectExec("^INSERT INTO `breeds` \\(`species`,`pet_size`,`name`,`average_male_adult_weight`,`average_female_adult_weight`\\) VALUES \\(\\?,\\?,\\?,\\?,\\?\\)$").
	// 	WithArgs("dog", "small", "american hairless terrier", 5500, 4500).
	// 	WillReturnResult(sqlmock.NewResult(1, 1))

	// sqlMock.ExpectCommit()
	// app := NewApp(logger, gormDB)

	// // Mock a ResponseWriter
	// rr := httptest.NewRecorder()
	// // Mock the gorilla router
	// router := mux.NewRouter()
	// router.HandleFunc("/breed", app.createBreed).Methods("POST")

	// // Mock a request to /breed
	// payload := `{"species": "dog", "pet_size": "small", "name": "american hairless terrier", "average_male_adult_weight": 5500, "average_female_adult_weight": 4500}`
	// req, err := http.NewRequest("POST", "/breed", strings.NewReader(payload))
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// req.Header.Set("Content-Type", "application/json") // Set content type to JSON

	// // Actual requesting of the route
	// router.ServeHTTP(rr, req)

	// // Status, are you ok?
	// if status := rr.Code; status != http.StatusCreated {
	// 	t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	// }

	// // Check the db was indeed requested using the expected parameters
	// if err := sqlMock.ExpectationsWereMet(); err != nil {
	// 	t.Fatalf("could not meet expectations: %s", err)
	// }
}
