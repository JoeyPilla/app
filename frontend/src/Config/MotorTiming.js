import React from 'react';
import './MotorTiming.css';
import InputContainerElement from './InputContainerElement'

export default function MotorTiming() {
  return (
      <div className="motor-timing">
        <InputContainerElement number={1}/>
        <InputContainerElement number={2}/>
        <InputContainerElement number={3}/>
        <InputContainerElement number={4}/>
        <InputContainerElement number={5}/>
        <InputContainerElement number={6}/>
        <InputContainerElement number={7}/>
      </div>
  )
}