import React, { useEffect, useState } from 'react';
import './Ingredient.css'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import {  faTimesCircle } from '@fortawesome/free-solid-svg-icons'
export default function Ingredient({ingredient, updated, setUpdated}) {
  return (
    <div className="ingredient">
      {ingredient.name}
      <FontAwesomeIcon
            className="times-circle-icon"
        icon={faTimesCircle}
        onClick={() => {
          fetch(`/api/ingredient?id=${ingredient.id}`, { method: 'DELETE' }).then(() => setUpdated(!updated))
        }}
          />
    </div>
  )
}

