package repository

import (
	"fmt"
	"github.com/MKA-Nigeria/mkanmedia-go/config"
	mgo "go.mongodb.org/mongo-driver/mongo"
	. "net/http"
	"strconv"
	"time"

	dbs "github.com/MKA-Nigeria/mkanmedia-go/config/db"
	httplib "github.com/MKA-Nigeria/mkanmedia-go/config/http"
	. "github.com/MKA-Nigeria/mkanmedia-go/config/responses"
	"github.com/MKA-Nigeria/mkanmedia-go/controllers/tracks"
	"github.com/MKA-Nigeria/mkanmedia-go/models"
)

type TracksRepository interface {
	RefreshAudioData(writer ResponseWriter, _ *Request)
	GetAllTracks(writer ResponseWriter, request *Request)
	GetAllPlaylists(writer ResponseWriter, _ *Request)
	GetCurrentAuthToken(writer ResponseWriter, request *Request)
	GetRecommendedMedia(writer ResponseWriter, request *Request)
	RefreshRecommendedMedia(writer ResponseWriter, request *Request)
}

type SoundCloudTracksRepository struct {
	remote      SoundCloudRemote
	store       SoundCloudStore
	auth        SoundCloudAuth
	recommender AudioRecommendationEngine
}

func NewSoundCloudRepository() TracksRepository {
	client := Client{Timeout: time.Duration(1) * time.Second}
	var db *mgo.Client
	env := config.Env

	db = dbs.ConnectMongodb()

	return SoundCloudTracksRepository{
		auth:        tracks.AuthImpl{Client: &client, Cache: dbs.ConnectRedis(), SoundCloudClientId: env.SOUND_CLOUD_CLIENT_ID, SoundCloudClientSecret: env.SOUND_CLOUD_CLIENT_SECRET},
		store:       tracks.StoreImpl{TracksCollection: db.Database("mkan-media").Collection("tracks"), PlaylistCollection: db.Database("mkan-media").Collection("playlists")},
		remote:      tracks.RemoteImpl{Client: &client, TracksStartUrl: tracks.TrackStartPoint, PlaylistsStartUrl: tracks.PlaylistStartPoint},
		recommender: tracks.RecommenderImpl{Cache: dbs.ConnectRedis()},
	}
}

func (repository SoundCloudTracksRepository) GetRecommendedMedia(writer ResponseWriter, _ *Request) {
	recommendedTracks, err := repository.recommender.FetchRecommendations()
	if respondedWithError(writer, err) {
		return
	}

	playLists, err := repository.store.FetchSavedPlaylists(true)
	if respondedWithError(writer, err) {
		return
	}

	httplib.ResponseJSON(writer, StatusOK, GeneralResponse{
		Success: true,
		Data: models.Recommendations{
			Playlists:         playLists,
			RecommendedTracks: recommendedTracks,
		},
		Error: nil,
	})
}

func (repository SoundCloudTracksRepository) RefreshRecommendedMedia(writer ResponseWriter, _ *Request) {
	recommendedTracks, err := repository.store.FetchSavedTracks("")
	if respondedWithError(writer, err) {
		return
	}

	count, err := repository.recommender.GenerateRecommendations(recommendedTracks)
	if respondedWithError(writer, err) {
		return
	}

	httplib.ResponseJSON(writer, StatusOK, GeneralResponse{
		Success: true,
		Data:    nil,
		Error:   nil,
		Message: "Generated : " + strconv.Itoa(count) + " recommended tracks",
	})
}

func (repository SoundCloudTracksRepository) GetAllTracks(writer ResponseWriter, request *Request) {
	playlistId := request.FormValue("playlistId")
	trackList, err := repository.store.FetchSavedTracks(playlistId)

	if respondedWithError(writer, err) {
		return
	}

	httplib.ResponseJSON(writer, StatusOK, GeneralResponse{
		Success: true,
		Data:    trackList,
		Error:   nil,
	})
}

func (repository SoundCloudTracksRepository) GetAllPlaylists(writer ResponseWriter, _ *Request) {
	playLists, err := repository.store.FetchSavedPlaylists(false)

	fmt.Println("Repo Got all tracks: size", len(playLists))

	if respondedWithError(writer, err) {
		return
	}

	httplib.ResponseJSON(writer, StatusOK, GeneralResponse{
		Success: true,
		Data:    playLists,
		Error:   nil,
	})
}

func (repository SoundCloudTracksRepository) getToken() (string, *error) {
	token, err := repository.auth.GetToken()
	return token.AccessToken, err
}

func (repository SoundCloudTracksRepository) RefreshAudioData(writer ResponseWriter, _ *Request) {
	token, _ := repository.getToken()

	trackCount, err := repository.refreshTrackData(token)
	if respondedWithError(writer, err) {
		return
	}

	playlistCount, err := repository.refreshPlaylistData(token)
	if respondedWithError(writer, err) {
		return
	}

	httplib.ResponseJSON(writer, StatusOK,
		GeneralResponse{
			Success: true,
			Data:    nil,
			Error:   nil,
			Message: "Stored -> tracks: size " + strconv.Itoa(trackCount) + "\n Playlists: size " + strconv.Itoa(playlistCount),
		})
}

func respondedWithError(writer ResponseWriter, err *error) bool {
	if err != nil && *err != nil {
		httplib.ResponseJSON(writer, StatusInternalServerError, GeneralResponse{
			Success: false,
			Data:    nil,
			Error:   *err,
			Message: fmt.Sprintf("error: %v", *err),
		})
		fmt.Println("Error and then return", *err)
		return true
	}

	return false
}

func (repository SoundCloudTracksRepository) refreshTrackData(token string) (int, *error) {
	trackList, _ := repository.remote.FetchAllTracks(token)

	err0 := repository.store.ClearTracks()
	if err0 != nil && *err0 != nil {
		println("Error occurred while clearing tracks", *err0)
	}
	err1 := repository.store.SaveTracks(trackList)
	if err1 != nil && *err1 != nil {
		println("Error occurred while saving tracks", *err1)
	}

	return len(trackList), err1
}

func (repository SoundCloudTracksRepository) refreshPlaylistData(token string) (int, *error) {
	playlists, _ := repository.remote.FetchAllPlaylists(token)

	err0 := repository.store.ClearPlaylists()
	if err0 != nil && *err0 != nil {
		println("Error occurred while saving tracks", *err0)
	}

	err1 := repository.store.SavePlaylists(playlists)
	if err1 != nil && *err1 != nil {
		println("Error occurred while saving tracks", *err1)
	}

	return len(playlists), err1
}

func (repository SoundCloudTracksRepository) GetCurrentAuthToken(writer ResponseWriter, _ *Request) {
	token, err := repository.auth.GetToken()
	if respondedWithError(writer, err) {
		return
	}

	httplib.ResponseJSON(writer, StatusOK, GeneralResponse{
		Success: true,
		Data:    token,
		Error:   nil,
	})
}
