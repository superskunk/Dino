package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type animal struct {
	id         int
	animalType string
	nickname   string
	zone       int
	age        int
}

func main() {
	db, err := sql.Open("mysql", "gouser:gouser@/Dino")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//general query with arguments
	rows, err := db.Query("select * from Dino.animals where age > ?", 5)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	animals := []animal{}
	for rows.Next() {
		a := animal{}
		err := rows.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
		if err != nil {
			log.Println(err)
			continue
		}
		animals = append(animals, a)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	log.Println("Results from General Query.....:")
	for _, animal := range animals {
		fmt.Println(animal)
	}

	//query a single row
	row := db.QueryRow("select * from Dino.animals where age > ?", 5)
	a := animal{}
	err = row.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Results from General Query.....:")
	fmt.Println(a)

	/*
		// insert a row
		result, err := db.Exec("Insert into Dino.animals (animal_type, nickname, zone, age) values (?, ?, 3, 22)", "Carnotaurus", "Carno")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("LastInsertId.....: ")
		fmt.Println(result.LastInsertId())
		fmt.Print("RowsAffected.....: ")
		fmt.Println(result.RowsAffected())
	*/

	//update a row
	age := 14
	id := 2
	result, err := db.Exec("Update Dino.animals set age=? where id=?", age, id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("LastInsertId.....: ")
	fmt.Println(result.LastInsertId())
	fmt.Print("RowsAffected.....: ")
	fmt.Println(result.RowsAffected())

}
