import { useSession } from "next-auth/react"
import { useState } from "react"
import Artist from "./components/Artist"
import CircleLoader from "./components/CircleLoader"

export default function FavoriteArtists(){
    const { data: session} = useSession()
    const url = process.env.backendUrl + '/artists?limit=3'
    const [artists, setArtists] = useState([])
    const [fetching, setFetching] = useState(false)

    if (session){
        const options = {
            method: "GET",
            headers: new Headers({
                'Authorization': session.accessToken
            })
        }
        
        //TODO:Implement the request
        const getFavoriteArtists = () => {
            setFetching(true)
            fetch(url, options)
            .then(response => response.json().then(json => setArtists(json)))
            .finally(() => setFetching(false))
            .catch(err => console.error(err))
        }

        return(
            <div className="favorite-artists-container">
                <button onClick={() => getFavoriteArtists()}>Display Favorite Artists</button>
                {fetching ? <CircleLoader /> : artists.map((artistData, i) => { if (artistData =! null) return <Artist key={i} name={artistData.name} popularity={artistData.popularity} image={artistData.images[2]}/> })}
            </div>
        )
    } 
}