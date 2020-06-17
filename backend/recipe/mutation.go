package recipe

import (
	"context"
	"fmt"
	"log"
)

func addRecipe(name string) Recipe {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	result, err := tx.ExecContext(ctx, "INSERT INTO recipe (recipe_name) VALUES ($1)", name)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		log.Fatal(err)
		tx.Rollback()
		return Recipe{}
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rows != 1 {
		log.Fatalf("expected to affect 1 row, affected %d", rows)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return getRecipeByName(name)
}

func deleteRecipe(id int) int64 {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	result, err := tx.ExecContext(ctx, `
	DELETE FROM recipe 
	WHERE recipe_id=$1`, id)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		log.Fatal(err)
		tx.Rollback()
		return 0
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rows != 1 {
		log.Fatalf("expected to affect 1 row, affected %d", rows)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return rows
}
