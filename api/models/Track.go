package models

type Track struct{
	Id string
	Name string
	Popularity float64
	Artists []Artist
}