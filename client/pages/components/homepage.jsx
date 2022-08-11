import { useSession, signIn, signOut } from "next-auth/react"
import ProfileInformation from "./profileInformation"
import PlaylistInfromation from "./playlistInformation"

export default function Homepage() {
    const { data: session } = useSession()

    if (session) {
        return (
            <div className="signed-in-home">
                <div className="signed-in-header">
                    <ProfileInformation />
                    <button onClick={() => signOut('spotify')}>Sign out</button>
                </div>
                <PlaylistInfromation />
            </div>
        )
    } 
    return (
        <div className="not-signed-in-div">
            <h2 className="welcome-text">Welcome to MySpotify</h2>
            <p>Sign in to start</p>
            <button onClick={() => signIn('spotify')} id="sign-out-button">Sign in with Spotify</button>
        </div>
    )
}