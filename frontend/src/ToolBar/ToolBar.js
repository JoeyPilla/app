import React, { useEffect, useState, useRef } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faCog, faBars } from '@fortawesome/free-solid-svg-icons'
import './ToolBar.css'
import { navigate, Link } from '@reach/router';
import { useSpring, animated } from 'react-spring'

export function useOnClickOutside(ref, handler) {


  useEffect(
    () => {
      const listener = event => {
        // Do nothing if clicking ref's element or descendent elements
        if (!ref.current || ref.current.contains(event.target)) {
          return;
        }

        handler(event);
      };

      document.addEventListener('mousedown', listener);
      document.addEventListener('touchstart', listener);

      return () => {
        document.removeEventListener('mousedown', listener);
        document.removeEventListener('touchstart', listener);
      };
    },
    // Add ref and handler to effect dependencies
    // It's worth noting that because passed in handler is a new ...
    // ... function on every render that will cause this effect ...
    // ... callback/cleanup to run every render. It's not a big deal ...
    // ... but to optimize you can wrap handler in useCallback before ...
    // ... passing it into this hook.
    [ref, handler]
  );
}


export default function ToolBar() {
  const ref = useRef();
  const [toggle, setToggle] = useState(false)
  const props = useSpring({
    left: toggle ? '0px' : '-150px'
  })

  const opacity = useSpring({
    opacity: toggle ? 0.5 : 0
  })

  useOnClickOutside(ref, ()=>setToggle(false))

  return (
    <>
      <div style={props} className="tool-bar">
          <FontAwesomeIcon className="cog-icon" size='30px' icon={faBars} onClick={() => setToggle(!toggle)}/>
          <FontAwesomeIcon
            className="cog-icon"
            icon={faCog}
            onClick={() => navigate('/config')}
          />
        </div>
      <animated.div style={props} className="nav-bar" ref={ref}>
        <div className="empty" />
        <div className="nav-bar-links">
          <Link
            to='/'
            className="tool-bar-link"
            onClick={() => setToggle(false)}
          >
            Home
          </Link>
          <Link
            to='drink'
            className="tool-bar-link"
            onClick={() => setToggle(false)}
          >
            Drinks
          </Link>
          <Link
            to='ingredients'
            className="tool-bar-link"
            onClick={() => setToggle(false)}
          >
            Ingredients
          </Link>
        </div>
      </animated.div>
      {toggle && <animated.div style={opacity} className="overlay"/>}
    </>
  )
}
