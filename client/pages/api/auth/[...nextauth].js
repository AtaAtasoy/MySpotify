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
    async jwt({token, account}) {
      if (account) {
        token.accessToken = account.access_token
      }
      return token
      },
    async session({ session, token }) {
        // Send properties to the client, like an access_token from a provider.
        session.accessToken = token.accessToken
        return session
      }
  }
})