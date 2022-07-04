package tracks

import (
	"encoding/json"
	"errors"
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
	url := remote.PlaylistsStartUrl

	for &url != nil && len(url) > 0 {
		fmt.Println("Current next is " + url)
		nextResponse, err := remote.fetchPlaylistsFromSoundCloud(url, accessToken)
		playlists = append(playlists, nextResponse.Playlists...)
		if err != nil {
			return nil, err
		}
		if nextResponse.NextUrl == nil || len(*nextResponse.NextUrl) == 0 {
			break
		}

		url = *nextResponse.NextUrl
	}

	fmt.Println("Returning all playlists with size ", len(playlists))
	return playlists, nil
}

func (remote RemoteImpl) FetchAllTracks(accessToken string) ([]models.Track, *error) {
	var allTracks []models.Track
	url := remote.TracksStartUrl

	for &url != nil && len(url) > 0 {
		fmt.Println("Current next is " + url)
		nextResponse, err := remote.fetchTracksFromSoundCloud(url, accessToken)
		allTracks = append(allTracks, nextResponse.Tracks...)
		if *err != nil {
			return nil, err
		}
		if nextResponse.NextUrl == nil || len(*nextResponse.NextUrl) == 0 {
			break
		}

		url = *nextResponse.NextUrl
	}

	fmt.Println("Returning all tracks with size ", len(allTracks))
	return allTracks, nil
}

func cleanup(data interface{}) (interface{}, *error) {
	var err error
	if r, ok := recover().(error); ok {
		fmt.Println("An error occurred", r)
		err = errors.New("internal error")
		return nil, &err
	}
	return data, nil
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
