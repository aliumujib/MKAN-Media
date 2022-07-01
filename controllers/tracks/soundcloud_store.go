package tracks

import (
	"context"
	"errors"
	"fmt"
	"github.com/MKA-Nigeria/mkanmedia-go/models"
	"go.mongodb.org/mongo-driver/bson"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
)

type StoreImpl struct {
	//http client
	TracksCollection   *mgo.Collection
	PlaylistCollection *mgo.Collection
	Context            context.Context
}

func (store StoreImpl) FetchSavedPlaylists(recent bool) ([]models.Playlist, *error) {
	var playlists []models.Playlist
	var err error

	if recent {
		sortOptions := options.Find()
		sortOptions.SetSort(bson.D{{"createdat", -1}})
		sortOptions.SetLimit(10)

		cursor, err_ := store.PlaylistCollection.Find(store.Context, bson.D{}, sortOptions)
		if err_ != nil {
			return nil, &err_
		}
		err = cursor.All(store.Context, &playlists)
	} else {
		cursor, err_ := store.PlaylistCollection.Find(store.Context, bson.D{})
		if err_ != nil {
			return nil, &err_
		}
		err = cursor.All(store.Context, &playlists)
	}

	return playlists, &err
}

func (store StoreImpl) SavePlaylists(playlists []models.Playlist) *error {
	for _, track := range playlists {
		_, err := store.PlaylistCollection.InsertOne(store.Context, track)
		if err != nil {
			return &err
		}
	}
	return nil
}

func (store StoreImpl) ClearTracks() *error {
	err := store.TracksCollection.Drop(store.Context)
	return &err
}

func (store StoreImpl) ClearPlaylists() *error {
	err := store.PlaylistCollection.Drop(store.Context)
	return &err
}

func (store StoreImpl) getTrackIds(playlistId string) ([]int, *error) {
	trackIds := make([]int, 0)

	var playlist []models.Playlist
	playlistId_, _ := strconv.Atoi(playlistId)
	cursor, err := store.PlaylistCollection.Find(store.Context, bson.M{"id": playlistId_})
	if err != nil {
		err = errors.New("Error finding playlist with Id: " + playlistId)
		return nil, &err
	}

	err = cursor.All(store.Context, &playlist)
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
	tracks := make([]models.Track, 0)

	if len(playlistId) > 0 {
		trackIdList, err := store.getTrackIds(playlistId)
		fmt.Println("Playlist Id ", playlistId)

		if err != nil {
			return nil, err
		}

		cursor, err_ := store.TracksCollection.Find(store.Context, bson.M{"id": bson.M{"$in": trackIdList}})
		if err_ != nil {
			err_ = errors.New("No playlist found with id: " + playlistId)
			return nil, &err_
		}

		err_ = cursor.All(store.Context, &tracks)

		return tracks, &err_
	}

	cursor, err := store.TracksCollection.Find(store.Context, bson.M{})
	if err != nil {
		return nil, &err
	}

	err = cursor.All(store.Context, &tracks)

	return tracks, &err
}

func (store StoreImpl) SaveTracks(tracks []models.Track) *error {
	for _, track := range tracks {
		_, err := store.TracksCollection.InsertOne(store.Context, track)
		if err != nil {
			return &err
		}
	}
	return nil
}
