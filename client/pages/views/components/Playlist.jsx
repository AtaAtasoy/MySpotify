import React from 'react';
import Track from './Track';
import Image from 'next/image';
/** 
type Playlist struct {
    Name   string        `json:"name"`
	Tracks []track.Track `json:"tracks"`
	Images []interface{} `json:"images"`
}
*/

export default function Playlist({name, tracks, image}) {
    return (
        <div className='playlist-container'>
            <Image alt='playlist-image' width={150} height={150} src={image.url} style={{"borderRadius": "50%"}}/>
            <h3>{name}</h3>
            {tracks.map((trackData, i) => <Track key={i} {...trackData} />)}
        </div>
    )
}