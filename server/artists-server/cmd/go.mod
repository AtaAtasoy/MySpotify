module artist-server/cmd

go 1.18

replace artists-server/artists => ../pkg

require (
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.4.0
	github.com/rs/cors v1.8.2
)

require artists-server/artists v0.0.0-00010101000000-000000000000
