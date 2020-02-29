package databaselayer

import (
	"database/sql"
	"fmt"
	"log"
)

// SQLHandler is the object that abstracts the operations of Dino for any SQL based database. It implements the DinoDBHandler interface.
type SQLHandler struct {
	*sql.DB
}

// GetAvailableDynos returns all of Dinosaurus animals in the DB
func (handler *SQLHandler) GetAvailableDynos() ([]Animal, error) {
	return handler.sendQuery("select * from animals")
}

// GetDynoByNickname returns a Dinosaurus whose nickname matches with the nickname passed as argument
func (handler *SQLHandler) GetDynoByNickname(nickname string) (Animal, error) {
	a := new(Animal)
	err := handler.QueryRow(fmt.Sprintf("select * from animalsl where nickname='%s'", nickname)).Scan(a.ID, a.AnimalType, a.Nickname, a.Zone, a.Age)
	return *a, err
}

// GetDynosByType returns all of Dinosaurus whose animal_type is "dinoType"
func (handler *SQLHandler) GetDynosByType(dinoType string) ([]Animal, error) {
	return handler.sendQuery(fmt.Sprintf("select * from animals where animal_type='%s'", dinoType))
}

// AddAnimal adds animal to the database
func (handler *SQLHandler) AddAnimal(animal Animal) error {
	_, err := handler.Exec(fmt.Sprintf("insert into animals(animal_type, nickname, zone, age) values('%s', '%s', '%d', '%d')",
		animal.AnimalType, animal.Nickname, animal.Zone, animal.Age))
	return err
}

// UpdateAnimal updates the animal identified by nickname with the animal object data
func (handler *SQLHandler) UpdateAnimal(animal Animal, nickname string) error {
	_, err := handler.Exec("update animals set animal_type='%s', nickname='%s', zone='%s', age='%d'", animal.AnimalType, animal.Nickname, animal.Zone, animal.Age)
	return err
}

// sendQuery executes the generic query received as a parameter
func (handler *SQLHandler) sendQuery(q string) ([]Animal, error) {
	rows, err := handler.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	animals := []Animal{}
	for rows.Next() {
		a := new(Animal)
		err := rows.Scan(a.ID, a.AnimalType, a.Nickname, a.Zone, a.Age)
		if err != nil {
			log.Println(err)
			continue
		}
		animals = append(animals, *a)
	}

	return animals, rows.Err()
}
