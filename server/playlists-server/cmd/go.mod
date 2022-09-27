module playlists-server/cmd

go 1.18

replace playlists-server/playlists => ../pkg

replace tracks-server/tracks => ../../tracks-server/pkg

require (
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.4.0
	github.com/rs/cors v1.8.2
	playlists-server/playlists v0.0.0-00010101000000-000000000000
)

require tracks-server/tracks v0.0.0-00010101000000-000000000000 // indirect
