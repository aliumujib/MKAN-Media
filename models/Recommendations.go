package models

type Recommendations struct {
	Playlists         []Playlist `json:"playlists"`
	RecommendedTracks []Track    `json:"tracks"`
}
