package routes

import (
	httplib "github.com/MKA-Nigeria/mkanmedia-go/config/http"
	responses "github.com/MKA-Nigeria/mkanmedia-go/config/responses"
	"github.com/MKA-Nigeria/mkanmedia-go/controllers/tracks/repository"
	mws "github.com/MKA-Nigeria/mkanmedia-go/middlewares"
	"github.com/gorilla/mux"
	"net/http"
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
	audioRoute := route.PathPrefix("/v1/audio").Subrouter()
	audioRoute.HandleFunc("/refresh-tracks", tracksRepo.RefreshAudioData).Methods("GET")
	audioRoute.HandleFunc("/fetch-all-tracks", tracksRepo.GetAllTracks).Methods("GET")
	audioRoute.HandleFunc("/fetch-all-playlists", tracksRepo.GetAllPlaylists).Methods("GET")
	audioRoute.HandleFunc("/current-auth", tracksRepo.GetCurrentAuthToken).Methods("GET")
	audioRoute.HandleFunc("/recommendations", tracksRepo.GetRecommendedMedia).Methods("GET")
	audioRoute.HandleFunc("/refresh-recommendations", tracksRepo.RefreshRecommendedMedia).Methods("GET")

	return route
}
