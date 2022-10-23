import React from "react";
import { VictoryChart, VictoryBar, VictoryScatter } from "victory";

export default function PopularityAreaChart({ tracks }) {
    const data = []
    if (tracks){
        for (let i = 0; i < tracks.length; i++) {
            const currentTrack = tracks[i]
            data.push({ x: currentTrack.name, y: currentTrack.popularity })
        }
        return (
            <div className="popularity-area-chart">
                <VictoryChart horizontal maxDomain={{ y: 100 }}>
                    <VictoryBar data={data} />
                    <VictoryScatter data={data} />
                </VictoryChart>
            </div>
        )
    }
}