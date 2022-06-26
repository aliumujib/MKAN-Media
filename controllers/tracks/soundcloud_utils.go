package tracks

import "time"

var baseSoundcloudAPIUrl = "https://api.soundcloud.com"

var MkaUserId = "204709701"

var SoundCloudTokenKey = "SoundCloudTokenKey"
var RecommendedTracksKey = "RecommendedTracksKey"

// urls
var authEndPoint = baseSoundcloudAPIUrl + "/oauth2/token"
var params = "?linked_partitioning=true&page_size=200"
var TrackStartPoint = baseSoundcloudAPIUrl + "/users/" + MkaUserId + "/tracks" + params
var PlaylistStartPoint = baseSoundcloudAPIUrl + "/users/" + MkaUserId + "/playlists" + params

// timeouts
var TokenCacheTimeOut = 1 * time.Hour
var RecommenderCacheTimeOut = 24 * time.Hour
