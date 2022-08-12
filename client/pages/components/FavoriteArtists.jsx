import { useSession } from "next-auth/react"

export default function FavoriteArtists(){
    const { data } = useSession()
    const { accessToken } = data

    const BACKEND_URI = process.env.BACKEND_URI
    const options = {
        headers: {
            'Authorization': accessToken
        }
    }

    //TODO:Implement the request
    const getFavoriteArtists = () => {
        fetch(BACKEND_URI + 'me/favorite/artists?limit=3', options)
        .then(result => console.log(result))
        .catch(err => console.error(err))
    }


    return(
        <div className="favorite-artists-container">
            <button onClick={() => getFavoriteArtists()}>Display Favorite Artists</button>
        </div>
    )

}