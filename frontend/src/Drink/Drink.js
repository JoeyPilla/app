import React, { useEffect, useState } from 'react';
import DrinkListElement from './DrinkListElement'
import AddIngredient from './AddIngredient'
import Pour from './Pour'

export default function Drink({id}) {
  const [updated, setUpdated] = useState(false)
  const [ingredients, setIngredients] = useState([])
  const [possibleIngredients, setPossibleIngredients] = useState([])
  const [recipe, setRecipe] = useState({})
  useEffect(() => {
    const fetchRecipe = async () => {
      const ingredients = await fetch(`/api/recipe?id=${id}`)
      const json = await ingredients.json()
      setRecipe(json)
      console.log(json)
    }
    const fetchIngredients = async () => {
      const ingredients = await fetch(`/api/drink?recipeId=${id}`)
      const json = await ingredients.json()
      setIngredients(json)
    }
    fetchIngredients()
    fetchRecipe()
  }, [id])
  useEffect(() => {
    const fetchIngredients = async () => {
      const data = await fetch('/api/ingredient/all')
      const json = await data.json()
      setPossibleIngredients(json)
    }
    fetchIngredients()
  }, [])
  const ingred = ingredients.map((i) => <h3>{i.measure} {i.unitOfMeasurement} {i.ingredient.name}</h3>)
  return (
    <div className="ingredients">
      <h1>
      {recipe.name}
      </h1>
      {ingred}
      <Pour recipeId={recipe.id} />
      <AddIngredient updated={updated} setUpdated={setUpdated} recipeId={id}/>
    </div>
  )
}

