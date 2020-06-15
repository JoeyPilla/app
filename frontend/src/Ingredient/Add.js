import React, { useState, useRef} from 'react';
import './Add.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faPlusCircle } from '@fortawesome/free-solid-svg-icons'

export default function Add({updated, setUpdated}) {
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
       <FontAwesomeIcon
        className="plus-circle-icon"
        icon={faPlusCircle}
        onClick={() => setShow(true)}
      />
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
    <div>
      <form onSubmit={(e) => handleSubmit(e)}>
        <label>
          Ingredient Name:
        <input
          ref={ref}
          type="text"
          name="name"
          onChange={(e) => setIngredient(e.target.value)}
          />
      </label>
    </form>
    </div>
  )
}