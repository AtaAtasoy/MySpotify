import React from 'react';
import Image from "next/future/image"

export default function Artist({name, popularity, image}) {
    if (image) {
        return (
            <div className='artist-container'>
                <Image alt='artist-image' width={image.width} height={image.height} src={image.url} style={{"borderRadius": "50%"}} />
                <br/>
                <label>{name} {popularity}</label>
            </div>
        )
    } else {
        return (
            <div className='artist-container'>
                <Image alt='artist-image' width={250} height={250} src={"https://thispersondoesnotexist.com/image"} style={{"borderRadius": "50%"}} />
                <br/>
                <label>{"John Doe"} {90}</label>
            </div>
        )
    }
}