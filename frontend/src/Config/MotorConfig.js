import React, {useEffect, useState} from 'react';
import './Config.css';
import IngredientDropdown from './IngredientDropdown'
import { navigate } from '@reach/router'

export default function Config() {
  const [ingredients, setIngredients] = useState([])
  const [motor0, setMotor0] = useState(0)
  const [motor1, setMotor1] = useState(0)
  const [motor2, setMotor2] = useState(0)
  const [motor3, setMotor3] = useState(0)
  const [motor4, setMotor4] = useState(0)
  const [motor5, setMotor5] = useState(0)
  const [motor6, setMotor6] = useState(0)
  useEffect(() => {
    const fetchIngredients = async () => {
      const data = await fetch('/api/ingredient/all')
      const json = await data.json()
      setIngredients(json)
    }
    fetchIngredients()
  }, [])
  useEffect(() => {
    const fetchMotorMap = async () => {
      const data = await fetch('/getMotors')
      const json = await data.json()
      setMotor0(json.motor1)
      setMotor1(json.motor2)
      setMotor2(json.motor3)
      setMotor3(json.motor4)
      setMotor4(json.motor5)
      setMotor5(json.motor6)
      setMotor6(json.motor7)
    }
    fetchMotorMap()
  }, [])

  const options = ingredients.map(ingredient => <option value={ingredient.id}>{ingredient.name}</option>)

  const handleSubmit = (e) => {
    e.preventDefault()
    fetch('/setMotors', {
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

  return (
    <form onSubmit={(e) => handleSubmit(e)}>
      <div className="input-container">
        <IngredientDropdown
          item={1}
          value={motor0}
          setValue={setMotor0}
              options={options}
              className='select-box'
          />
        <IngredientDropdown 
          item={2} 
          value={motor1} 
          setValue={setMotor1} 
          options={options}
          />
        <IngredientDropdown 
          item={3} 
          value={motor2} 
          setValue={setMotor2} 
          options={options}
          />
        <IngredientDropdown 
          item={4} 
          value={motor3} 
          setValue={setMotor3} 
          options={options}
          />
        <IngredientDropdown 
          item={5} 
          value={motor4} 
          setValue={setMotor4} 
          options={options}
          />
        <IngredientDropdown 
          item={6} 
          value={motor5} 
          setValue={setMotor5} 
          options={options}
          />
        <IngredientDropdown 
          item={7} 
          value={motor6} 
          setValue={setMotor6} 
          options={options}
          />
        <button type="submit" className="form-button">Submit</button>
        </div>
    </form>
  )
}
