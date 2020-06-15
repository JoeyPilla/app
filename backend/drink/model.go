package drink

import (
	"github.com/JoeyPilla/rest-api-example/ingredient"
	"github.com/JoeyPilla/rest-api-example/recipe"
)

type Drink struct {
	Ingredient        ingredient.Ingredient `json:"ingredient"`
	Recipe            recipe.Recipe         `json:"recipe"`
	Measure           float64               `json:"measure"`
	UnitOfMeasurement string                `json:"unitOfMeasurement"`
}
