import React, { useEffect, useState } from 'react';
import './DrinkListElement.css'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faTimesCircle } from '@fortawesome/free-solid-svg-icons'
import { navigate } from '@reach/router';

export default function DrinkListElement({ recipe, updated, setUpdated }) {
  return (
    <div className="ingredient" onClick={()=> navigate(`/drink/${recipe.id}`)}>
      {recipe?.name}
      <FontAwesomeIcon
            className="times-circle-icon"
        icon={faTimesCircle}
        onClick={() => {
          fetch(`/api/recipe?id=${recipe?.id}`, { method: 'DELETE' })
            .then(() => setUpdated(!updated))
        }}
      />
    </div>
  )
}

