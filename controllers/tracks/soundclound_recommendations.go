package tracks

import (
	"errors"
	"github.com/MKA-Nigeria/mkanmedia-go/models"
	"github.com/go-redis/redis"
	"math/rand"
	"time"
)

type RecommenderImpl struct {
	//http client
	Cache *redis.Client
}

func (recommender RecommenderImpl) GenerateRecommendations(tracks []models.Track) (int, *error) {
	rand.Seed(time.Now().Unix())
	var randomTracks []models.Track

	for turn := 1; turn <= 10; turn++ {
		n := rand.Int() % len(tracks)
		randomTracks = append(randomTracks, tracks[n])
	}

	tracksList := models.TracksList{
		Tracks: randomTracks,
	}

	data, err := tracksList.ToJson()
	err_ := recommender.Cache.Set(RecommendedTracksKey, data, RecommenderCacheTimeOut).Err()
	if err != nil && err_ != nil {
		return 0, &err_
	}

	return len(tracksList.Tracks), &err_
}

func (recommender RecommenderImpl) FetchRecommendations() ([]models.Track, *error) {
	var tracksList models.TracksList
	result, _ := recommender.Cache.Get(RecommendedTracksKey).Result()

	if len(result) == 0 {
		err := errors.New("no tracks to recommend, please contact admin")
		return nil, &err
	}

	tracksList, err := tracksList.TracksListFromJson(result)

	if *err != nil {
		err := errors.New("no tracks to recommend, please contact admin")
		return nil, &err
	}

	return tracksList.Tracks, nil
}
