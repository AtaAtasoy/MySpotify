import React from "react";
import { VictoryArea, VictoryChart, VictoryPolarAxis, VictoryTheme } from "victory"

export default function AttributesRadialAreaChart({ data }) {
    return(
        <div className="radial-area-chart">
            {console.log(data)}
            <VictoryChart 
                polar
                theme={VictoryTheme.material}
                animate={{ duration: 2000, onLoad: { duration: 1000 }}}
                domain={{y: [0, 100]}}
            >
                <VictoryArea data={data} style={{ data: { fill: "#1ED760"}}} labels={({ datum }) => datum.y}/>
                <VictoryPolarAxis labelPlacement={"perpendicular"}/>
            </VictoryChart>
        </div>
    )
}