package models

import (
	"encoding/json"
)

type TracksResponse struct {
	Tracks  []Track `json:"collection"`
	NextUrl string  `json:"next_href"`
}

type Track struct {
	Kind         string      `json:"kind"`
	Id           int         `json:"id"`
	CreatedAt    string      `json:"created_at"`
	Duration     int         `json:"duration"`
	Genre        *string     `json:"genre"`
	Title        string      `json:"title"`
	Uri          string      `json:"uri"`
	User         User        `json:"user"`
	PermalinkUrl string      `json:"permalink_url"`
	ArtworkUrl   string      `json:"artwork_url"`
	StreamUrl    string      `json:"stream_url"`
	DownloadUrl  interface{} `json:"download_url"`
	WaveformUrl  string      `json:"waveform_url"`
	Downloadable bool        `json:"downloadable"`
	Access       string      `json:"access"`
	Policy       interface{} `json:"policy"`
}

type TracksList struct {
	Tracks []Track `json:"collection"`
}

func (receiver TracksList) ToJson() ([]byte, *error) {
	data, err := json.Marshal(receiver)
	return data, &err
}

func (receiver TracksList) TracksListFromJson(tracksJson string) (TracksList, *error) {
	err := json.Unmarshal([]byte(tracksJson), &receiver)
	return receiver, &err
}
