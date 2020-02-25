package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type animal struct {
	id         int
	animalType string
	nickname   string
	zone       int
	age        int
}

var stdin *bufio.Reader
var actions map[int]string

func init() {

	stdin = bufio.NewReader(os.Stdin)

	actions = map[int]string{
		1: "generalQuery",
		2: "querySingleRow",
		3: "insert",
		4: "update",
		5: "end",
	}
}

func menu() string {
	option := 0
	for option < 1 || option > 5 {
		fmt.Println("1. General Query with parameters")
		fmt.Println("2. Query a Single Row")
		fmt.Println("3. Insert a Row")
		fmt.Println("4. Update a Row")
		fmt.Println("5. Exit")
		fmt.Printf("\nChoose an option....:")
		if _, err := fmt.Fscanf(stdin, "%d", &option); err != nil {
			// In case of not introducing a number
			option = 0
		}
		stdin.ReadLine() //This line is necessary to flush the buffer because there is a "\n" left

	}
	return actions[option]
}

func main() {

	connStr := "user=postgres dbname=dino sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var option string
	for option != "end" {
		option = menu()
		switch {
		case option == "generalQuery":
			rows, err := db.Query("select * from animals where age > $1", 5)
			handlerows(rows, err)
		case option == "querySingleRow":
			a := animal{}
			db.QueryRow("select * from animals where age > $1", 5).Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
			// a := animal{}
			// err = row.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			log.Println("\nResults from Single Query.....:")
			fmt.Println(a)
		case option == "insert":
			result, err := db.Exec("Insert into animals (animal_type, nickname, zone, age) values ('Carnotaurus', 'Carno', $1, $2)", 3, 22)
			if err != nil {
				log.Fatal(err)
			}
			processRows(result, err)
		case option == "update":
			age := 14
			id := 2
			result, err := db.Exec("Update animals set age=$1 where id=$2", age, id)
			processRows(result, err)
		case option == "end":
			return
		}
		fmt.Println("\nPress <ENTER>......")
		stdin.ReadLine()
	}
	os.Exit(0)
}

func handlerows(rows *sql.Rows, err error) {
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
}

func processRows(result sql.Result, err error) {
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("LastInsertId.....: ")
	fmt.Println(result.LastInsertId())
	fmt.Print("RowsAffected.....: ")
	fmt.Println(result.RowsAffected())
}
