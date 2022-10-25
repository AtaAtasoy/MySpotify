import React from "react";
import { VictoryChart, VictoryBar, VictoryScatter, VictoryAxis, VictoryLabel, VictoryLegend } from "victory";

export default function PopularityAreaChart({ tracks }) {
    const data = []
    if (tracks) {
        const length = tracks.length
        let totalPopularity = 0
        for (let i = 0; i < length; i++) {
            const currentTrack = tracks[i]
            data.push({ x: currentTrack.name, y: currentTrack.popularity })
            totalPopularity += currentTrack.popularity
        }
        const meanPopularity = totalPopularity / length

        if (length > 10) {
            data.sort((t1, t2) => t1.y - t2.y)
            const leastPopularTracks = data.slice(0, 5)
            const mostPopularTracks = data.slice(length - 6, length - 1)
            leastPopularTracks.push({ x: "Mean Value", y: Number(meanPopularity.toFixed(2)) })

            const visualizedData = leastPopularTracks.concat(mostPopularTracks)
            return (
                <div className="popularity-area-chart">
                    <VictoryChart domainPadding={10} title="Your Most Popular and Least Popular 5 Songs">
                        <VictoryAxis dependentAxis={true} labelComponent={(<VictoryLabel />)} label={"Popularity"} />
                        <VictoryAxis
                            axisLabelComponent={(
                                <VictoryLabel
                                    verticalAnchor="middle"    
                                />
                            )}
                            style={{
                                tickLabels: {
                                    fontSize: 10
                                },
                            }}
                        />
                        <VictoryBar
                            data={visualizedData}
                            horizontal={true}
                            domain={{ y: [0, 100] }}
                            style={{
                                data: {
                                    fill: ({ datum }) =>
                                        datum.x === 'Mean Value' ? 'red' : '#1ED760'
                                }
                            }}
                        />
                        <VictoryScatter data={visualizedData}
                            style={{ labels: { fill: "black" } }}
                            labels={({ datum }) => datum.y}
                            labelComponent={<VictoryLabel />} />
                    </VictoryChart>
                </div>
            )
        }
        else {
            return (
                <div className="popularity-area-chart">
                    <VictoryChart domainPadding={10} title="Your Most Popular and Least Popular 5 Songs">
                        <VictoryAxis dependentAxis={true} labelComponent={(<VictoryLabel x={0} />)} />
                        <VictoryAxis
                            axisLabelComponent={(
                                <VictoryLabel
                                    verticalAnchor="middle"
                                    textAnchor="start"
                                    x={0}
                                />
                            )}
                        />
                        <VictoryBar
                            data={data}
                            horizontal={true}
                            domain={{ y: [0, 100] }}
                            style={{ data: { fill: '#1ED760' } }}
                        />
                        <VictoryScatter data={data}
                            style={{ labels: { fill: "black" } }}
                            labels={({ datum }) => datum.y}
                            labelComponent={<VictoryLabel />} />
                    </VictoryChart>
                </div>
            )
        }
    }
}