package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-adodb"
)

var (
	imie     string
	nazwisko string
)

func main() {

	workingPath, _ := os.Getwd()

	dbPath := fmt.Sprintf(`Provider=Microsoft.ACE.OLEDB.12.0;Data Source=%s/example.accdb`, workingPath)

	db, err := sql.Open("adodb", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()

	if err != nil {
		fmt.Println("No connection")
	}

	effect, err := db.Query("SELECT imie, nazwisko FROM Tabelka")

	if err != nil {
		log.Fatal(err)
	}
	defer effect.Close()

	for effect.Next() {

		err := effect.Scan(&imie, &nazwisko)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(imie, nazwisko)

	}

	err = effect.Err()
	if err != nil {
		log.Fatal(err)
	}
}
