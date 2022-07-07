package models

type User struct {
	AvatarUrl     string `json:"avatar_url"`
	Id            int    `json:"id"`
	Kind          string `json:"kind"`
	PermalinkUrl  string `json:"permalink_url"`
	Uri           string `json:"uri"`
	Username      string `json:"username"`
	Permalink     string `json:"permalink"`
	CreatedAt     string `json:"created_at"`
	FullName      string `json:"full_name"`
	City          string `json:"city"`
	Description   string `json:"description"`
	Country       string `json:"country"`
	TrackCount    int    `json:"track_count"`
	PlaylistCount int    `json:"playlist_count"`
}
