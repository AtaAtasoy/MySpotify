export default function Playlists(){
    const BACKEND_URI = process.env.BACKEND_URI

    //TODO:Implement the request
    const getUserPlaylists = () => {
        fetch(BACKEND_URI + 'me/favorite/artists')
    }

    return(
        <div className="playlists-container">
            <button onClick={() => getUserPlaylists()}>Display Playlists</button>
        </div>
    )
}