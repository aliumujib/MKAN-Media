package models

type PlaylistsResponse struct {
	Playlists []Playlist `json:"collection"`
	NextUrl   string     `json:"next_href"`
}

type Playlist struct {
	Duration     int     `json:"duration"`
	Description  *string `json:"description"`
	Uri          string  `json:"uri"`
	TagList      string  `json:"tag_list"`
	TrackCount   int     `json:"track_count"`
	UserId       int     `json:"user_id"`
	LastModified string  `json:"last_modified"`
	License      string  `json:"license"`
	User         User    `json:"user"`
	PlaylistType string  `json:"playlist_type"`
	Type         string  `json:"type"`
	Id           int     `json:"id"`
	CreatedAt    string  `json:"created_at"`
	Tags         string  `json:"tags"`
	Kind         string  `json:"kind"`
	Title        string  `json:"title"`
	ArtworkUrl   *string `json:"artwork_url"`
	TracksUri    string  `json:"tracks_uri"`
	Tracks       []struct {
		Id int `json:"id"`
	} `json:"tracks"`
}
