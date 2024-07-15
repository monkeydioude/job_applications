package internal

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"gorm.io/gorm"
)

// buildBreedsQuery builds a gorm Query out of the URL values.
// `id` has the highest priority over every other filters, they won't even matter.
func buildBreedsQuery(queryParams url.Values, gormDB *gorm.DB) (*gorm.DB, error) {
	if gormDB == nil {
		return nil, fmt.Errorf("gorm is nil omg")
	}

	query := gormDB.Model(&Breed{})
	id := queryParams.Get("id")
	if id != "" {
		return query.Where("id = ?", id), nil
	}
	species := queryParams.Get("species")
	if species != "" {
		query.Where("species = ?", species)
	}
	petSize := queryParams.Get("pet_size")
	if petSize != "" {
		query.Where("pet_size = ?", petSize)
	}
	name := queryParams.Get("name")
	if name != "" {
		query.Where("name = ?", name)
	}
	// this chunk match subfilters for average_male_adult_weight and average_female_adult_weight
	// ex: ?average_female_adult_weight[lte]=4000&average_female_adult_weight[gt]=0
	for key, values := range queryParams {
		if !strings.Contains(key, "average_male_adult_weight") &&
			!strings.Contains(key, "average_female_adult_weight") {
			continue
		}
		for _, value := range values {
			if value == "" {
				continue
			}
			if strings.Contains(key, "[lte]") {
				field := strings.TrimSuffix(key, "[lte]")
				query = query.Where(fmt.Sprintf("%s <= ?", field), value)
			} else if strings.Contains(key, "[gte]") {
				field := strings.TrimSuffix(key, "[gte]")
				query = query.Where(fmt.Sprintf("%s >= ?", field), value)
			} else if strings.Contains(key, "[lt]") {
				field := strings.TrimSuffix(key, "[lt]")
				query = query.Where(fmt.Sprintf("%s < ?", field), value)
			} else if strings.Contains(key, "[gt]") {
				field := strings.TrimSuffix(key, "[gt]")
				query = query.Where(fmt.Sprintf("%s > ?", field), value)
			} else if strings.Contains(key, "[eq]") {
				field := strings.TrimSuffix(key, "[eq]")
				query = query.Where(fmt.Sprintf("%s = ?", field), value)
			} else {
				query = query.Where(fmt.Sprintf("%s = ?", key), value)
			}
		}
	}

	return query, nil
}

// getBreed, through a GET method, handles the modification of a breed.
//
//	Possible filters: `id`, `species`, `name`, `pet_size`,
//	`average_male_adult_weight` (and its subfilters `[eq]`, `[lt]`, `[lte]`, `[gt]`, `[gte]`)
//	and  `average_female_adult_weight` (and its subfilters `[eq]`, `[lt]`, `[lte]`, `[gt]`, `[gte]`).
//
// `id` has the highest priority, other filters won't even matter.
// Ex: _GET /v1/breed?species=dog&average_female_adult_weight[gte]=3000_
func (a *App) getBreed(w http.ResponseWriter, r *http.Request) {
	query, err := buildBreedsQuery(r.URL.Query(), a.gormDB)
	if err != nil {
		a.InternalServerError(w, err)
		return
	}
	var breeds []Breed
	if err := query.Find(&breeds).Error; err != nil {
		a.InternalServerError(w, errors.New("Failed to query breeds"))
		return
	}
	if len(breeds) == 0 {
		a.BadRequest(w, errors.New("Breed not found"))
		return
	}
	a.Ok(w, breeds)
}
