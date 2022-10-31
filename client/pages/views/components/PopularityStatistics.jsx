import React, { useState } from "react";
import PopularityStatisticVisualization from "./PopularityStatisticVisualization";
import { Drawer, Col, Typography, Row } from "antd";

const { Text, Link } = Typography;

export default function PopularityStatistics({ tracks, name }) {
    const data = []
    const [open, setOpen] = useState(false);

    const showDrawer = () => {
        setOpen(true);
    };

    const onClose = () => {
        setOpen(false);
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

            return (
                <Col className="popularity-area-chart">
                    <Col style={{ "paddingRight": "20px" }}>
                        <h3>5 Most Popular Tracks</h3>
                        {mostPopularTracks.map((track, index) => <p>{index + 1}.{track.x}</p>)}
                        <Link target="_blank" onClick={showDrawer}>
                            Click to view details
                        </Link>
                        <Drawer title={`Popularity statistics of ${name}`} open={open} onClose={onClose} placement="bottom" height={800} footer={`Mean Value = ${meanPopularity.toFixed(2)}`}>
                            {length > 50 ? <Text italic>{length} tracks... You listen to all of these? Really?</Text> : <Text>{length} tracks</Text>}
                            <PopularityStatisticVisualization data={data} width={700} height={800} domainPadding={10} />
                        </Drawer>
                    </Col>
                    <Col>
                        <h3>5 Least Popular Tracks</h3>
                        {leastPopularTracks.map((track, index) => <p>{index + 1}.{track.x}</p>)}
                    </Col>
                </Col>
            )
        }
        else {
            return (
                <Col>
                    <Row>
                        <h3>Most Popular Tracks</h3>
                    </Row>
                    <Row>
                        <Col style={{"textAlign": "start", "paddingRight": "20px"}}>
                            {data.map((track, index) => {
                                if (index <= 4)
                                    return <p>{index + 1}.{track.x}</p>
                            })}
                            <Link target="_blank" onClick={showDrawer}>
                                Click to view details
                            </Link>
                            <Drawer title="Track popularity" open={open} onClose={onClose} placement="bottom" footer={`Mean Value = ${meanPopularity.toFixed(2)}`}>
                                <PopularityStatisticVisualization data={data} width={300} height={300} domainPadding={15}/>
                            </Drawer>
                        </Col>
                        <Col style={{"textAlign": "start"}}>
                            {data.map((track, index) => {
                                if (index > 4)
                                    return <p>{index + 1}.{track.x}</p>
                            })}
                        </Col>
                    </Row>
                </Col>
            )
        }
    }
}