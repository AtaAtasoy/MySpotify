import React, { useRef, useEffect } from "react";

/**
 type Track struct {
     Id         string `json:"id"`
     Name       string `json:"name"`
     Popularity float64 `json:"popularity"`
     Acousticness float64 `json:"acousticness"`
     Danceability float64 `json:"danceability"`
     Duration_ms float64 `json:"duration_ms"`
     Energy float64 `json:"energy"`
     Instrumentalness float64 `json:"instrumentalness"`
     Liveness float64 `json:"liveness"`
     Loudness float64 `json:"loudness"`
     Mode float64 `json:"mode"`
     Speechiness float64 `json:"speechiness"`
     Tempo float64 `json:"tempo"`
     Valence float64 `json:"valence"`
 }
 */

export default function PlaylistInformation({tracks}){
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

    console.log('Tracks', tracks)
    return(
        <div className="playlist-info-container">
            <p className="playlist-info-text">
                Popularity: {avgPopularity.current}<br/>
                Acousticness: {avgAcousticness.current}<br/>
                Danceability: {avgDanceability.current}<br/>
                Duration: {avgDuration.current}<br/>
                Energy: {avgEnergy.current}<br/>
                Instrumentalness: {avgInstrumentalness.current}<br/>
                Tempo: {avgTempo.current}<br/>
                Loudness: {avgLoudness.current}
            </p>
        </div>
    )

}