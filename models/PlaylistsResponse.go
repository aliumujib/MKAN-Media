package models

type PlaylistsResponse struct {
	Playlists []Playlist `json:"collection"`
	NextUrl   string     `json:"next_href"`
}

type Playlist struct {
	Duration     int         `json:"duration"`
	Genre        string      `json:"genre"`
	ReleaseDay   *int        `json:"release_day"`
	Permalink    string      `json:"permalink"`
	PermalinkUrl string      `json:"permalink_url"`
	ReleaseMonth *int        `json:"release_month"`
	ReleaseYear  *int        `json:"release_year"`
	Description  *string     `json:"description"`
	Uri          string      `json:"uri"`
	LabelName    *string     `json:"label_name"`
	LabelId      interface{} `json:"label_id"`
	Label        interface{} `json:"label"`
	TagList      string      `json:"tag_list"`
	TrackCount   int         `json:"track_count"`
	UserId       int         `json:"user_id"`
	LastModified string      `json:"last_modified"`
	License      string      `json:"license"`
	User         struct {
		AvatarUrl            string      `json:"avatar_url"`
		Id                   int         `json:"id"`
		Kind                 string      `json:"kind"`
		PermalinkUrl         string      `json:"permalink_url"`
		Uri                  string      `json:"uri"`
		Username             string      `json:"username"`
		Permalink            string      `json:"permalink"`
		CreatedAt            string      `json:"created_at"`
		LastModified         string      `json:"last_modified"`
		FirstName            string      `json:"first_name"`
		LastName             string      `json:"last_name"`
		FullName             string      `json:"full_name"`
		City                 string      `json:"city"`
		Description          string      `json:"description"`
		Country              string      `json:"country"`
		TrackCount           int         `json:"track_count"`
		PublicFavoritesCount int         `json:"public_favorites_count"`
		RepostsCount         int         `json:"reposts_count"`
		FollowersCount       int         `json:"followers_count"`
		FollowingsCount      int         `json:"followings_count"`
		Plan                 string      `json:"plan"`
		MyspaceName          interface{} `json:"myspace_name"`
		DiscogsName          interface{} `json:"discogs_name"`
		WebsiteTitle         string      `json:"website_title"`
		Website              string      `json:"website"`
		CommentsCount        int         `json:"comments_count"`
		Online               bool        `json:"online"`
		LikesCount           int         `json:"likes_count"`
		PlaylistCount        int         `json:"playlist_count"`
		Subscriptions        []struct {
			Product struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"product"`
		} `json:"subscriptions"`
	} `json:"user"`
	PlaylistType  string      `json:"playlist_type"`
	Type          string      `json:"type"`
	Id            int         `json:"id"`
	Downloadable  interface{} `json:"downloadable"`
	LikesCount    int         `json:"likes_count"`
	Sharing       string      `json:"sharing"`
	CreatedAt     string      `json:"created_at"`
	Release       interface{} `json:"release"`
	Tags          string      `json:"tags"`
	Kind          string      `json:"kind"`
	Title         string      `json:"title"`
	PurchaseTitle interface{} `json:"purchase_title"`
	Ean           interface{} `json:"ean"`
	Streamable    bool        `json:"streamable"`
	EmbeddableBy  string      `json:"embeddable_by"`
	ArtworkUrl    *string     `json:"artwork_url"`
	PurchaseUrl   interface{} `json:"purchase_url"`
	TracksUri     string      `json:"tracks_uri"`
	Tracks        []struct {
		Id int `json:"id"`
	} `json:"tracks"`
}
