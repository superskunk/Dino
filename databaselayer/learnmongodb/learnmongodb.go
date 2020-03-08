package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type animal struct {
	AnimalType string `bson:"animal_type"`
	Nickname   string `bson:"nickname"`
	Zone       int    `bson:"zone"`
	Age        int    `bson:"age"`
}

var stdin *bufio.Reader

type options uint

const (
	insertList = iota + 1
	update
	remove
	listAll
	listOverAge
	end
)

func init() {
	stdin = bufio.NewReader(os.Stdin)
}

func main() {
	session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	animalCollection := session.DB("Dino").C("animals")
exit:
	for {
		option := menu()
		switch option {
		case insertList:
			if insertDinos(animalCollection) != nil {
				log.Fatal(err)
			}
		case update:
			if updateDino(animalCollection) != nil {
				log.Fatal(err)
			}
		case remove:
			if removeDino(animalCollection) != nil {
				log.Fatal(err)
			}
		case listAll:
			var dinos []animal
			var err error
			if dinos, err = listAllDinos(animalCollection); err != nil {
				log.Fatal(err)
			}
			fmt.Println("List of the whole dinos...: ")
			printDinos(dinos)
		case listOverAge:
			var dinos []animal
			var err error
			if dinos, err = listDinosOverAge(animalCollection); err != nil {
				log.Fatal(err)
			}
			fmt.Println("List of dinos over age founded...: ")
			printDinos(dinos)
		case end:
			break exit
		}
	}
	os.Exit(0)
}

func menu() options {
	option := 0
	for option < 1 || option > 6 {
		fmt.Println("1. Insert Dinos")
		fmt.Println("2. Update Dino")
		fmt.Println("3. Remove Dino")
		fmt.Println("4. List All Dinos")
		fmt.Println("5. List Dinos older than five and in zone 1 or 2")
		fmt.Println("6. Exit")
		fmt.Printf("\nChoose an option....:")
		if _, err := fmt.Fscanf(stdin, "%d", &option); err != nil {
			// In case of not introducing a number
			option = 0
		}
		stdin.ReadLine() //This line is necessary to flush the buffer because there is a "\n" left

	}
	return options(option)
}

func insertDinos(animals *mgo.Collection) error {
	a := []interface{}{
		animal{
			AnimalType: "Tyrannosaurus rex",
			Nickname:   "rex",
			Zone:       1,
			Age:        11,
		},
		animal{
			AnimalType: "Velociraptor",
			Nickname:   "rapto",
			Zone:       2,
			Age:        17,
		},
		animal{
			AnimalType: "Velociraptor",
			Nickname:   "Velo",
			Zone:       2,
			Age:        9,
		},
	}
	return animals.Insert(a...)
}

func updateDino(animals *mgo.Collection) error {
	var name string
	var age int
	fmt.Printf("\nGive the Dino's name to update: ")
	fmt.Fscanf(stdin, "%s", &name)
	stdin.ReadLine() //This line is necessary to flush the buffer because there is a "\n" left
	fmt.Printf("\nGive the new Age: ")
	fmt.Fscanf(stdin, "%d", &age)
	stdin.ReadLine()

	return animals.Update(bson.M{"nickname": name}, bson.M{"$set": bson.M{"age": age}})
}

func removeDino(animals *mgo.Collection) error {
	var name string
	fmt.Printf("\nGive the Dino's name to remove: ")
	fmt.Fscanf(stdin, "%s", &name)
	stdin.ReadLine() //This line is necessary to flush the buffer because there is a "\n" left

	return animals.Remove(bson.M{"nickname": name})
}

func listDinosOverAge(animals *mgo.Collection) ([]animal, error) {
	var zone1, zone2 int
	var age int
	fmt.Printf("\nGive the Dino's age to find: ")
	fmt.Printf("\nGive the new Age: ")
	fmt.Fscanf(stdin, "%d", &age)
	stdin.ReadLine()
	fmt.Printf("\nGive zone 1 and 2: ")
	fmt.Fscanf(stdin, "%d%d", &zone1, &zone2)
	stdin.ReadLine()

	query := bson.M{
		"age": bson.M{
			"$gt": 5,
		},
		"zone": bson.M{
			"$in": []int{zone1, zone2},
		},
	}
	result := []animal{}

	err := animals.Find(query).All(&result)

	return result, err
}

func listAllDinos(animals *mgo.Collection) ([]animal, error) {
	result := []animal{}
	err := animals.Find(bson.M{}).All(&result)
	return result, err
}

func printDinos(dinos []animal) {
	for _, dino := range dinos {
		fmt.Println(dino)
	}
	fmt.Println()
}
