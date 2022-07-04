package models

import (
	"encoding/json"
)

type TracksResponse struct {
	Tracks  []Track `json:"collection"`
	NextUrl string  `json:"next_href"`
}

type Track struct {
	Kind          string      `json:"kind"`
	Id            int         `json:"id"`
	CreatedAt     string      `json:"created_at"`
	Duration      int         `json:"duration"`
	Commentable   bool        `json:"commentable"`
	CommentCount  int         `json:"comment_count"`
	Sharing       string      `json:"sharing"`
	TagList       string      `json:"tag_list"`
	Streamable    bool        `json:"streamable"`
	EmbeddableBy  string      `json:"embeddable_by"`
	PurchaseUrl   interface{} `json:"purchase_url"`
	PurchaseTitle interface{} `json:"purchase_title"`
	Genre         *string     `json:"genre"`
	Title         string      `json:"title"`
	Description   *string     `json:"description"`
	LabelName     interface{} `json:"label_name"`
	Release       interface{} `json:"release"`
	KeySignature  interface{} `json:"key_signature"`
	Isrc          interface{} `json:"isrc"`
	Bpm           interface{} `json:"bpm"`
	ReleaseYear   interface{} `json:"release_year"`
	ReleaseMonth  interface{} `json:"release_month"`
	ReleaseDay    interface{} `json:"release_day"`
	License       string      `json:"license"`
	Uri           string      `json:"uri"`
	User          struct {
		AvatarUrl            string      `json:"avatar_url"`
		Id                   int         `json:"id"`
		Kind                 string      `json:"kind"`
		PermalinkUrl         string      `json:"permalink_url"`
		Uri                  string      `json:"uri"`
		Username             string      `json:"username"`
		Permalink            string      `json:"permalink"`
		CreatedAt            string      `json:"created_at"`
		LastModified         string      `json:"last_modified"`
		FirstName            string      `json:"first_name"`
		LastName             string      `json:"last_name"`
		FullName             string      `json:"full_name"`
		City                 string      `json:"city"`
		Description          string      `json:"description"`
		Country              string      `json:"country"`
		TrackCount           int         `json:"track_count"`
		PublicFavoritesCount int         `json:"public_favorites_count"`
		RepostsCount         int         `json:"reposts_count"`
		FollowersCount       int         `json:"followers_count"`
		FollowingsCount      int         `json:"followings_count"`
		Plan                 string      `json:"plan"`
		MyspaceName          interface{} `json:"myspace_name"`
		DiscogsName          interface{} `json:"discogs_name"`
		WebsiteTitle         string      `json:"website_title"`
		Website              string      `json:"website"`
		CommentsCount        int         `json:"comments_count"`
		Online               bool        `json:"online"`
		LikesCount           int         `json:"likes_count"`
		PlaylistCount        int         `json:"playlist_count"`
		Subscriptions        []struct {
			Product struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"product"`
		} `json:"subscriptions"`
	} `json:"user"`
	PermalinkUrl          string      `json:"permalink_url"`
	ArtworkUrl            string      `json:"artwork_url"`
	StreamUrl             string      `json:"stream_url"`
	DownloadUrl           interface{} `json:"download_url"`
	WaveformUrl           string      `json:"waveform_url"`
	AvailableCountryCodes interface{} `json:"available_country_codes"`
	SecretUri             interface{} `json:"secret_uri"`
	UserFavorite          interface{} `json:"user_favorite"`
	UserPlaybackCount     interface{} `json:"user_playback_count"`
	PlaybackCount         int         `json:"playback_count"`
	DownloadCount         int         `json:"download_count"`
	FavoritingsCount      int         `json:"favoritings_count"`
	RepostsCount          int         `json:"reposts_count"`
	Downloadable          bool        `json:"downloadable"`
	Access                string      `json:"access"`
	Policy                interface{} `json:"policy"`
	MonetizationModel     interface{} `json:"monetization_model"`
}

type TracksList struct {
	Tracks []Track `json:"collection"`
}

func (receiver TracksList) ToJson() ([]byte, *error) {
	data, err := json.Marshal(receiver)
	return data, &err
}

func (receiver TracksList) TracksListFromJson(tracksJson string) (TracksList, *error) {
	err := json.Unmarshal([]byte(tracksJson), &receiver)
	return receiver, &err
}
