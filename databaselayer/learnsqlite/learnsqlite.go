package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
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
		5: "preparedStatement",
		6: "testTransaction",
		7: "insertTransaction",
		8: "end",
	}
}

func menu() string {
	option := 0
	for option < 1 || option > 8 {
		fmt.Println("1. General Query with parameters")
		fmt.Println("2. Query a Single Row")
		fmt.Println("3. Insert a Row")
		fmt.Println("4. Update a Row")
		fmt.Println("5. Prepared Statement")
		fmt.Println("6. testTransactions")
		fmt.Println("7. Insert Transaction")
		fmt.Println("8. Exit")
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
	log.Println("Connect to database...")
	db, err := sql.Open("sqlite3", "dino.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := createDatabase(db); err != nil {
		log.Fatal(err)
	}

	if err := insertSampleData(db); err != nil {
		log.Fatal(err)
	}

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
		case option == "preparedStatement":
			stmt, err := db.Prepare("select * from animals where age > $1")
			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()
			rows, err := stmt.Query(5)
			handlerows(rows, err)
		case option == "testTransaction":
			testTransaction(db)
		case option == "insertTransaction":
			a := &animal{0, "AlloSaurus", "Allo", 3, 25}
			result, err := insertAnimal(db, a)
			processRows(result, err)
		case option == "end":
			return
		}
		fmt.Println("\nPress <ENTER>......")
		stdin.ReadLine()
	}
	os.Exit(0)

}

func createDatabase(db *sql.DB) error {
	_, err := db.Exec(` CREATE TABLE IF NOT EXISTS
	animals(id INTEGER PRIMARY KEY AUTOINCREMENT,
		animal_type TEXT,
		nickname TEXT,
		zone INTEGER,
		age INTEGER)`)
	return err
}

func insertSampleData(db *sql.DB) error {
	_, err := db.Exec(`
	INSERT INTO animals(animal_type, nickname, zone, age)
	VALUES('Tyrannosaurus rex', 'rex', 1, 10),
	('Velociraptor', 'rapto', 2, 15)
	`)
	return err
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

func testTransaction(db *sql.DB) {
	fmt.Println("Transactions...")
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("select * from animals where age > $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(6)
	handlerows(rows, err)
	rows, err = stmt.Query(2)
	handlerows(rows, err)
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

}

func insertAnimal(db *sql.DB, a *animal) (sql.Result, error) {
	fmt.Printf("Starting transaction for inserting animal: %v\n", a)
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("insert into animals(animal_type, nickname, zone, age) values($1,$2,$3,$4)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(a.animalType, a.nickname, a.zone, a.age)
	if err != nil {
		fmt.Printf("Rollback transaction for inserting animal: %v\n", a)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return result, err
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
