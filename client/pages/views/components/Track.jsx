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
 import React from 'react'
 
 export default function Track(props){
    if (props) {
        return(
            <div className='playlist-tracks'>
                <label>{props.name}</label>
                <label>{props.popularity}</label>
                <label>{props.acousticness}</label>
                <label>{props.danceability}</label>
                <label>{props.duration_ms}</label>
                <label>{props.energy}</label>
                <label>{props.instrumentalness}</label>
                <label>{props.liveness}</label>
                <label>{props.tempo}</label>
            </div>
        )
    }
 }