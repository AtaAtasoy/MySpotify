import React from 'react';
import Image from "next/future/image"

export default function Artist({name, popularity, image}) {
    return (
        <div className='artist-container'>
            <Image alt='artist-image' width={image.width} height={image.height} src={image.url} style={{"borderRadius": "50%"}} />
            <br/>
            <label>{name} {popularity}</label>
        </div>
    )
}