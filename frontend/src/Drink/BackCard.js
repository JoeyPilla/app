
import React from 'react';
import './BackCard.css'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faTimesCircle } from '@fortawesome/free-solid-svg-icons'
import {useNavigate} from "@reach/router"
export default function BackCard({ recipe, ingredients, setFlip }) {
  const navigate = useNavigate()
  return (
    <div className='back-card-container'>
    <div className="back-card-ingredient" onClick={()=> navigate(`/drink/${recipe.id}`)}>
      <h2>{recipe?.name}</h2>
      <FontAwesomeIcon
            className="back-card-icon"
        icon={faTimesCircle}
        onClick={() => {
          fetch(`/api/recipe?id=${recipe?.id}`, { method: 'DELETE' })
        }}
      />
      </div>
      {ingredients}
      <div className="back-card-button-container">
      <button
            type="button"
            className="back-card-form-button-back"
            onClick={() => {
              setFlip()
            }}
          >Back</button>
        <button type="button" className="back-card-form-button" onClick={() => navigate(`/drink/${recipe.id}`)}>Edit</button>
      </div>

    </div>
  )
}


