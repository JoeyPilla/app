import React, {useEffect, useState} from 'react';
import logo from './logo.svg';
import './App.scss';

function App() {
  useEffect(() => {
    const fetchData = async () => {
      const data = await fetch('/api/todo/all')
      const json = await data.json()
      console.log(json)
    }
    fetchData()
 }, [])
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
