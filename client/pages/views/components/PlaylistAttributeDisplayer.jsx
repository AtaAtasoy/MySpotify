import React from "react";
import { CircularProgressbar, buildStyles } from "react-circular-progressbar";

export default function PlaylistAttributeDisplayer({attributeName, value, minValue, maxValue}) {
    return (
        <div className="playlist-attribute-displayer">
            <h4>{attributeName}</h4>
            <CircularProgressbar value={value} minValue={minValue} maxValue={maxValue} text={attributeName != "Duration" ? `${value}` : `${value} seconds` } styles={buildStyles({ textSize: '16px', textColor: '#21e065', pathColor: `rgb(33, 224, 101)`})}/>
        </div>
    )
}