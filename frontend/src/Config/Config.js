import React from 'react';
import './Config.css';
import MotorConfig from './MotorConfig';
import MotorTiming from './MotorTiming';

export default function Config() {
  return (
    <div className="not-ready-container">
      <MotorConfig />
      <MotorTiming />
    </div>
  )
}
