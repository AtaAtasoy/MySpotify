import { useSession, signIn, signOut, get } from "next-auth/react"
import Image from "next/image"

export default function LoginButton() {
    const { data: session } = useSession()

  if (session) {
    return (
      <>
        <Image src={session.user.image} width={200} height={200} alt={"user-profile-picture"}/>
        Signed in as {session.user.email}<br />
        With token {session.accessToken} <br />
        <button onClick={() => signOut()}>Sign out</button>
      </>
    )
  }
  return (
    <>
      Not signed in <br />
      <button onClick={() => signIn()}>Sign in</button>
    </>
  )
}