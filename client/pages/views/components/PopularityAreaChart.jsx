import React from "react";
import { VictoryChart, VictoryBar, VictoryScatter } from "victory";

export default function PopularityAreaChart({ tracks }) {

    const data = tracks.reduce((accumulator, current) => { return accumulator + { x: current.name, y: current.popularity } }, [])

    return (
        <div className="popularity-area-chart">
            <VictoryChart horizontal maxDomain={{ y: 100 }}>
                <VictoryBar data={data} />
                <VictoryScatter data={data} />
            </VictoryChart>
        </div>
    )

}