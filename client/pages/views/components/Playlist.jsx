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
        avgPopularity.current = tracks.reduce((accumulator, current) => {return accumulator + current.popularity}, 0) / length
        avgAcousticness.current = tracks.reduce((accumulator, current) => {return accumulator + current.acousticness}, 0) / length
        avgDanceability.current = tracks.reduce((accumulator, current) => {return accumulator + current.danceability}, 0) / length
        avgDuration.current = tracks.reduce((accumulator, current) => {return accumulator + current.duration_ms}, 0) / length
        avgEnergy.current = tracks.reduce((accumulator, current) => {return accumulator + current.energy}, 0) / length
        avgInstrumentalness.current = tracks.reduce((accumulator, current) => {return accumulator + current.instrumentalness}, 0) / length
        avgTempo.current = tracks.reduce((accumulator, current) => {return accumulator + current.tempo}, 0) / length
    }, [tracks, length])


    return (
        <div className='playlist-container'>
            <Image alt='playlist-image' width={150} height={150} src={image.url} style={{"borderRadius": "50%"}}/>
            <h3>{name}</h3>
            {/** Playlist may contain podcasts instead of tracks, thus have to check if it contains tracks}*/}
            {tracks ? <div className="playlist-info-text">
                <div style={{ width: 100, height: 100 }}>
                    <CircularProgressbar value={avgPopularity.current} minValue={0} maxValue={100} text={`Popularity\n${avgPopularity.current}`} styles={buildStyles({ textSize: '16px', pathColor: `rgb(33, 224, 101)`})}/>
                </div>
                <div style={{ width: 100, height: 100 }}>
                    <CircularProgressbar value={avgAcousticness.current} minValue={0} maxValue={1} text={`Acousticness\n${avgAcousticness.current}`} styles={buildStyles({pathColor: `rgb(33, 224, 101)`})}/>
                </div>
                <div style={{ width: 100, height: 100 }}>
                    <CircularProgressbar value={avgDanceability.current} minValue={0} maxValue={1} text={`Danceability\n${avgDanceability.current}`} styles={buildStyles({pathColor: `rgb(33, 224, 101)`})}/>
                </div>
                <div style={{ width: 100, height: 100 }}>
                    <CircularProgressbar value={avgEnergy.current} minValue={0} maxValue={1} text={`Energy\n${avgEnergy.current}`} styles={buildStyles({pathColor: `rgb(33, 224, 101)`})}/>
                </div>
                <div style={{ width: 100, height: 100 }}>
                    <CircularProgressbar value={avgInstrumentalness.current} minValue={0} maxValue={1} text={`Instrumentalness\n${avgInstrumentalness.current}`} styles={buildStyles({pathColor: `rgb(33, 224, 101)`})}/>
                </div>
                <p>
                    Duration: {avgDuration.current / 1000}<br/>
                    Tempo: {avgTempo.current} bpm<br/>
                </p>
            </div> : <div className='podcast-playlist'/>}
            {/**tracks ? tracks.map((trackData, i) => <Track key={i} {...trackData} />) : <div className='podcast-playlist'/> */}
        </div>
    )
}