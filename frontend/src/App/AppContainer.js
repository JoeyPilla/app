import React, {useEffect, useState} from 'react';
import ToolBar from '../ToolBar/ToolBar'
import './AppContainer.css'

export default function Home({children}) {
  return (
    <div className="app-container">
      <ToolBar/>
      {children}
    </div>
    )
}
