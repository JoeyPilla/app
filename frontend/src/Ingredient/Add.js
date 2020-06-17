import React, { useState, useRef} from 'react';
import './Add.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faPlusCircle } from '@fortawesome/free-solid-svg-icons'
import AddForm from './AddForm'
export default function Add({updated, setUpdated}) {
  const [recentlyAdded, setRecentlyAdded] = useState('');
  const [show, setShow] = useState(false)


  return (
    <>
    <Modal show={show}>
        <AddForm
          setRecentlyAdded={setRecentlyAdded}
          updated={updated}
          setUpdated={setUpdated}
          setShow={setShow}
        />
      </Modal>
      {recentlyAdded.length > 0 && <p>Successfully added {recentlyAdded}!</p>}
      {!show &&
        <FontAwesomeIcon
        className="plus-circle-icon"
        icon={faPlusCircle}
        onClick={() => {
          setShow(true);
          window.scrollTo(0, 0)
        }}
        />}
    </>
    )
}

const Modal = ({ children, show }) => {
  if (show) {
    return (
      <>
      <div className="modal-2">
          {children}
          <div className="color"/>
      </div>
        </>
  )
  } else {
    return <></>
}
}