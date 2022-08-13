import { useSession } from "next-auth/react"

export default function FavoriteArtists(){
    const { data: session} = useSession()
    const url = process.env.backendUrl + '/artists?limit=3'

    if (session){
        const options = {
            method: "GET",
            headers: new Headers({
                'Authorization': session.accessToken
            })
        }
        
        //TODO:Implement the request
        const getFavoriteArtists = () => {
            fetch(url, options)
            .then(response => response.json())
            .then(jsondata => console.log(jsondata))
            .catch(err => console.error(err))
        }

        return(
            <div className="favorite-artists-container">
                <button onClick={() => getFavoriteArtists()}>Display Favorite Artists</button>
            </div>
        )
    } 
}