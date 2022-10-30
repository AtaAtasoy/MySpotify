import React from "react";
import PopularityChartVisuals from "./PopularityChartVisuals";

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
                    <PopularityChartVisuals data={visualizedData} />
                </div>
            )
        }
        else {
            return (
                <div className="popularity-area-chart">
                    <PopularityChartVisuals data={data} />
                </div>
            )
        }
    }
}