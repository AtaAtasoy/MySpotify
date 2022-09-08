import React from 'react';

export default function Artist({name, popularity}) {
    
    return (
        <div className='artist-container'>
            <label>{name}</label>
            <label>{popularity}</label>
        </div>
    )
}