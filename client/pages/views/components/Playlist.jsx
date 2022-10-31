import React from 'react';
import Image from "next/future/image"
import AttributesRadialAreaChart from './AttributesRadialAreaChart';
import PopularityStatistics from './PopularityStatistics';
import { Col, Row } from 'antd';

/** 
type Playlist struct {
    Id string `json:"id"`
    Name   string        `json:"name"`
    Tracks []t.Track `json:"tracks"`
    Images []interface{} `json:"images"`
    Attributes map[string]float64 `json:"attributes"`
}*/

export default function Playlist({ name, tracks, image, attributes }) {
    if (attributes) {
        return (
            <div className='playlist'>
                <Image alt='playlist-image' width={150} height={150} src={image ? image.url : "https://thispersondoesnotexist.com/image"} style={{ "borderRadius": "50%"}} />
                <h3 style={{'paddingTop': '20px'}}>{name ? name : "Playlist"}</h3>
                <Row className='playlist-visualization-container' align='middle'>
                    <Col>
                        <AttributesRadialAreaChart
                            acousticness={attributes["acousticness"]}
                            danceability={attributes["danceability"]}
                            energy={attributes["energy"]}
                            instrumentalness={attributes["instrumentalness"]}
                            speechines={attributes["speechiness"]}
                            valence={attributes["valence"]} />
                    </Col>
                    <Col>
                        <PopularityStatistics tracks={tracks} name={name} />
                    </Col>
                </Row>
            </div>
        )
    } else {
        return (
            <div></div>
        )
    }
}