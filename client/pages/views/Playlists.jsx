import { useSession } from "next-auth/react"

export default function Playlists() {
    const { data: session } = useSession()
    const url = process.env.backendUrl + '/playlists?limit=3'

    if (session) {
        const options = {
            method: "GET",
            headers: new Headers({
                'Authorization': session.accessToken
            }),
        }

        //TODO:Implement the request
        const getUserPlaylists = () => {
            fetch(encodeURI(url), options)
                .then(response => response.json())
                .then(jsondata => console.log(jsondata))
                .catch(err => console.error(err))
        }

        return (
            <div className="playlists-container">
                <button onClick={() => getUserPlaylists()}>Display Playlists</button>
            </div>
        )
    }
}