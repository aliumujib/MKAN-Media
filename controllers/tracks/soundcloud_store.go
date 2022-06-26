package tracks

import (
	"errors"
	"github.com/MKA-Nigeria/mkanmedia-go/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type StoreImpl struct {
	//http client
	TracksCollection   *mgo.Collection
	PlaylistCollection *mgo.Collection
}

func (store StoreImpl) FetchSavedPlaylists(recent bool) ([]models.Playlist, *error) {
	var playlists []models.Playlist
	var err error

	if recent {
		err_ := store.PlaylistCollection.Find(bson.M{}).Sort("-createdat").Limit(10).All(&playlists)
		err = err_
	} else {
		err_ := store.PlaylistCollection.Find(bson.M{}).All(&playlists)
		err = err_
	}

	return playlists, &err
}

func (store StoreImpl) SavePlaylists(playlists []models.Playlist) *error {
	for _, track := range playlists {
		err := store.PlaylistCollection.Insert(track)
		if err != nil {
			return &err
		}
	}
	return nil
}

func (store StoreImpl) ClearTracks() *error {
	err := store.TracksCollection.DropCollection()
	return &err
}

func (store StoreImpl) ClearPlaylists() *error {
	err := store.PlaylistCollection.DropCollection()
	return &err
}

func (store StoreImpl) getTrackIds(playlistId string) ([]int, *error) {
	trackIds := make([]int, 0)

	var playlist []models.Playlist
	playlistId_, _ := strconv.Atoi(playlistId)
	err := store.PlaylistCollection.Find(bson.M{"id": playlistId_}).All(&playlist)

	if err != nil || len(playlist) == 0 {
		err = errors.New("No playlist found with id: " + playlistId)
		return nil, &err
	}

	for _, trackId := range playlist[0].Tracks {
		trackIds = append(trackIds, trackId.Id)
	}

	return trackIds, nil
}

func (store StoreImpl) FetchSavedTracks(playlistId string) ([]models.Track, *error) {
	var tracks []models.Track

	if len(playlistId) > 0 {
		trackIdList, err := store.getTrackIds(playlistId)
		if err != nil {
			return nil, err
		}

		store.TracksCollection.Find(bson.M{"id": bson.M{"$in": trackIdList}}).All(&tracks)
		return tracks, nil
	}

	err := store.TracksCollection.Find(bson.M{}).All(&tracks)
	return tracks, &err
}

func (store StoreImpl) SaveTracks(tracks []models.Track) *error {
	for _, track := range tracks {
		err := store.TracksCollection.Insert(track)
		if err != nil {
			return &err
		}
	}
	return nil
}
