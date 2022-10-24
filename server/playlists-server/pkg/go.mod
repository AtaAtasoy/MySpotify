module playlists-server/playlists

go 1.18

replace tracks-server/tracks => ../../tracks-server/pkg

require tracks-server/tracks v0.0.0-00010101000000-000000000000

require github.com/shopspring/decimal v1.3.1 // indirect
