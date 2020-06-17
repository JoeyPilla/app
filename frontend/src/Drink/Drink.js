import React, { useEffect, useState } from 'react';
import AddIngredient from './AddIngredient'
import Pour from './Pour'
import './Drink.css'
import DrinkIngredient from './DrinkIngredient'
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
  }, [id, updated])
  useEffect(() => {
    const fetchIngredients = async () => {
      const data = await fetch('/api/ingredient/all')
      const json = await data.json()
      setPossibleIngredients(json)
    }
    fetchIngredients()
  }, [])
  const ingred = ingredients.map((i) => <DrinkIngredient ingredient={i} updated={updated} setUpdated={setUpdated}/>)
  return (
    <div className="single-drink-container">
      <h1>
      {recipe.name}
      </h1>
      {ingred}
      <AddIngredient updated={updated} setUpdated={setUpdated} recipeId={id}/>
    </div>
  )
}

