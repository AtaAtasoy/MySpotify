import { useSession } from "next-auth/react"
import Image from "next/future/image"
import { Row } from "antd"

export default function UserProfile (){
    const { data: session } = useSession()
    if (session){
        return(
            <Row className="profile-information" align="center">
                <Image className="profile-picture" alt="profile-picture" src={session.user.image} width={150} height={150} />
                <h2 className="user-name">Welcome {session.user.name}</h2>
            </Row>
        )
    } 
}
