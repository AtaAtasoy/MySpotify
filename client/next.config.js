/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  images: {
    remotePatterns: [
      {
        // The `src` property hostname must end with `.example.com`,
        // otherwise the API will respond with 400 Bad Request.
        protocol: 'https',
        hostname: '**.scdn.co'
      },
      {
        // The `src` property hostname must end with `.example.com`,
        // otherwise the API will respond with 400 Bad Request.
        protocol: 'https',
        hostname: '**.spotifycdn.com'
      },
    ],
  },
  env:{
    backendUrl: process.env.BACKEND_URL
  },
}

module.exports = nextConfig