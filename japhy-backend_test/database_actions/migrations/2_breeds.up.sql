LOAD DATA INFILE '/breeds.csv'
INTO TABLE core.breeds
FIELDS TERMINATED BY ','
ENCLOSED BY '"'
LINES TERMINATED BY '\n'
IGNORE 1 ROWS
(id, species, pet_size, name, average_male_adult_weight, average_female_adult_weight);
