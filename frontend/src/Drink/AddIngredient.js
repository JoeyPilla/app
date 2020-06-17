import React, { useState} from 'react';
import './AddIngredient.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faPlusCircle } from '@fortawesome/free-solid-svg-icons'
import { navigate } from '@reach/router';

export default function AddIngredient({updated, setUpdated, recipeId}) {
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
          recipeId={recipeId}
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

const Modal = ({ children, show, recipeId}) => {
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

const Form = ({ updated, setUpdated, setShow, recipeId }) => {
  const [ingredient, setIngredient] = useState(0)
  const [amount, setAmount] = useState('')
  const [units, setUnits] = useState('')
  const [possibleIngredients, setPossibleIngredients] = useState([])
  const possibleUnits = ['oz', 'cups', 'ml']
  
  React.useEffect(() => {
    const fetchIngredients = async () => {
      const data = await fetch('/api/ingredient/all')
      const json = await data.json()
      setPossibleIngredients(json)
    }
    fetchIngredients()
  }, [])
  const ingredientOptions = possibleIngredients.map(ingredient => <option value={ingredient.id}>{ingredient.name}</option>)
  const unitsOptions = possibleUnits.map(unit => <option value={unit}>{unit}</option>)
  const handleSubmit = (e) => {
    e.preventDefault()
    console.log(ingredient)
    console.log(amount)
    console.log(units)
  if (ingredient > 0 && amount !== '' && units !== '') {
    fetch(`/api/drink?ingredientId=${encodeURI(ingredient)}&recipeId=${encodeURI(recipeId)}&measure=${eval(amount)}&unitOfMeasurement=${units}`, {
      method: 'POST'
    }).then(async (res) => {
      const json = await res.json()
      setUpdated(!updated)
      setShow(false)
    })
  }
}
  return (
    <div classNmae='not-ready-container'>

    <div className='add-ingredient-input-container'>
    <form className="add-ingredient-form" onSubmit={(e) => handleSubmit(e)}>
    <div className='temp'>
      Ingredient:
      <select className='select-box' value={ingredient} onChange={e => setIngredient(e.target.value)}>
        <option value={0}>Select an ingredient...</option>
        {ingredientOptions}
      </select>
    </div>
    <div className='temp'>
      Amount:
      <input
        value={amount}
        onChange={(e) => setAmount(e.target.value)}
        placeholder="enter ingredient amount"
        />
    </div>
    <div className='temp'>
      Units:
      <select className='select-box' value={units} onChange={e => setUnits(e.target.value)}>
        <option value={''}>Select a unit...</option>
        {unitsOptions}
      </select>
      </div>
          <button
            type='button'
            className='form-button-cancel'
            onClick={() => {
              setIngredient('')
              setAmount('')
              setUnits(0)
              setShow(false)
            }}
          >cancel</button>
      <button type='submit' className='form-button'>submit</button>
      </form>
      </div>
        </div>
  )
}