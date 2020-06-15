import React, { useEffect, useState } from 'react';
import Add from './Add'
import Ingredient from './Ingredient'
import './Ingredients.css'
export default function AddDrink() {
  const [ingredients, setIngredients] = useState([])
  const [updated, setUpdated] = useState(false)

  useEffect(() => {
    const fetchData = async () => {
      const data = await fetch('/api/ingredient/all')
      const json = await data.json()
      setIngredients(json)
    }
    fetchData()
  }, [updated])
  const ingredients2 = ingredients.map(ingredient => <Ingredient
    updated={updated}
    setUpdated={setUpdated}
    ingredient={ingredient} />)
  return (
    <div className="ingredients">
      {ingredients2}
      <Add updated={updated} setUpdated={setUpdated}/>
    </div>
  )
}

