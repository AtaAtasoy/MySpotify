import { useSession } from "next-auth/react"

export default function FavoriteArtists(){
    const { data: session} = useSession()
    const BACKEND_URI = process.env.backendUrl + 'me/favorite/artists?limit=3'

    if (session){
        const options = {
            headers: {
                'Authorization': session.accessToken
            }
        }
        
        //TODO:Implement the request
        const getFavoriteArtists = () => {
            console.log(options)
            fetch(BACKEND_URI, options)
            .then(result => console.log(result))
            .catch(err => console.error(err))
        }

        return(
            <div className="favorite-artists-container">
                <button onClick={() => getFavoriteArtists()}>Display Favorite Artists</button>
            </div>
        )

    } 
}