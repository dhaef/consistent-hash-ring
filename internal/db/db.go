package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Conn *sql.DB
	Id   string
}

func NewDatabase(id string) Database {
	return Database{
		Id: id,
	}
}

func (db *Database) Start(dbPath string) error {
	conn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	db.Conn = conn

	return nil
}

func (db Database) Stop(dbPath string) error {
	// err := os.Remove(dbPath)
	// if err != nil {
	// 	return err
	// }

	err := db.Conn.Close()
	if err != nil {
		return err
	}

	return nil
}

// func (db Database) Setup() error {
// 	_, err := db.Conn.Exec(`create table data(
//         id text NOT NULL,
//         value text,
//         PRIMARY KEY (id)
//     )`)
// 	if err != nil {
// 		fmt.Printf("error setting up new db node %v", db.Id)
// 		return err
// 	}

// 	return nil
// }
