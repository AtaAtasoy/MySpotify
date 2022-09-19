import React, { useRef, useEffect } from 'react';
import Image from "next/future/image"
/** 
type Playlist struct {
    Name   string        `json:"name"`
	Tracks []track.Track `json:"tracks"`
	Images []interface{} `json:"images"`
}
*/

export default function Playlist({name, tracks, image}) {
    const length = tracks.length
    const avgPopularity = useRef(0)
    const avgAcousticness = useRef(0)
    const avgDanceability = useRef(0)
    const avgDuration = useRef(0)
    const avgEnergy = useRef(0)
    const avgInstrumentalness = useRef(0)
    const avgTempo = useRef(0)
    const avgLoudness = useRef(0)

    useEffect(() => {
        avgPopularity.current = tracks.reduce((accumulator, current) => {return accumulator + current.popularity}, 0) / length
        avgAcousticness.current = tracks.reduce((accumulator, current) => {return accumulator + current.acousticness}, 0) / length
        avgDanceability.current = tracks.reduce((accumulator, current) => {return accumulator + current.danceability}, 0) / length
        avgDuration.current = tracks.reduce((accumulator, current) => {return accumulator + current.duration_ms}, 0) / length
        avgEnergy.current = tracks.reduce((accumulator, current) => {return accumulator + current.energy}, 0) / length
        avgInstrumentalness.current = tracks.reduce((accumulator, current) => {return accumulator + current.instrumentalness}, 0) / length
        avgTempo.current = tracks.reduce((accumulator, current) => {return accumulator + current.tempo}, 0) / length
        avgLoudness.current = tracks.reduce((accumulator, current) => {return accumulator + current.loudness}, 0) / length
    }, [tracks, length])


    return (
        <div className='playlist-container'>
            <Image alt='playlist-image' width={150} height={150} src={image.url} style={{"borderRadius": "50%"}}/>
            <h3>{name}</h3>
            {/** Playlist may contain podcasts instead of tracks, thus have to check if it contains tracks}*/}
            {tracks ? <p className="playlist-info-text">
                Popularity: {avgPopularity.current}<br/>
                Acousticness: {avgAcousticness.current}<br/>
                Danceability: {avgDanceability.current}<br/>
                Duration: {avgDuration.current}<br/>
                Energy: {avgEnergy.current}<br/>
                Instrumentalness: {avgInstrumentalness.current}<br/>
                Tempo: {avgTempo.current}<br/>
                Loudness: {avgLoudness.current}
            </p> : <div className='podcast-playlist'/>}
            {/**tracks ? tracks.map((trackData, i) => <Track key={i} {...trackData} />) : <div className='podcast-playlist'/> */}
        </div>
    )
}