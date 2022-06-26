package repository

import (
	"github.com/MKA-Nigeria/mkanmedia-go/models"
)

type SoundCloudRemote interface {
	FetchAllTracks(accessToken string) ([]models.Track, *error)
	FetchAllPlaylists(accessToken string) ([]models.Playlist, *error)
}

type AudioRecommendationEngine interface {
	GenerateRecommendations([]models.Track) (int, *error)
	FetchRecommendations() ([]models.Track, *error)
}

type SoundCloudStore interface {
	FetchSavedTracks(string) ([]models.Track, *error)
	SaveTracks([]models.Track) *error
	FetchSavedPlaylists(recents bool) ([]models.Playlist, *error)
	SavePlaylists([]models.Playlist) *error
	ClearTracks() *error
	ClearPlaylists() *error
}

type SoundCloudAuth interface {
	GetToken() (*models.TokenResponse, *error)
}
