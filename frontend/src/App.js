import React, { useEffect, useState } from 'react';
import { Router, Link, useNavigate } from "@reach/router"
import logo from './logo.svg';
import './App.css';
import Ready from './Ready'
import Add from './Add'
import NotReady from './NotReady'
export default function App() {
  // const navigate = useNavigate()
  useEffect(() => {
    const fetchData = async () => {
      const data = await fetch('/hello')
      const json = await data.json()
      if (!json.ready) {
        // navigate("/config")
      }
    }
    fetchData()
  }, [])
  return (
    <Router>
      <Ready path="/" />
      <NotReady path="config" />
      <Add path="add" />
    </Router>
  )
}

