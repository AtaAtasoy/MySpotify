import { useSession } from "next-auth/react"
import { useState } from "react"
import CircleLoader from "./components/CircleLoader"
import Playlist from "./components/Playlist"

export default function Playlists() {
    const { data: session } = useSession()
    const url = process.env.backendUrl + '/playlists?limit=1'
    const [playlists, setPlaylists] = useState([])
    const [fetching, setFetching] = useState(false)

    if (session) {
        const options = {
            method: "GET",
            headers: new Headers({
                'Authorization': session.accessToken,
                'Content-Type': 'application/json'
            }),
        }

        //TODO:Implement the request
        const getUserPlaylists = () => {
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
            <div className="playlists-container">
                <button onClick={() => getUserPlaylists()}>Display Playlists</button>
                {fetching ? <CircleLoader /> : playlists.map((playlistData, i) => <Playlist key={i} name={playlistData.name} tracks={playlistData.tracks} image={playlistData.images[0]} />)}
            </div>
        )
    }
}