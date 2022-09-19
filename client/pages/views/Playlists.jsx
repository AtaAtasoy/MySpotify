import { useSession } from "next-auth/react"
import { useState } from "react"
import CircleLoader from "./components/CircleLoader"
import Playlist from "./components/Playlist"

export default function Playlists() {
    const { data: session } = useSession()
    const url = process.env.backendUrl + '/playlists'
    const [playlists, setPlaylists] = useState([])
    const [fetching, setFetching] = useState(false)

    const renderTableRows = () =>{
        const rows = []
        for (let i = 0; i < playlists.length - 2; i = i + 3){
            rows.push(
                <tr>
                    <td><Playlist name={playlists[i].name} tracks={playlists[i].tracks} image={playlists[i].images[0]} /></td>
                    <td><Playlist name={playlists[i + 1].name} tracks={playlists[i + 1].tracks} image={playlists[i + 1].images[0]} /></td>
                    <td><Playlist name={playlists[i + 2].name} tracks={playlists[i + 2].tracks} image={playlists[i + 2].images[0]} /></td>
                </tr>
            )
        }
        return rows
    }

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
            <div className="playlists-parent-container">
                <button onClick={() => getUserPlaylists()}>Display Playlists</button>
                <div className="playlists-child-container">
                    {fetching ? <CircleLoader /> : 
                        <table className="playlists-table">
                           {renderTableRows()}
                        </table>
                    }
                </div>
            </div>
        )
    }
}