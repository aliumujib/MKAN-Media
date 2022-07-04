package scheduling

import (
	"fmt"
	"github.com/MKA-Nigeria/mkanmedia-go/controllers/tracks/repository"
	"github.com/go-co-op/gocron"
	"time"
)

type MediaRefresher interface {
	ScheduleMediaRefreshing()
}

type MediaRefresherImpl struct {
	Scheduler        *gocron.Scheduler
	TracksRepository repository.TracksRepository
}

func NewMediaRefresher(scheduler *gocron.Scheduler, tracksRepository repository.TracksRepository) MediaRefresher {
	return MediaRefresherImpl{
		Scheduler:        scheduler,
		TracksRepository: tracksRepository,
	}
}

func (refresher MediaRefresherImpl) ScheduleMediaRefreshing() {
	refresher.Scheduler.Every(24).Hours().Do(func() {
		fmt.Println("Running scheduler tasks at ", time.Now().Format(time.RFC850))
		refresher.TracksRepository.RefreshTrackData()
		refresher.TracksRepository.RefreshPlaylistData()
		refresher.TracksRepository.RefreshRecommendedMedia()
	})

	refresher.Scheduler.StartAsync()
}
