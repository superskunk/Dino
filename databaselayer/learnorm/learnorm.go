package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type animal struct {
	ID         int    `gorm:"primary_key;not null;unique;AUTO_INCREMENT"`
	AnimalType string `gorm:"type:TEXT"`
	Nickname   string `gorm:"type:TEXT"`
	Zone       int    `gorm:"type:INTEGER"`
	Age        int    `gorm:"type:INTEGER"`
}

var stdin *bufio.Reader

type options uint

const (
	dropTableIfExits = iota + 1
	createTable
	insertDino
	listDinosOverAnAge
	end
)

func init() {

	stdin = bufio.NewReader(os.Stdin)
}

func menu() options {
	option := 0
	for option < 1 || option > 8 {
		fmt.Println("1. Drop Table If Exists")
		fmt.Println("2. Create Table")
		fmt.Println("3. Insert a Dino")
		fmt.Println("4. List Dinos over an age")
		fmt.Println("5. Exit")
		fmt.Printf("\nChoose an option....:")
		if _, err := fmt.Fscanf(stdin, "%d", &option); err != nil {
			// In case of not introducing a number
			option = 0
		}
		stdin.ReadLine() //This line is necessary to flush the buffer because there is a "\n" left

	}
	return options(option)
}

func main() {
	db, err := gorm.Open("sqlite3", "dino.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	option := menu()
	switch option {
	case dropTableIfExits:
		fmt.Println("Drop Table If Exists....")
		db.Table("dinos").DropTableIfExists(&animal{})
		db.AutoMigrate(&animal{})
	case createTable:
		fmt.Println("Create Table")
		db.CreateTable(&animal{})
	case insertDino:
		fmt.Println("Insert a Dino")
	case listDinosOverAnAge:
		fmt.Println("List Dinos over an age")
	case end:
		fmt.Println("End")
	}
}
