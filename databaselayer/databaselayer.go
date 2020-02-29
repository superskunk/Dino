package databaselayer

import "errors"

// DBType defines an enumeration to keep the diffeent db types
type DBType uint8

const (
	// MYSQL constant indicates to the GetDatabaseHandler method to execute the database MySql constructor
	MYSQL DBType = iota
	// SQLITE constant indicates to the GetDatabaseHandler method to execute the database SqLite constructor
	SQLITE
	// POSTGRESS constant indicates to the GetDatabaseHandler method to execute the database Postgress constructor
	POSTGRESS
	// MONGODB constant indicates to the GetDatabaseHandler method to execute the database MongoDB constructor
	MONGODB
)

// Animal is the struct that defines what an animal is
type Animal struct {
	ID         int    `bson:"-"`
	AnimalType string `bson:"animal_type"`
	Nickname   string `bson:"nickname"`
	Zone       int    `bson:"zone"`
	Age        int    `bson:"age"`
}

// DinoDBHandler defines the interface that every databaseHandler has to fullfill
type DinoDBHandler interface {
	GetAvaialbeDynos() ([]Animal, error)
	GeDynoByNickName(string) (Animal, error)
	GetDynosByType(string) ([]Animal, error)
	AddAnimal(Animal) error
	UpdateAnimal(Animal, string) error
}

// DBTypeNotSupported is an error Object that will be returned when using a database type not supported
var DBTypeNotSupported = errors.New("The Database type provided is not supported...")

// GetDatabaseHandler is a factory function to create one or another database handler according to the dbtype parameter
func GetDatabaseHandler(dbtype DBType, connection string) (DinoDBHandler, error) {

	switch dbtype {
	case MYSQL:
	case SQLITE:
	case POSTGRESS:
	case MONGODB:
	}

	return nil, DBTypeNotSupported
}
