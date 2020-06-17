import React, { useState, useRef} from 'react';
import './AddForm.css';

export default function AddForm({setRecentlyAdded, updated, setUpdated, setShow}) {
  const [ingredient, setIngredient] = useState('');
  const ref = useRef(null)

  React.useEffect(() => {
    ref.current.focus()
  }, [])

  const handleSubmit = (e) => {
  e.preventDefault()
  if (ingredient.length > 0) {
    fetch(`/api/ingredient?name=${encodeURI(ingredient)}`, {
      method: 'POST'
    }).then(() => {
      setRecentlyAdded(ingredient)
      setIngredient('')
      setUpdated(!updated)
      setShow(false)
    })
  }
}
  return (
    <div className="form-container">
      <form onSubmit={(e) => handleSubmit(e)}>
        <input
          ref={ref}
          type="text"
          name="name"
          placeholder="Ingredient Name"
          className="input"
          onChange={(e) => setIngredient(e.target.value)}
        />
        <div className="form-button-container">
          <button
            type="button"
            className="form-button-cancel"
            onClick={() => {
              setIngredient('')
              setShow(false)
            }}
          >Cancel</button>
          <button type="submit" className="form-button">Add</button>
        </div>
    </form>
    </div>
  )
}