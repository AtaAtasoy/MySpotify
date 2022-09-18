import NextAuth from "next-auth/next"
import SpotifyProvider from "next-auth/providers/spotify"

export default NextAuth({
  // Configure one or more authentication providers
  providers: [
    SpotifyProvider({
      clientId: process.env.SPOTIFY_ID,
      clientSecret: process.env.SPOTIFY_SECRET,
      authorization: 'https://accounts.spotify.com/authorize?scope=' + process.env.SPOTIFY_SCOPE
    }),
    // ...add more providers here
  ],
  callbacks: {
    async jwt({token, account, user}) {
      if (account && user) {
        token.accessToken = account.access_token
        token.uid = user.id
      }
      return token
      },
    async session({ session, token }) {
        // Send properties to the client, like an access_token from a provider.
        session.accessToken = token.accessToken
        session.userId = token.uid
        return session
      }
  }
})