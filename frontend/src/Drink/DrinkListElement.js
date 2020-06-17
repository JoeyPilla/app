import React, { useEffect, useState } from 'react';
import './DrinkListElement.css'
import { useSpring, animated as a } from 'react-spring'
import FrontCard from './FrontCard'
import BackCard from './BackCard';

export default function DrinkListElement({ recipe, updated, setUpdated }) {
  const [flipped, setFlipped] = useState(false)
  const [ingredients, setIngredients] = useState([])
  const [possibleIngredients, setPossibleIngredients] = useState([])
  const { transform, opacity } = useSpring({
    opacity: flipped ? 1 : 0,
    transform: `perspective(600px) rotateX(${flipped ? 180 : 0}deg)`,
    config: { mass: 5, tension: 500, friction: 80 }
  })

  useEffect(() => {
    const fetchIngredients = async () => {
      const ingredients = await fetch(`/api/drink?recipeId=${recipe.id}`)
      const json = await ingredients.json()
      setIngredients(json)
    }
    fetchIngredients()
  }, [recipe])
  useEffect(() => {
    const fetchIngredients = async () => {
      const data = await fetch('/api/ingredient/all')
      const json = await data.json()
      setPossibleIngredients(json)
    }
    fetchIngredients()
  }, [])
  const ingred = ingredients.map((i) => <p key={i.ingredient_id}>{i.measure} {i.unitOfMeasurement} {i.ingredient.name}</p>)

  
  return (
    <div
      className="drink-list-element-container">
      {!flipped && <a.div className="c front" style={{ opacity: opacity.interpolate(o => 1 - o), transform }}>
        <FrontCard recipe={recipe} setFlip={(e) => { setFlipped(state => !state) }}/>
      </a.div>
      }
      {flipped && <a.div className="c back" style={{ opacity, transform: transform.interpolate(t => `${t} rotateX(180deg)`) }}>
        <BackCard recipe={recipe} ingredients={ingred} setFlip={(e) => { setFlipped(state => !state) }}/>
      </a.div>}
</div>
  )
}

