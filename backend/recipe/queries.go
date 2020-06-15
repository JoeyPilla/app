package recipe

import (
	"context"
	"log"

	"github.com/JoeyPilla/rest-api-example/global"
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

func getAvailableRecipes() []Recipe {
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, `
	SELECT 
		DISTINCT	
			recipe_name, recipe_id
	FROM 
	drink_ingredient
		INNER JOIN 
			ingredient ON (drink_ingredient.fk_ingredient_id = ingredient.ingredient_id)
		INNER JOIN 
			recipe ON (drink_ingredient.fk_recipe_id = recipe.recipe_id
					AND recipe.recipe_id NOT IN
					(
						SELECT
							fk_recipe_id
						FROM
							drink_ingredient
						WHERE drink_ingredient.fk_ingredient_id NOT IN ($1, $2, $3, $4, $5, $6, $7)
					) 
				)   
	`,
		global.MotorMap["motor1"],
		global.MotorMap["motor2"],
		global.MotorMap["motor3"],
		global.MotorMap["motor4"],
		global.MotorMap["motor5"],
		global.MotorMap["motor6"],
		global.MotorMap["motor7"])
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	recipes := make([]Recipe, 0)

	for rows.Next() {
		recipe := Recipe{}
		if err := rows.Scan(
			&recipe.Name,
			&recipe.ID,
		); err != nil {
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
