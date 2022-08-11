import { useState } from "react"

export default function FavoriteArtists(){
    const BACKEND_URI = process.env.BACKEND_URI

    //TODO:Implement the request
    const getFavoriteArtists = () => {
        fetch(BACKEND_URI + 'me/favorite/artists')
    }


    return(
        <div className="favorite-artists-container">
            <button onClick={() => getFavoriteArtists()}>Display Favorite Artists</button>
            
        </div>
    )

}