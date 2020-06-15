import React, { useEffect } from 'react';
import { Router } from "@reach/router"
import './App.css';
import AppContainer from './AppContainer'
import Home from '../Home/Home'
import Config from '../Config/Config'
import AddDrink from '../Drink/AddDrink'
import Ingredients from '../Ingredient/Ingredients'
import Drinks from '../Drink/Drinks'
import Drink from '../Drink/Drink'
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
      <AppContainer path = "/">
        <Home path="/" />
        <Config path="config" />
        {/* <IngredientAdd path="/ingredient/add" /> */}
        <AddDrink path="/new" />
        <Ingredients path="/ingredients" />
        <Drinks path="/drink"/>
        <Drink path="/drink/:id"/>
      </AppContainer>
    </Router>
  )
}

