import React, { useState } from "react";
import PopularityStatisticVisualization from "./PopularityStatisticVisualization";
import { Drawer, Col, Typography, Row } from "antd";

const { Text, Link } = Typography;

export default function PopularityStatistics({ tracks, name }) {
    const data = []
    const [open, setOpen] = useState(false);
    const [fontSize, setFontSize] = useState(0)

    const showDrawer = (length) => {
        decideLabelFontSize(length)
        setOpen(true);
    };

    const onClose = () => {
        setOpen(false);
    };

    const decideLabelFontSize = (length) => {
        if (length <= 10) {
            setFontSize(24)
        }
        else if (length <= 40) {
            setFontSize(20)
        }
        else if (length <= 60) {
            setFontSize(14)
        }
        else if (length <= 80) {
            setFontSize(12)
        }
        else {
            setFontSize(10)
        }
    }

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
            const mostPopularTracks = data.slice(-5).sort((t1, t2) => t2.y - t1.y)

            return (
                <Col style={{ "textAlign": "start"}}>
                    <Row>
                        <Col style={{ "paddingRight": "20px" }}>
                            <h3>5 Most Popular Tracks</h3>
                            {mostPopularTracks.map((track, index) => <p>{index + 1}.{track.x}</p>)}
                            <Link target="_blank" onClick={() => showDrawer(length)}>
                                Click to view details.
                            </Link>
                            <Drawer title={`Popularity statistics of ${name}`}
                                open={open} onClose={onClose}
                                placement="bottom" height={800}
                                footer={`Mean Value = ${meanPopularity.toFixed(2)}`}
                                style={{ overflowY: "scroll" }}
                            >
                                {length > 50 ? <Text italic>{length} tracks... You listen to all of these? Really?</Text> : <Text italic>{length} tracks</Text>}
                                <PopularityStatisticVisualization data={data} width={700} height={900} domainPadding={10} fontSize={fontSize} />
                            </Drawer>
                        </Col>
                        <Col>
                            <h3>5 Least Popular Tracks</h3>
                            {leastPopularTracks.map((track, index) => <p>{index + 1}.{track.x}</p>)}
                            {length >= 80 && <Text type="danger">Playlists containing 80 or more tracks might not render properly.</Text>}
                        </Col>
                    </Row>
                </Col>
            )
        }
        else {
            data.sort((t1, t2) => t2.y - t1.y)
            return (
                <Col>
                    <Row>
                        <h3>Most Popular Tracks</h3>
                    </Row>
                    <Row>
                        <Col style={{ "textAlign": "start", "paddingRight": "20px" }}>
                            {data.map((track, index) => {
                                if (index + 1 <= 5)
                                    return <p>{index + 1}.{track.x}</p>
                            })}
                            <Link target="_blank" onClick={() => showDrawer(length)}>
                                Click to view details
                            </Link>
                            <Drawer title="Track popularity" open={open} onClose={onClose} placement="bottom" footer={`Mean Value = ${meanPopularity.toFixed(2)}`} height={600}>
                                <Text italic>Just {length} tracks? Is this a phase?</Text>
                                <PopularityStatisticVisualization data={data.sort((t1, t2) => t1.y - t2.y)} width={400} height={600} domainPadding={15} fontSize={fontSize} />
                            </Drawer>
                        </Col>
                        <Col style={{ "textAlign": "start" }}>
                            {data.map((track, index) => {
                                if (index + 1 > 5)
                                    return <p>{index + 1}.{track.x}</p>
                            })}
                        </Col>
                    </Row>
                </Col>
            )
        }
    }
}