package drink

import (
	"context"
	"log"

	"github.com/JoeyPilla/rest-api-example/global"
)

func getAllDrinks() []Drink {
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, `
	SELECT 
		recipe_name, recipe_id, ingredient_name, ingredient_id, measure, unit_of_measurement 
	FROM 
		drink_ingredient
		INNER JOIN 
			ingredient ON (drink_ingredient.fk_ingredient_id = ingredient.ingredient_id)
		INNER JOIN 
			recipe ON (drink_ingredient.fk_recipe_id = recipe.recipe_id)
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	drinks := make([]Drink, 0)

	for rows.Next() {
		drink := Drink{}
		if err := rows.Scan(
			&drink.Recipe.Name,
			&drink.Recipe.ID,
			&drink.Ingredient.Name,
			&drink.Ingredient.ID,
			&drink.Measure,
			&drink.UnitOfMeasurement,
		); err != nil {
			log.Fatal(err)
		}
		drinks = append(drinks, drink)
	}

	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(err)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return drinks
}

func getAvailableDrinks() []Drink {
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, `
	SELECT 
		recipe_name, recipe_id, ingredient_name, ingredient_id, measure, unit_of_measurement 
	FROM 
		drink_ingredient
		INNER JOIN 
			ingredient ON (drink_ingredient.fk_ingredient_id = ingredient.ingredient_id)
		INNER JOIN 
			recipe ON (drink_ingredient.fk_recipe_id = recipe.recipe_id
					AND recipe.recipe_id IN
					(
						SELECT
							fk_recipe_id
						FROM
							drink_ingredient
						WHERE drink_ingredient.fk_ingredient_id IN ($1, $2, $3, $4, $5, $6, $7)
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
	drinks := make([]Drink, 0)

	for rows.Next() {
		drink := Drink{}
		if err := rows.Scan(
			&drink.Recipe.Name,
			&drink.Recipe.ID,
			&drink.Ingredient.Name,
			&drink.Ingredient.ID,
			&drink.Measure,
			&drink.UnitOfMeasurement,
		); err != nil {
			log.Fatal(err)
		}
		drinks = append(drinks, drink)
	}

	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(err)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return drinks
}

func getDrinkByRecipe(id int) []Drink {
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, `
	SELECT 
		recipe_name, recipe_id, ingredient_name, ingredient_id, measure, unit_of_measurement 
	FROM 
		drink_ingredient
		INNER JOIN 
			ingredient ON (drink_ingredient.fk_ingredient_id = ingredient.ingredient_id)
		INNER JOIN 
			recipe ON (drink_ingredient.fk_recipe_id = recipe.recipe_id)
	WHERE recipe.recipe_id=$1
	`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	drinks := make([]Drink, 0)

	for rows.Next() {
		drink := Drink{}
		if err := rows.Scan(
			&drink.Recipe.Name,
			&drink.Recipe.ID,
			&drink.Ingredient.Name,
			&drink.Ingredient.ID,
			&drink.Measure,
			&drink.UnitOfMeasurement,
		); err != nil {
			log.Fatal(err)
		}
		drinks = append(drinks, drink)
	}

	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(err)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return drinks
}

func getDrinkByIngredient(id int) []Drink {
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, `
	SELECT 
		recipe_name, recipe_id, ingredient_name, ingredient_id, measure, unit_of_measurement 
	FROM 
		drink_ingredient
		INNER JOIN 
			ingredient ON (drink_ingredient.fk_ingredient_id = ingredient.ingredient_id)
		INNER JOIN 
			recipe ON (drink_ingredient.fk_recipe_id = recipe.recipe_id)
	WHERE ingredient.ingredient_id=$1
	`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	drinks := make([]Drink, 0)

	for rows.Next() {
		drink := Drink{}
		if err := rows.Scan(
			&drink.Recipe.Name,
			&drink.Recipe.ID,
			&drink.Ingredient.Name,
			&drink.Ingredient.ID,
			&drink.Measure,
			&drink.UnitOfMeasurement,
		); err != nil {
			log.Fatal(err)
		}
		drinks = append(drinks, drink)
	}

	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(err)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return drinks
}
