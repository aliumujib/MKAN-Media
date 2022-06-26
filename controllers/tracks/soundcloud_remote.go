package tracks

import (
	"encoding/json"
	"fmt"
	"github.com/MKA-Nigeria/mkanmedia-go/models"
	"io/ioutil"
	. "net/http"
)

type RemoteImpl struct {
	//http client
	Client            *Client
	TracksStartUrl    string
	PlaylistsStartUrl string
}

func constructAuthenticatedHeader(request *Request, accessToken string) *Request {
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36")
	request.Header.Add("Authorization", "OAuth "+accessToken)
	return request
}

func (remote RemoteImpl) FetchAllPlaylists(accessToken string) ([]models.Playlist, *error) {
	var playlists []models.Playlist
	response, err := remote.fetchPlaylistsFromSoundCloud(remote.PlaylistsStartUrl, accessToken)
	fmt.Println("Got first playlist list of size: ", len(response.Playlists))
	if *err != nil {
		return nil, err
	}
	playlists = append(response.Playlists)

	for response.NextUrl != nil && len(*response.NextUrl) > 0 {
		fmt.Println("Current next is " + *response.NextUrl)
		nextResponse, err := remote.fetchPlaylistsFromSoundCloud(*response.NextUrl, accessToken)
		playlists = append(playlists, nextResponse.Playlists...)
		if *err != nil {
			return nil, err
		}
		response = nextResponse
	}

	fmt.Println("Returning all playlists with size ", len(playlists))
	return playlists, nil
}

func (remote RemoteImpl) FetchAllTracks(accessToken string) ([]models.Track, *error) {
	var allTracks []models.Track
	response, err := remote.fetchTracksFromSoundCloud(remote.TracksStartUrl, accessToken)
	fmt.Println("Got first track list of size: ", len(response.Tracks))
	if *err != nil {
		return nil, err
	}
	allTracks = append(response.Tracks)

	for response.NextUrl != nil && len(*response.NextUrl) > 0 {
		fmt.Println("Current next is " + *response.NextUrl)
		nextResponse, err := remote.fetchTracksFromSoundCloud(*response.NextUrl, accessToken)
		allTracks = append(allTracks, nextResponse.Tracks...)
		if *err != nil {
			return nil, err
		}
		response = nextResponse
	}

	fmt.Println("Returning all tracks with size ", len(allTracks))
	return allTracks, nil
}

func (remote RemoteImpl) fetchPlaylistsFromSoundCloud(playlistsURl string, accessToken string) (models.PlaylistsResponse, *error) {
	playlists := models.PlaylistsResponse{}

	request, err := NewRequest(MethodGet, playlistsURl, nil)
	if err != nil {
		return playlists, &err
	}
	constructAuthenticatedHeader(request, accessToken)

	resp, err := remote.Client.Do(request)
	if err != nil {
		return playlists, &err
	}
	defer resp.Body.Close()

	if resp.StatusCode == StatusUnauthorized {
	}
	if resp.StatusCode >= StatusBadRequest {
		return playlists, &err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return playlists, &err
	}

	if err := json.Unmarshal(body, &playlists); err != nil {
		return playlists, &err
	}

	return playlists, &err
}

func (remote RemoteImpl) fetchTracksFromSoundCloud(tracksURl string, accessToken string) (models.TracksResponse, *error) {
	tracks := models.TracksResponse{}

	request, err := NewRequest(MethodGet, tracksURl, nil)
	if err != nil {
		fmt.Println("An error occurred while constructing tracks call to tracks")
		return tracks, &err
	}
	constructAuthenticatedHeader(request, accessToken)

	resp, err := remote.Client.Do(request)
	if err != nil {
		fmt.Println("An error occurred while constructing tracks call to tracks", err.Error())
		return tracks, &err
	}
	defer resp.Body.Close()

	if resp.StatusCode == StatusUnauthorized {
		//talk to sandra about handling retries and things like self or this from the JS and Python world
	}

	if resp.StatusCode >= StatusBadRequest {
		fmt.Println("An error occurred while performing request call to sound cloud", err.Error())
		return tracks, &err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("An error occurred while reading response from sound cloud")
		return tracks, &err
	}

	if err := json.Unmarshal(body, &tracks); err != nil {
		fmt.Println("An error occurred while parsing response from sound cloud")
		return tracks, &err
	}

	return tracks, &err
}
