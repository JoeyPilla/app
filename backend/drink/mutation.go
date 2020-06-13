package drink

import (
	"context"
	"log"
)

func addDrink(ingredientId int, recipeId int, measure float64, unitOfMeasurement string) int64 {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	result, err := tx.ExecContext(ctx, `
	INSERT INTO drink_ingredient  
		(fk_ingredient_id, fk_recipe_id, measure, unit_of_measurement) 
	VALUES ($1, $2, $3, $4)`, ingredientId, recipeId, measure, unitOfMeasurement)
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

func deleteDrink(ingredientId int, recipeId int) int64 {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	result, err := tx.ExecContext(ctx, `
	DELETE FROM drink_ingredient  
	WHERE 
	fk_ingredient_id=$1
	AND fk_recipe_id=$2`, ingredientId, recipeId)
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

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return rows
}
