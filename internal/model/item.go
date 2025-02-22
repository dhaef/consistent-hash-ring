package model

import (
	"database/sql"
	"fmt"
)

type Item struct {
	Id    string
	Value string
}

func CreateItemsTable(conn *sql.DB) error {
	_, err := conn.Exec(`create table items(
        id text NOT NULL,
        value text,
        PRIMARY KEY (id)
    )`)
	if err != nil {
		fmt.Println("error setting up table", err.Error())
		return err
	}

	return nil
}

func Get(conn *sql.DB, id string) (Item, error) {
	var item Item

	if err := conn.QueryRow("SELECT * from items where id = ?",
		id).Scan(&item.Id, &item.Value); err != nil {
		if err == sql.ErrNoRows {
			return item, nil
		}
		return item, err
	}

	return item, nil
}

func Create(conn *sql.DB, id string, value string) (Item, error) {
	item := Item{
		Id:    id,
		Value: value,
	}

	_, err := conn.Exec("INSERT INTO items VALUES(?,?)", item.Id, item.Value)
	if err != nil {
		return item, err
	}

	return item, nil
}
