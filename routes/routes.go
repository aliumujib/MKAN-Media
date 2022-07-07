package routes

import (
	httplib "github.com/MKA-Nigeria/mkanmedia-go/config/http"
	responses "github.com/MKA-Nigeria/mkanmedia-go/config/responses"
	"github.com/MKA-Nigeria/mkanmedia-go/controllers/scheduling"
	"github.com/MKA-Nigeria/mkanmedia-go/controllers/tracks/repository"
	mws "github.com/MKA-Nigeria/mkanmedia-go/middlewares"
	"github.com/go-co-op/gocron"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

//Router for all routes
func Router() *mux.Router {
	route := mux.NewRouter()

	//BASE ROUTE
	route.HandleFunc("/v1", func(res http.ResponseWriter, req *http.Request) {
		resp := responses.GeneralResponse{Success: true, Message: "vms  server running....", Data: "vsm SERVER v1.0"}
		httplib.ResponseJSON(res, http.StatusOK, resp)
	})

	route.Use(mws.AccessLogToConsole)

	//************************
	// AUDIO  ROUTES
	//************************
	tracksRepo := repository.NewSoundCloudRepository()
	scheduler := scheduling.NewMediaRefresher(gocron.NewScheduler(time.UTC), tracksRepo)

	audioRoute := route.PathPrefix("/v1/audio").Subrouter()
	audioRoute.HandleFunc("/tracks", tracksRepo.GetAllTracks).Methods("GET")
	audioRoute.HandleFunc("/playlists", tracksRepo.GetAllPlaylists).Methods("GET")
	audioRoute.HandleFunc("/stream-auth", tracksRepo.GetCurrentAuthToken).Methods("GET")
	audioRoute.HandleFunc("/recommendations", tracksRepo.GetRecommendedMedia).Methods("GET")

	scheduler.ScheduleMediaRefreshing()

	return route
}
