import React from 'react';
import { motion } from 'framer-motion'


const containerStyle = {
    position: 'relative',
    widt: "3rem",
    height: "3rem"
};

const spinTransition = {
    loop: Infinity,
    ease: "linear",
    duration: 1,
};

const circleStyle = {
    display: "block",
    width: "3rem",
    height: "3rem",
    border: "0.5rem solid #e9e9e9",
    borderTop: "0.5rem solid #21e065",
    borderRadius: "50%",
    position: "absolute",
    top: 0,
    left: 0
};

export default function CircleLoader(){

    return(
        <div style={containerStyle}>
            <motion.span style={circleStyle} animate={{ rotate: 360 }} transition={spinTransition}/>
        </div>
    )
}