import { useSession } from "next-auth/react"
import { useState } from "react"
import CircleLoader from "./components/CircleLoader"
import Playlist from "./components/Playlist"

export default function Playlists() {
    const { data: session } = useSession()
    const url = process.env.PLAYLISTS_SERVER_URI + '/playlists'
    const [playlists, setPlaylists] = useState([])
    const [fetching, setFetching] = useState(false)

    if (session) {
        const options = {
            method: "GET",
            headers: new Headers({
                'Authorization': session.accessToken,
                'Content-Type': 'application/json',
                'Username': session.userId
            }),
        }

        const getUserPlaylists = () => {
            playlists.length = 0
            console.log(url)
            setFetching(true)
            fetch(url, options)
                .then(response => {
                    response.json().then(json => {
                        setPlaylists(playlists.concat(json))
                        console.log(json)
                    })
                })
                .finally(() => setFetching(false))
                .catch(err => console.error(err))
        }

        return (
            <div className="fetch-playlists-container">
                <button onClick={() => getUserPlaylists()}>Display Playlists</button>
                <div className="playlists-child-container">
                    {fetching ? <CircleLoader /> : playlists.map((playlistData, i) => {
                        return(
                        <div className="playlist-container" key={i}> 
                            <Playlist key={i} name={playlistData.name} tracks={playlistData.tracks} image={playlistData.images[0]} />
                        </div>
                        ) 
                    })}
                </div>
            </div>
        )
    }
}