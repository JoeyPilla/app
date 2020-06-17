import React, { useState} from 'react';
import './MotorTiming.css';
import { navigate } from '@reach/router'

export default function InputContainerElement({ setMotor, number }) {
  const handleClick = (e, motor) => {
    e.preventDefault()
    fetch(`/TestMotor?motor=${motor}`)
  }

  return (
    <div className="input-container-element">
      <input
        type="text"
        name="name"
        placeholder="Motor 1 Offset"
        className="input"
        onChange={(e) => setMotor(e.target.value)}
      />
      <button
        type="button"
        onClick={(e) => handleClick(e, number-1)}
      >
        run motor {number}
      </button>
    </div>
  )
}