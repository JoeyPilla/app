import React, { useState} from 'react';
import './AddDrink.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faPlusCircle } from '@fortawesome/free-solid-svg-icons'
import { navigate } from '@reach/router';

export default function AddDrink({updated, setUpdated}) {
  const [recentlyAdded, setRecentlyAdded] = useState('');
  const [show, setShow] = useState(false)


  return (
    <>
    <Modal show={show}>
        <Form
          setRecentlyAdded={setRecentlyAdded}
          updated={updated}
          setUpdated={setUpdated}
          setShow={setShow}
        />
      </Modal>
      {recentlyAdded.length > 0 && <p>Successfully added {recentlyAdded}!</p>}
      {!show && <FontAwesomeIcon
        className="plus-circle-icon"
        icon={faPlusCircle}
        onClick={() => setShow(true)}
      />}
    </>
    )
}

const Modal = ({ children, show }) => {
  if (show) {
    return (
      <div className="modal">
      {children}
    </div>
  )
  } else {
    return <></>
}
}

const Form = ({setRecentlyAdded, updated, setUpdated, setShow}) => {
  const [recipe, setRecipe] = useState('');

  const handleSubmit = (e) => {
  e.preventDefault()
  if (recipe.length > 0) {
    fetch(`/api/recipe?name=${encodeURI(recipe)}`, {
      method: 'POST'
    }).then(async (res) => {
      const json = await res.json()
      console.log(json)
      setRecentlyAdded(recipe)
      setRecipe('')
      setUpdated(!updated)
      setShow(false)
      navigate(`/drink/${json.id}`)
    })
  }
}
  return (
    <div className="form-container">
      <form onSubmit={(e) => handleSubmit(e)}>
        <input
          autofocus="true"
          type="text"
          name="name"
          placeholder="Recipe Name"
          className="input"
          onChange={(e) => setRecipe(e.target.value)}
        />
        <div className="form-button-container">
          <button
            type="button"
            className="form-button-cancel"
            onClick={() => {
              setRecipe('')
              setShow(false)
            }}
          >Cancel</button>
          <button type="submit" className="form-button">Add</button>
        </div>
      </form>
      
    </div>
  )
}