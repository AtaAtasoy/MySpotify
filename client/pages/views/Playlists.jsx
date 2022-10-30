import { useSession } from "next-auth/react"
import { useState, useEffect } from "react"
import CircleLoader from "./components/CircleLoader"
import Playlist from "./components/Playlist"

export default function Playlists() {
    const { data: session } = useSession()
    const [playlists, setPlaylists] = useState([])
    const [fetching, setFetching] = useState(false)
    const [error, setError] = useState('')

    useEffect(() => {
        const item = JSON.parse(localStorage.getItem('playlists'))
        if (item){
            setPlaylists(item)
        }
    })

    if (session) {
        const options = {
            method: "GET",
            headers: {
                'Authorization': session.accessToken,
                'Content-Type': 'application/json',
                'Username': session.userId
            },
        }

        const getUserPlaylists = () => {
            if (!JSON.parse(localStorage.getItem('playlists'))){
                playlists.length = 0
                setFetching(true)
                fetch(process.env.NEXT_PUBLIC_PLAYLISTS_SERVER_URI + "/playlists", options)
                    .then(response => {
                        if (response.status === 400) {
                            throw new Error('Could not fetch playlists.' + response.statusText);
                        }
                        response.json().then(json => {
                            localStorage.setItem('playlists', JSON.stringify(json))
                            setPlaylists(playlists.concat(json))
                        })
                    })
                    .finally(() => setFetching(false))
                    .catch(err => {
                        console.error(err)
                        setError(err)
                    })
            }
        }
        if (error.length === 0) {
            return (
                <div className="fetch-playlists-container">
                    <button disabled={fetching} onClick={() => getUserPlaylists()}>{fetching ? "Loading" : "Display Playlists"}</button>
                    <div className="playlists-child-container">
                        {fetching ? <CircleLoader /> : playlists.map((playlistData, i) => {
                            return (
                                <div className="playlist-container" key={i}>
                                    <Playlist key={i} name={playlistData.name} tracks={playlistData.tracks} image={playlistData.images[0]} attributes={playlistData.attributes} />
                                </div>
                            )
                        })}
                    </div>
                </div>
            )
        }
        else {
            return (
                <div className="fetch-playlists-container">
                    <button disabled={fetching} onClick={() => getUserPlaylists()}>{fetching ? "Loading" : "Display Playlists"}</button>
                    <p>{error.toString()}</p>
                    <p>Signing out and signing back in can solve your problem...</p>
                </div>
            )
        }
    }
}