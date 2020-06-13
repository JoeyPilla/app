package api

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/JoeyPilla/rest-api-example/drink"
	"github.com/JoeyPilla/rest-api-example/ingredient"
	"github.com/JoeyPilla/rest-api-example/recipe"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "me"
	password = "1234"
	dbname   = "bartenderDB"
)

var tables []string = []string{
	"ingredient",
	"drink_ingredient",
	"recipe",
}

func Connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	err = InitalizeDB(db)
	if err != nil {
		panic(err)
	}
	ingredient.Initialize(db)
	recipe.Initialize(db)
	drink.Initialize(db)
	return db
}

func ExampleDB_QueryContext(db *sql.DB) []string {
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, `SELECT table_name
		FROM information_schema.tables
	 WHERE table_schema='public'
		 AND table_type='BASE TABLE'`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	tables := make([]string, 0)

	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		tables = append(tables, table)
	}
	// If the database is being written to ensure to check for Close
	// errors that may be returned from the driver. The query may
	// encounter an auto-commit error and be forced to rollback changes.
	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(err)
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return tables
}

func InitalizeDB(db *sql.DB) error {
	currentTables := ExampleDB_QueryContext(db)
	for _, table := range tables {
		if !contains(currentTables, table) {
			err := createTable(table, db)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func createTable(table string, db *sql.DB) error {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	switch table {
	case "ingredient":
		_, err = tx.ExecContext(ctx, `CREATE TABLE ingredient (
				ingredient_id integer PRIMARY KEY,
				ingredient_name TEXT
				)`)
		if err != nil {
			tx.Rollback()
			return err
		}
	case "drink_ingredient":
		_, err = tx.ExecContext(ctx, `CREATE TABLE drink_ingredient (
				fk_ingredient_id integer REFERENCES ingredient (ingredient_id),
				fk_recipe_id integer REFERENCES recipe (recipe_id),
				measure decimal,
				unit_of_measurement TEXT
				)`)
		if err != nil {
			tx.Rollback()
			return err
		}
	case "recipe":
		_, err = tx.ExecContext(ctx, `CREATE TABLE recipe (
				recipe_id integer PRIMARY KEY,
				recipe_name TEXT
				)`)
		if err != nil {
			tx.Rollback()
			return err
		}
	default:
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully initalized", table)
	return err
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
