import React, { useRef, useEffect } from 'react';
import Image from "next/future/image"
import { buildStyles, CircularProgressbar } from 'react-circular-progressbar';
import 'react-circular-progressbar/dist/styles.css';

/** 
type Playlist struct {
    Name   string        `json:"name"`
	Tracks []track.Track `json:"tracks"`
	Images []interface{} `json:"images"`
}
*/

export default function Playlist({name, tracks, image}) {
    const length = tracks.length
    const avgPopularity = useRef(0) // 0 <= popularity <= 100
    const avgAcousticness = useRef(0) // 0 <= acousticness <= 1
    const avgDanceability = useRef(0) // 0 <= danceability <= 1
    const avgDuration = useRef(0) // duration in milliseconds
    const avgEnergy = useRef(0) // 0 <= energy <= 1
    const avgInstrumentalness = useRef(0) // 0 <= avgInstrumentalness <= 1
    const avgTempo = useRef(0) // BPM

    useEffect(() => {
        avgPopularity.current = (tracks.reduce((accumulator, current) => {return accumulator + current.popularity}, 0) / length).toPrecision(2)
        avgAcousticness.current = (tracks.reduce((accumulator, current) => {return accumulator + current.acousticness}, 0) / length).toPrecision(2)
        avgDanceability.current = (tracks.reduce((accumulator, current) => {return accumulator + current.danceability}, 0) / length).toPrecision(2)
        avgEnergy.current = (tracks.reduce((accumulator, current) => {return accumulator + current.energy}, 0) / length).toPrecision(2)
        avgInstrumentalness.current = (tracks.reduce((accumulator, current) => {return accumulator + current.instrumentalness}, 0) / length).toPrecision(2)
        avgDuration.current = (tracks.reduce((accumulator, current) => {return accumulator + current.duration_ms}, 0) / length) / 100
        avgTempo.current = tracks.reduce((accumulator, current) => {return accumulator + current.tempo}, 0) / length
    }, [tracks, length])

    /**
     * TODO: Fix floating point errors
     * TODO: Progress bar view also for duration and tempo
     * TODO: Column view for the circular progress bars
     */
    return (
        <div className='playlist-container'>
            <Image alt='playlist-image' width={150} height={150} src={image.url} style={{"borderRadius": "50%"}}/>
            <h3>{name}</h3>
            {tracks ? <div className="playlist-info-text">
                <div style={{ width: 150, height: 150 }}>
                    <CircularProgressbar value={(avgPopularity.current)} minValue={0} maxValue={100} text={`${avgPopularity.current}%`} styles={buildStyles({ textSize: '16px', textColor: '#21e065', pathColor: `rgb(33, 224, 101)`})}/>
                </div>
                <div style={{ width: 150, height: 150 }}>
                    <CircularProgressbar value={(avgAcousticness.current)} minValue={0} maxValue={1} text={`${avgAcousticness.current * 100}%`} styles={buildStyles({ textSize: '16px', textColor: '#21e065', pathColor: `rgb(33, 224, 101)`})}/>
                </div>
                <div style={{ width: 150, height: 150 }}>
                    <CircularProgressbar value={(avgDanceability.current)} minValue={0} maxValue={1} text={`${avgDanceability.current * 100}%`} styles={buildStyles({ textSize: '16px', textColor: '#21e065', pathColor: `rgb(33, 224, 101)`})}/>
                </div>
                <div style={{ width: 150, height: 150 }}>
                    <CircularProgressbar value={(avgEnergy.current)} minValue={0} maxValue={1} text={`${avgEnergy.current * 100}%`} styles={buildStyles({ textSize: '16px', textColor: '#21e065', pathColor: `rgb(33, 224, 101)`})}/>
                </div>
                <div style={{ width: 150, height: 150 }}>
                    <CircularProgressbar value={(avgInstrumentalness.current)} minValue={0} maxValue={1} text={`${avgInstrumentalness.current * 100}%`} styles={buildStyles({ textSize: '16px', textColor: '#21e065', pathColor: `rgb(33, 224, 101)`})}/>
                </div>
                <p>
                    Duration: {Math.round(avgDuration.current)} seconds<br/>
                    Tempo: {Math.round(avgTempo.current)} bpm<br/>
                </p>
            </div> : <div className='podcast-playlist'/>}
        </div>
    )
}