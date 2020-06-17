import React, { useState} from 'react';
import './MotorTiming.css';
import { navigate } from '@reach/router'
import InputContainerElement from './InputContainerElement'

export default function MotorTiming() {
  const [motor0, setMotor0] = useState(0)
  const [motor1, setMotor1] = useState(0)
  const [motor2, setMotor2] = useState(0)
  const [motor3, setMotor3] = useState(0)
  const [motor4, setMotor4] = useState(0)
  const [motor5, setMotor5] = useState(0)
  const [motor6, setMotor6] = useState(0)

  const handleSubmit = (e) => {
    e.preventDefault()
    fetch('/setMotorOffset', {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        motor1: motor0,
        motor2: motor1,
        motor3: motor2,
        motor4: motor3,
        motor5: motor4,
        motor6: motor5,
        motor7: motor6,
      })
    }).then(() => navigate('/'))
  }

  const handleClick = (e, motor) => {
    e.preventDefault()
    fetch(`/TestMotor?motor=${motor}`)
  }

  return (
    <form onSubmit={(e) => handleSubmit(e)}>
      <div className="input-container">
        <InputContainerElement setMotor={setMotor0} number={1}/>
        <InputContainerElement setMotor={setMotor1} number={2}/>
        <InputContainerElement setMotor={setMotor2} number={3}/>
        <InputContainerElement setMotor={setMotor3} number={4}/>
        <InputContainerElement setMotor={setMotor4} number={5}/>
        <InputContainerElement setMotor={setMotor5} number={6}/>
        <InputContainerElement setMotor={setMotor6} number={7}/>
        <button type="submit" className="form-button">Submit</button>
      </div>
    </form>
  )
}