import { Col, Row } from "antd"
import { useSession, signIn, signOut } from "next-auth/react"
import Playlists from "./Playlists"
import UserProfile from "./UserProfile"

export default function Homepage() {
    const { data: session } = useSession()

    if (session) {
        return (
            <div className="signed-in-home">
                <Row className="signed-in-header" align="center">
                    <Col>
                        <UserProfile />
                    </Col>
                    <Col offset={8}>
                        <button className="spotify-themed-button" onClick={() => signOut('spotify')}>Sign out</button>
                    </Col>
                </Row>
                <Playlists />
            </div>
        )
    }
    return (
        <div className="not-signed-in-div">
            <h2 className="welcome-text">Welcome to MySpotify</h2>
            <p>Sign in to start</p>
            <button className="spotify-themed-button" onClick={() => signIn('spotify')} id="sign-out-button">Sign in with Spotify</button>
        </div>
    )
}