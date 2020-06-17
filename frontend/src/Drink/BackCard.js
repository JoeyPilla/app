
import React, { useEffect, useState } from 'react';
import './BackCard.css'
import { useSpring, animated as a } from 'react-spring'
import Pour from './Pour'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faTimesCircle } from '@fortawesome/free-solid-svg-icons'
import {navigate} from "@reach/router"
export default function BackCard({ recipe, ingredients}) {
  return (
    <div className='back-card-container'>
    <div className="back-card-ingredient" onClick={()=> navigate(`/drink/${recipe.id}`)}>
      {recipe?.name}
      <FontAwesomeIcon
            className="back-card-icon"
        icon={faTimesCircle}
        onClick={() => {
          fetch(`/api/recipe?id=${recipe?.id}`, { method: 'DELETE' })
        }}
      />
      </div>
      {ingredients}
    </div>
  )
}


