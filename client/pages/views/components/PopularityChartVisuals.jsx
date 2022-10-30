import React from "react";
import { VictoryChart, VictoryBar, VictoryScatter, VictoryAxis, VictoryLabel } from "victory";

export default function PopularityChartVisuals({ data }) {
    return (
        <VictoryChart domainPadding={10} title="Your Most Popular and Least Popular 5 Songs" padding={{ bottom: 50, left: 70, top: 10 }}>
            <VictoryAxis dependentAxis={true} labelComponent={(<VictoryLabel />)} label={"Popularity"} />
            <VictoryAxis
                tickLabelComponent={(
                    <VictoryLabel
                        textAnchor={'end'}
                        dx={5}
                    />
                )}
                style={{
                    tickLabels: {
                        fontSize: 10
                    },
                }}
            />
            <VictoryBar
                padding={{ right: 5 }}
                data={data}
                horizontal={true}
                domain={{ y: [0, 100] }}
                style={{
                    data: {
                        fill: ({ datum }) =>
                            datum.x === 'Mean Value' ? 'red' : '#1ED760'
                    }
                }}
            />
            <VictoryScatter data={data}
                style={{ labels: { fill: "black" } }}
                labels={({ datum }) => datum.y}
                labelComponent={<VictoryLabel />} />
        </VictoryChart>
    )
}