import React, { useEffect, useState } from 'react';
import DrinkListElement from './DrinkListElement'
import AddDrink from './AddDrink'

export default function Drinks() {
  const [updated, setUpdated] = useState(false)
  const [recipes, setRecipes] = useState([])
  useEffect(() => {
    const fetchData = async () => {
      const data = await fetch('/api/recipe/available')
      const json = await data.json()
      setRecipes(json)
    }
    fetchData()
  }, [])
  const availableRecipes = recipes.map((recipe) => <DrinkListElement
    recipe={recipe}
    updated={updated}
    setUpdated={setUpdated} 
  />
  )
  return (
    <div className="ingredients">
      {availableRecipes}
      <AddDrink updated={updated} setUpdated={setUpdated}/>
    </div>
  )
}

