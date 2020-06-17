import React, {useEffect, useState} from 'react';
import './IngredientDropdown.css';

export default function IngredientDropdown({item, value, setValue, options}) {
  return (
    <div className="ingredient-dropdown">
      {item}:
        <select value={value} onChange={e => setValue(e.target.value)} className='select-box'>
        <option value={0}>Select ingredient {item}</option>
          {options}
        </select>
    </div>
  )
}
