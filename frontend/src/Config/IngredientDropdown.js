import React, {useEffect, useState} from 'react';
import './IngredientDropdown.css';

export default function IngredientDropdown({item, value, setValue, options}) {
  return (
    <div className="ingredient-dropdown">
        Ingredient {item}:
        <select value={value} onChange={e => setValue(e.target.value)}>
          <option value={0}>Select an ingredient...</option>
          {options}
        </select>
    </div>
  )
}
