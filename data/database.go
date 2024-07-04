package data

import (
	"Gokedex/utils"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

var dbFileName string

func InitDB() {
	dbFileName, err := utils.GetSetting("DatabaseName")

	if err != nil {
		panic("Could not get DatabaseName. Did you set it in settings.json?")
	}

	db, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		panic("Could not connect to the database.")
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(5)

	DB = db

	setuptables()
	seedData()
}

func CloseDB() error {
	return DB.Close()
}

func setuptables() {
	usersTable := `
  CREATE TABLE IF NOT EXISTS "Users" (
	  "Id"	INTEGER NOT NULL UNIQUE,
	  "Username"	TEXT NOT NULL UNIQUE,
	  "Email"	TEXT UNIQUE,
	  "Password"	TEXT,
	  PRIMARY KEY("Id" AUTOINCREMENT)
  );`

	_, err := DB.Exec(usersTable)
	if err != nil {
		fmt.Println("Could not create Users table!")
	}

	pokemonTable := `
  CREATE TABLE IF NOT EXISTS "Pokemon" (
	  "Id"	INTEGER NOT NULL UNIQUE,
	  "Name"	TEXT,
	  "Category"	TEXT,
	  "Height"	INTEGER,
	  "Weight"	INTEGER,
	  "Type"	TEXT,
	  "Weakness"	TEXT,
	  "Abilities"	TEXT,
	  PRIMARY KEY("Id" AUTOINCREMENT)
  );`

	_, err = DB.Exec(pokemonTable)
	if err != nil {
		fmt.Println("Could not create Pokemon table!")
	}
}

func seedData() {
	insertQuery := `
	INSERT INTO Users(Username, Email, Password)
	VALUES(?,?,?)
	`
	hashedPassword, err := utils.HashPassword("GokedexAdmin")

	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = DB.Exec(insertQuery, "admin", "admin@gokedex.com", hashedPassword)

	if err != nil {
		fmt.Println(err.Error())
	}
}
