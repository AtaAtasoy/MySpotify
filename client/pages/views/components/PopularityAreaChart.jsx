import React, { useState } from "react";
import PopularityChartVisuals from "./PopularityChartVisuals";
import { Button, Drawer, Modal } from "antd";

export default function PopularityAreaChart({ tracks }) {
    const data = []
    const [isModalOpen, setIsModalOpen] = useState(false);

    const handleOk = () => {
        setIsModalOpen(false);
    };

    const handleCancel = () => {
        setIsModalOpen(false);
    };

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

            const visualizedData = leastPopularTracks.concat(mostPopularTracks)
            //TODO:Display top5 and bottom5 most popular via text.
            //Provide button to display the details.
            //Show the detiled visualization in the Drawer component
            return (
                <div className="popularity-area-chart">
                    <div style={{ "paddingRight": "20px" }}>
                        <h3>5 Most Popular Tracks</h3>
                        {mostPopularTracks.map((track) => <p>{track.x}</p>)}
                        <Modal title="Basic Modal" open={isModalOpen} onOk={handleOk} onCancel={handleCancel}>
                            <p>Some contents...</p>
                            <p>Some contents...</p>
                            <p>Some contents...</p>
                        </Modal>
                    </div>
                    <div>
                        <h3>5 Least Popular Tracks</h3>
                        {leastPopularTracks.map((track) => <p>{track.x}</p>)}
                    </div>
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