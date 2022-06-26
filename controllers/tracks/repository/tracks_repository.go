package repository

import (
	"fmt"
	dbs "github.com/MKA-Nigeria/mkanmedia-go/config/db"
	httplib "github.com/MKA-Nigeria/mkanmedia-go/config/http"
	"github.com/MKA-Nigeria/mkanmedia-go/controllers/tracks"
	"github.com/MKA-Nigeria/mkanmedia-go/models"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	. "net/http"
	"strconv"
	"time"
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
	var db *mgo.Session
	env := viper.GetString("env")

	if env == "prod" {
		db = dbs.ConnectMongodbTLS()
	} else {
		db = dbs.ConnectMongodb()
	}

	return SoundCloudTracksRepository{
		auth:        tracks.AuthImpl{Client: &client, Cache: dbs.ConnectRedis()},
		store:       tracks.StoreImpl{TracksCollection: db.DB("mkan-media").C("tracks"), PlaylistCollection: db.DB("mkan-media").C("playlists")},
		remote:      tracks.RemoteImpl{Client: &client, TracksStartUrl: tracks.TrackStartPoint, PlaylistsStartUrl: tracks.PlaylistStartPoint},
		recommender: tracks.RecommenderImpl{Cache: dbs.ConnectRedis()},
	}
}

func (repository SoundCloudTracksRepository) GetRecommendedMedia(writer ResponseWriter, _ *Request) {
	recommendedTracks, err := repository.recommender.FetchRecommendations()
	if err != nil && *err != nil {
		fmt.Println("Error 1:", *err)
		httplib.ResponseJSON(writer, StatusInternalServerError, *err)
		return
	}

	playLists, err := repository.store.FetchSavedPlaylists(true)
	if err != nil && *err != nil {
		fmt.Println("Error 2:", *err)
		httplib.ResponseJSON(writer, StatusInternalServerError, *err)
		return
	}

	httplib.ResponseJSON(writer, StatusOK, models.Recommendations{
		Playlists:         playLists,
		RecommendedTracks: recommendedTracks,
	})
}

func (repository SoundCloudTracksRepository) RefreshRecommendedMedia(writer ResponseWriter, _ *Request) {
	recommendedTracks, err := repository.store.FetchSavedTracks("")
	if err != nil && *err != nil {
		httplib.ResponseJSON(writer, StatusInternalServerError, *err)
		return
	}

	count, err := repository.recommender.GenerateRecommendations(recommendedTracks)
	if err != nil && *err != nil {
		httplib.ResponseJSON(writer, StatusInternalServerError, *err)
		return
	}

	httplib.ResponseJSON(writer, StatusOK, "Generated : "+strconv.Itoa(count)+" recommended tracks")
}

func (repository SoundCloudTracksRepository) GetAllTracks(writer ResponseWriter, request *Request) {
	playlistId := request.FormValue("playlistId")
	trackList, err := repository.store.FetchSavedTracks(playlistId)

	if err != nil {
		fmt.Println("Error: ", *err)
		httplib.ResponseJSON(writer, StatusInternalServerError, *err)
		return
	}

	httplib.ResponseJSON(writer, StatusOK, trackList)
}

func (repository SoundCloudTracksRepository) GetAllPlaylists(writer ResponseWriter, _ *Request) {
	playLists, err := repository.store.FetchSavedPlaylists(false)

	fmt.Println("Repo Got all tracks: size", len(playLists))

	if *err != nil {
		fmt.Println("Error: ", *err)
		httplib.ResponseJSON(writer, StatusInternalServerError, err)
		return
	}

	httplib.ResponseJSON(writer, StatusOK, playLists)
}

func (repository SoundCloudTracksRepository) getToken() (string, *error) {
	token, err := repository.auth.GetToken()
	return token.AccessToken, err
}

func (repository SoundCloudTracksRepository) RefreshAudioData(writer ResponseWriter, _ *Request) {
	token, _ := repository.getToken()
	trackCount, err := repository.refreshTrackData(token)
	respondIfNeeded(writer, err)

	playlistCount, err := repository.refreshPlaylistData(token)
	respondIfNeeded(writer, err)

	httplib.ResponseJSON(writer, StatusOK, "Stored -> tracks: size "+strconv.Itoa(trackCount)+"\n Playlists: size "+strconv.Itoa(playlistCount))
}

func respondIfNeeded(writer ResponseWriter, err *error) {
	if err != nil {
		httplib.ResponseJSON(writer, StatusInternalServerError, err)
		fmt.Println("Error and then return", *err)
		return
	}
}

func (repository SoundCloudTracksRepository) refreshTrackData(token string) (int, *error) {
	trackList, _ := repository.remote.FetchAllTracks(token)

	_ = repository.store.ClearTracks()
	err1 := repository.store.SaveTracks(trackList)

	return len(trackList), err1
}

func (repository SoundCloudTracksRepository) refreshPlaylistData(token string) (int, *error) {
	playlists, _ := repository.remote.FetchAllPlaylists(token)

	_ = repository.store.ClearPlaylists()
	err2 := repository.store.SavePlaylists(playlists)

	return len(playlists), err2
}

func (repository SoundCloudTracksRepository) GetCurrentAuthToken(writer ResponseWriter, request *Request) {
	token, err := repository.auth.GetToken()
	respondIfNeeded(writer, err)

	httplib.ResponseJSON(writer, StatusOK, token)
}

//home -> playlists and recommended tracks random everyday->()
//getAuthToken

//youtube stuff
