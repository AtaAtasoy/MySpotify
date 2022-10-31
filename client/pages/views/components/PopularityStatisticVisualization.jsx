import React from "react";
import { VictoryChart, VictoryBar, VictoryScatter, VictoryAxis, VictoryLabel } from "victory";

export default function PopularityStatisticVisualization({ data, width, height, domainPadding, fontSize }) {
    if (width && data && height && domainPadding && fontSize) {
        return (
            <VictoryChart domainPadding={domainPadding} title="Popularity of the songs in your playlist" width={width} height={height}>
                <VictoryAxis dependentAxis={true} labelComponent={(<VictoryLabel />)} label={"Popularity"} />
                <VictoryAxis
                    tickLabelComponent={(
                        <VictoryLabel
                            textAnchor={'end'}
                        />
                    )}
                    style={{
                        tickLabels: {
                            fontSize: fontSize
                        },
                    }}
                />
                <VictoryBar
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
                    style={{ labels: { fill: "black", fontSize: fontSize } }}
                    labels={({ datum }) => datum.y}
                    labelComponent={<VictoryLabel />} />
            </VictoryChart>
        )
    }
}