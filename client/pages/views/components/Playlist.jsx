import React, { useRef, useEffect } from 'react';
import Image from "next/future/image"
import AttributesRadialAreaChart from './AttributesRadialAreaChart';
import PopularityAreaChart from './PopularityAreaChart';

/** 
type Playlist struct {
    Name   string        `json:"name"`
    Tracks []track.Track `json:"tracks"`
    Images []interface{} `json:"images"`
}
*/

export default function Playlist({ name, tracks, image }) {
    const length = tracks ? tracks.length : 1
    const avgPopularity = useRef(0) // 0 <= popularity <= 100
    const avgAcousticness = useRef(0) // 0 <= acousticness <= 1
    const avgDanceability = useRef(0) // 0 <= danceability <= 1
    const avgDuration = useRef(0) // duration in milliseconds
    const avgEnergy = useRef(0) // 0 <= energy <= 1
    const avgInstrumentalness = useRef(0) // 0 <= avgInstrumentalness <= 1
    const avgSpeechines = useRef(0)
    const avgValence = useRef(0)
    const avgTempo = useRef(0) // BPM

    useEffect(() => {
        avgPopularity.current = (tracks.reduce((accumulator, current) => { return accumulator + current.popularity }, 0) / length).toFixed(2)
        avgAcousticness.current = ((tracks.reduce((accumulator, current) => { return accumulator + current.acousticness }, 0) / length) * 100).toFixed(2)
        avgDanceability.current = ((tracks.reduce((accumulator, current) => { return accumulator + current.danceability }, 0) / length) * 100).toFixed(2)
        avgEnergy.current = ((tracks.reduce((accumulator, current) => { return accumulator + current.energy }, 0) / length) * 100).toFixed(2)
        avgInstrumentalness.current = ((tracks.reduce((accumulator, current) => { return accumulator + current.instrumentalness }, 0) / length) * 100).toFixed(2)
        avgSpeechines.current = ((tracks.reduce((accumulator, current) => { return accumulator + current.speechiness }, 0) / length) * 100).toFixed(2)
        avgValence.current = ((tracks.reduce((accumulator, current) => { return accumulator + current.valence }, 0) / length) * 100).toFixed(2)
        avgDuration.current = Math.round((tracks.reduce((accumulator, current) => { return accumulator + current.duration_ms }, 0) / length) / 1000)
        avgTempo.current = Math.round(tracks.reduce((accumulator, current) => { return accumulator + current.tempo }, 0) / length)
    }, [tracks, length])

    /**
     * TODO:Adjust the data to show on click to playlist image
     */
    if (tracks) {
        return (
            <div className='playlist'>
                <Image alt='playlist-image' width={150} height={150} src={image ? image.url : "https://thispersondoesnotexist.com/image"} style={{ "borderRadius": "50%" }} />
                <h3>{name ? name : "Playlist"}</h3>
                <AttributesRadialAreaChart data={[
                    { x: "Speechiness", y: Number(avgSpeechines.current) },
                    { x: "Acousticness", y: Number(avgAcousticness.current) },
                    { x: "Danceability", y: Number(avgDanceability.current) },
                    { x: "Energy", y: Number(avgEnergy.current) },
                    { x: "Instrumentalness", y: Number(avgInstrumentalness.current) },
                    { x: "Valence", y: Number(avgValence.current) }
                ]} />
                <PopularityAreaChart tracks={tracks} />
            </div>
        )
    }
}