import React, { useState} from 'react';
import './App.css';

export default function Add() {
  const [ingredient, setIngredient] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault()
    fetch(`/api/ingredient?name=${encodeURI(ingredient)}`, {
      method: 'POST'
    })
    setIngredient('')
  }

  return (
    <form onSubmit={(e) => handleSubmit(e)}>
        <label>
        Ingredient Name:
        <input
          type="text"
          name="name"
          onChange={(e) => setIngredient(e.target.value)}
        />
      </label>
    </form>
    )
}
