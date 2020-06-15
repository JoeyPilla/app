package ingredient

import (
	"context"
	"log"
)

func getAllIngredients() []Ingredient {
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, `SELECT * FROM ingredient ORDER BY ingredient_name ASC`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	ingredients := make([]Ingredient, 0)

	for rows.Next() {
		ingredient := Ingredient{}
		if err := rows.Scan(&ingredient.ID, &ingredient.Name); err != nil {
			log.Fatal(err)
		}
		ingredients = append(ingredients, ingredient)
	}

	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(err)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return ingredients
}

func getIngredientById(id int) Ingredient {
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, `
	SELECT 
		* 
	FROM ingredient 
	WHERE ingredient_id=$1`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	ingredient := Ingredient{}

	for rows.Next() {
		if err := rows.Scan(&ingredient.ID, &ingredient.Name); err != nil {
			log.Fatal(err)
		}
	}
	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(err)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return ingredient
}

func getIngredientByName(name string) Ingredient {
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, `
	SELECT 
		* 
	FROM ingredient 
	WHERE ingredient_name=$1`, name)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	ingredient := Ingredient{}

	for rows.Next() {
		if err := rows.Scan(&ingredient.ID, &ingredient.Name); err != nil {
			log.Fatal(err)
		}
	}
	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(err)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return ingredient
}
