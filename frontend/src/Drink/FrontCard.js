
import React, { useEffect, useState } from 'react';
import './FrontCard.css'
import { useSpring, animated as a } from 'react-spring'
import Pour from './Pour'

export default function FrontCard({ recipe, setFlip}) {
  return (
    <div className='front-card-container'>
      <h1 className='front-card-cotainer-title'>{recipe?.name}</h1>
      <p onClick={(e) => {
        setFlip()
      }}>See ingredients</p>
      <Pour recipeId={recipe.id} />
    </div>
  )
}


