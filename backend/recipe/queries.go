package recipe

import (
	"context"
	"log"
)

func getAllRecipes() []Recipe {
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, `SELECT * FROM recipe`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	recipes := make([]Recipe, 0)

	for rows.Next() {
		recipe := Recipe{}
		if err := rows.Scan(&recipe.ID, &recipe.Name); err != nil {
			log.Fatal(err)
		}
		recipes = append(recipes, recipe)
	}

	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(err)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return recipes
}

func getRecipeById(id int) Recipe {
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, `
	SELECT 
		* 
	FROM recipe 
	WHERE recipe_id=$1`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	recipe := Recipe{}

	for rows.Next() {
		if err := rows.Scan(&recipe.ID, &recipe.Name); err != nil {
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
	return recipe
}

func getRecipeByName(name string) Recipe {
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, `
	SELECT 
		* 
	FROM recipe 
	WHERE recipe_name=$1`, name)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	recipe := Recipe{}

	for rows.Next() {
		if err := rows.Scan(&recipe.ID, &recipe.Name); err != nil {
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
	return recipe
}
