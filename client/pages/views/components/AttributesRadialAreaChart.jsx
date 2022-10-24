import React from "react";
import { VictoryArea, VictoryChart, VictoryPolarAxis, VictoryTheme } from "victory"

export default function AttributesRadialAreaChart({ speechines, acousticness,danceability, energy, instrumentalness, valence}) {
    if (speechines && acousticness && danceability && energy && instrumentalness && valence){
        const data = [
            {x: "Speechines", y: speechines },
            {x: "Acousticness", y: acousticness },
            {x: "Danceability", y: danceability },
            {x: "Energy", y: energy },
            {x: "Instrumentalness", y: instrumentalness},
            {x: "Valence", y: valence}
        ]
        console.log(data)
    
        return(
            <div className="radial-area-chart">
                {console.log("Rendering area chart" + data)}
                <VictoryChart 
                    polar
                    theme={VictoryTheme.material}
                    domain={{y: [0, 100]}}
                    animate={{duration: 2000, onLoad:{ duration: 1000}}}
                >
                    <VictoryArea data={data} style={{ data: { fill: "#1ED760"}}} labels={({ datum }) => datum.y}/>
                    <VictoryPolarAxis labelPlacement={"perpendicular"}/>
                </VictoryChart>
            </div>
        )
    }
}