import React, { useRef, useEffect } from 'react';
import Image from "next/future/image"
import { buildStyles, CircularProgressbar } from 'react-circular-progressbar';
import 'react-circular-progressbar/dist/styles.css';
import PlaylistAttributeDisplayer from './PlaylistAttributeDisplayer';

/** 
type Playlist struct {
    Name   string        `json:"name"`
	Tracks []track.Track `json:"tracks"`
	Images []interface{} `json:"images"`
}
*/

export default function Playlist({name, tracks, image}) {
    const length = tracks ? tracks.length : 1
    const avgPopularity = useRef(0) // 0 <= popularity <= 100
    const avgAcousticness = useRef(0) // 0 <= acousticness <= 1
    const avgDanceability = useRef(0) // 0 <= danceability <= 1
    const avgDuration = useRef(0) // duration in milliseconds
    const avgEnergy = useRef(0) // 0 <= energy <= 1
    const avgInstrumentalness = useRef(0) // 0 <= avgInstrumentalness <= 1
    const avgTempo = useRef(0) // BPM

    useEffect(() => {
        avgPopularity.current = (tracks.reduce((accumulator, current) => {return accumulator + current.popularity}, 0) / length).toFixed(2)
        avgAcousticness.current = ((tracks.reduce((accumulator, current) => {return accumulator + current.acousticness}, 0) / length) * 100).toFixed(2)
        avgDanceability.current = ((tracks.reduce((accumulator, current) => {return accumulator + current.danceability}, 0) / length) * 100).toFixed(2)
        avgEnergy.current = ((tracks.reduce((accumulator, current) => {return accumulator + current.energy}, 0) / length) * 100).toFixed(2)
        avgInstrumentalness.current = ((tracks.reduce((accumulator, current) => {return accumulator + current.instrumentalness}, 0) / length) * 100).toFixed(2)
        avgDuration.current = Math.round((tracks.reduce((accumulator, current) => {return accumulator + current.duration_ms}, 0) / length) / 1000)
        avgTempo.current = Math.round(tracks.reduce((accumulator, current) => {return accumulator + current.tempo}, 0) / length)
    }, [tracks, length])

    /**
     * TODO:Adjust the data to show on click to playlist image
     */
    return (
        <div className='playlist'>
            <Image alt='playlist-image' width={150} height={150} src={image ? image.url : "https://thispersondoesnotexist.com/image"} style={{"borderRadius": "50%"}}/>
            <h3>{name ? name : "Playlist"}</h3>
            {tracks ?
                <div className="playlist-circular-data-container">
                        <PlaylistAttributeDisplayer attributeName={"Popularity"} value={avgPopularity.current} minValue={0} maxValue={100} />
                        <PlaylistAttributeDisplayer attributeName={"Acousticness"} value={avgAcousticness.current} minValue={0} maxValue={100} />
                        <PlaylistAttributeDisplayer attributeName={"Danceability"} value={avgDanceability.current} minValue={0} maxValue={100} />
                        <PlaylistAttributeDisplayer attributeName={"Energy"} value={avgEnergy.current} minValue={0} maxValue={100} />
                        <PlaylistAttributeDisplayer attributeName={"Instrumentalness"} value={avgInstrumentalness.current} minValue={0} maxValue={100} />
                        <PlaylistAttributeDisplayer attributeName={"Tempo"} value={avgTempo.current} minValue={0} maxValue={avgTempo.current} />
                        <PlaylistAttributeDisplayer attributeName={"Duration"} value={avgDuration.current} minValue={0} maxValue={avgDuration.current} />
                </div>
                 : <div className='podcast-playlist'/>}
        </div>
    )
}