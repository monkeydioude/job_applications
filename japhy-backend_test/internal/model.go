package internal

type Breed struct {
	ID                       *int   `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`
	Species                  string `json:"species" gorm:"type:longtext"`
	PetSize                  string `json:"pet_size" gorm:"type:varchar(100)"`
	Name                     string `json:"name" gorm:"type:varchar(255);unique"`
	AverageMaleAdultWeight   int    `json:"average_male_adult_weight"`
	AverageFemaleAdultWeight int    `json:"average_female_adult_weight"`
}
