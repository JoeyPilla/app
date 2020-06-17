import React, { useEffect, useState } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import {  faTimesCircle } from '@fortawesome/free-solid-svg-icons'
export default function DrinkIngredient({ ingredient, updated, setUpdated }) {
  return (
    <div className="ingredient">
      {ingredient.measure} {ingredient.unitOfMeasurement} {ingredient.ingredient.name}
      <FontAwesomeIcon
            className="times-circle-icon"
        icon={faTimesCircle}
        onClick={() => {
          if (window.confirm(`Are you sure you want to delete ${ingredient.measure} ${ingredient.unitOfMeasurement} ${ingredient.ingredient.name}?`)) {
            fetch(`/api/drink?ingredientId=${ingredient.ingredient.id}&recipeId=${ingredient.recipe.id}`, { method: 'DELETE' }).then(() => setUpdated(!updated))
           }
        }}
          />
    </div>
  )
}
