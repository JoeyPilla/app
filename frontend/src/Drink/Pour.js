import React, { useEffect, useState, useRef } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faSpinner } from '@fortawesome/free-solid-svg-icons'
import './Pour.css'
export default function Pour({ recipeId }) {
    const [duration, setDuration] = useState(0)
    const [count, setCount] = useState(0)
    const pour = async () => {
        const data = await fetch(`/api/drink/pour?recipeId=${recipeId}`, {method: 'POST'})
        const json = await data.json()
        setDuration(json)
    }
    
    return (
        <>
            <h1 onClick={() => pour()}>Pour</h1>
            {duration > 0 && <Timer reset={() => setDuration(0)} stop={duration}/>}
        </>
    )
}

const Timer = ({reset, stop}) => {
    const [time, setTime] = useState(new Date().toLocaleTimeString());
    const secondsPassed = useRef(0);
  
    useEffect(() => {
      const timeout = setTimeout(() => {
        const date = new Date()
        secondsPassed.current = secondsPassed.current + 1;
        setTime(date.toLocaleTimeString());
      }, 1000);
      return () => {
        clearTimeout(timeout);
      }
    }, [time]);
    const percent = secondsPassed.current / stop * 100
    if (percent > 100) {
        reset()
    }
    return (
        <>
        <FontAwesomeIcon
            className="spinner-icon"
            icon={faSpinner}
            spin
        />
      <div className="timer-container">
        <div style={{width: `${percent}%`}} className="progress-bar"/>
      </div>
    </>
    )
  }