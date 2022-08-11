import { useSession } from "next-auth/react"
import Image from "next/image"

export default function UserProfile (){
    const { data: session } = useSession()
    return(
        <div className="profile-information">
            <Image className="profile-picture" alt="profile-picture" src={session.user.image} width={150} height={150} />
            <h2 className="user-name">{session.user.name}</h2>
        </div>
    )
}