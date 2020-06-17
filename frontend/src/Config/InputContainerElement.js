import React, { useState} from 'react';
import './MotorTiming.css';
import { navigate } from '@reach/router'

export default function InputContainerElement({ number }) {
  const [motor, setMotor] = useState('')
  const handleClick = (e, motor) => {
    e.preventDefault()
    fetch(`/TestMotor?motor=${motor}`)
  }


  const handleSubmit = (e) => {
    e.preventDefault()
    fetch(`/setMotorPourRate?motor=${number-1}&value=${motor}`, {
      method: 'POST',
    }).then(()=> setMotor(''))
  }

  return (
    <div className="input-container-element">
      <input
        type="text"
        name="name"
        placeholder={`Motor ${number} Pour Rate`}
        className="input-2"
        value={motor}
        onChange={(e) => setMotor(e.target.value)}
      />
      <button
        type="button"
        onClick={(e) => handleClick(e, number-1)}
      >
        run
      </button>
      <button
        type="button"
        onClick={(e) => handleSubmit(e, number - 1)}
        disabled={motor<=0}
      >
        save
      </button>
    </div>
  )
}