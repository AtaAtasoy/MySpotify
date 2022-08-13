/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  images: {
    domains: ["i.scdn.co"]
  },
  env:{
    backendUrl: process.env.BACKEND_URL
  }
}

module.exports = nextConfig