package main

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/jackc/pgx/v5/pgxpool"
)

type dataBaseEntity struct {
	id           int
	name         string
	armorType    string
	defense      int
	fireRes      int
	waterRes     int
	lightningRes int
	skill1       string
	skill2       string
	skill3       string
	gemSlot1     int
	gemSlot2     int
	gemSlot3     int
}

func main() {
	fmt.Println("hello world")

	// Define connection string
	databaseName := "postgres://anthonycalcagno:secret@localhost:5432/test"

	// Create a connection pool
	pool, err := pgxpool.New(context.Background(), databaseName)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer pool.Close()

	fmt.Println("Connected to PostgreSQL!")

	rows, err := pool.Query(context.Background(), "SELECT * FROM armor")
	if err != nil {
		log.Fatalf("Query error: %v", err)
	}
	defer rows.Close()

	fmt.Printf("printing row = %v\n", rows)

	dbe := dataBaseEntity{}

	for rows.Next() {
		rows.Scan(&dbe.id, &dbe.name, &dbe.armorType, &dbe.defense, &dbe.fireRes, &dbe.waterRes, &dbe.lightningRes, &dbe.skill1, &dbe.skill2, &dbe.skill3, &dbe.gemSlot1, &dbe.gemSlot2, &dbe.gemSlot3)
		fmt.Printf("entity values after scan = %v\n", dbe)
	}

	values := reflect.ValueOf(dbe)
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		fmt.Println(types.Field(i).Index[0], types.Field(i).Name, values.Field(i))
	}

}
