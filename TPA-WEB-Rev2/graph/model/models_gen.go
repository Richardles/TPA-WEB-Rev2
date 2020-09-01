// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Activity struct {
	ID      int    `json:"id"`
	UserID  string `json:"user_id"`
	VideoID *int   `json:"video_id"`
	PostID  *int   `json:"post_id"`
}

type Comment struct {
	ID        int    `json:"id"`
	Comment   string `json:"comment"`
	UserID    string `json:"user_id"`
	Date      string `json:"date"`
	Likes     int    `json:"likes"`
	Dislikes  int    `json:"dislikes"`
	RepliesID *int   `json:"replies_id"`
	VideoID   *int   `json:"video_id"`
}

type Playlist struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	TotalVideos *int   `json:"total_videos"`
	Views       *int   `json:"views"`
	LastUpdated string `json:"last_updated"`
	ViewType    string `json:"view_type"`
	Description string `json:"description"`
	UserID      string `json:"userId"`
	VideosID    string `json:"videos_id"`
}

type Post struct {
	ID          int     `json:"id"`
	PictureURL  *string `json:"picture_url"`
	Likes       int     `json:"likes"`
	Dislikes    int     `json:"dislikes"`
	Date        string  `json:"date"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ChannelID   string  `json:"channel_id"`
}

type Secuser struct {
	ID                 string  `json:"id"`
	Name               string  `json:"name"`
	ProfilePicture     string  `json:"profile_picture"`
	Subscriber         int     `json:"subscriber"`
	Email              string  `json:"email"`
	Location           string  `json:"location"`
	Premium            bool    `json:"premium"`
	Restriction        bool    `json:"restriction"`
	PremiumDate        string  `json:"premium_date"`
	ChannelIcon        string  `json:"channel_icon"`
	ChannelDescription string  `json:"channel_description"`
	ChannelJoinDate    string  `json:"channel_join_date"`
	ChannelViews       *int    `json:"channel_views"`
	ChannelLocation    string  `json:"channel_location"`
	ChannelArt         *string `json:"channel_art"`
	LikeComment        string  `json:"like_comment"`
	DislikeComment     string  `json:"dislike_comment"`
	Subscribed         string  `json:"subscribed"`
	NotifiedBy         string  `json:"notified_by"`
	LikeVideo          *string `json:"like_video"`
	DislikeVideo       *string `json:"dislike_video"`
	LikePost           *string `json:"like_post"`
	DislikePost        *string `json:"dislike_post"`
}

type Secvid struct {
	ID          int     `json:"id"`
	URL         string  `json:"url"`
	Title       string  `json:"title"`
	Likes       *int    `json:"likes"`
	Dislikes    *int    `json:"dislikes"`
	Description *string `json:"description"`
	Thumbnail   *string `json:"thumbnail"`
	UserID      *string `json:"userId"`
	Views       *int    `json:"views"`
	PlaylistID  int     `json:"playlist_id"`
	Category    *string `json:"category"`
	Audience    *string `json:"audience"`
	Visibility  *string `json:"visibility"`
	Premium     *string `json:"premium"`
	Date        string  `json:"date"`
}

type NewActivity struct {
	UserID  string `json:"user_id"`
	VideoID *int   `json:"video_id"`
	PostID  *int   `json:"post_id"`
}

type NewComment struct {
	Comment   string `json:"comment"`
	UserID    string `json:"user_id"`
	Likes     int    `json:"likes"`
	Dislikes  int    `json:"dislikes"`
	RepliesID *int   `json:"replies_id"`
	VideoID   *int   `json:"video_id"`
}

type NewPlaylist struct {
	Title       string  `json:"title"`
	TotalVideos *int    `json:"total_videos"`
	Views       *int    `json:"views"`
	ViewType    string  `json:"view_type"`
	Description *string `json:"description"`
	UserID      string  `json:"userId"`
	VideosID    string  `json:"videos_id"`
}

type NewPost struct {
	PictureURL  *string `json:"picture_url"`
	Likes       int     `json:"likes"`
	Dislikes    int     `json:"dislikes"`
	Date        string  `json:"date"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ChannelID   string  `json:"channel_id"`
}

type NewUser struct {
	ID                 string  `json:"id"`
	Name               string  `json:"name"`
	ProfilePicture     string  `json:"profile_picture"`
	Subscriber         int     `json:"subscriber"`
	Email              string  `json:"email"`
	Location           string  `json:"location"`
	Premium            bool    `json:"premium"`
	Restriction        bool    `json:"restriction"`
	ChannelIcon        string  `json:"channel_icon"`
	ChannelDescription string  `json:"channel_description"`
	ChannelJoinDate    string  `json:"channel_join_date"`
	ChannelViews       *int    `json:"channel_views"`
	ChannelLocation    string  `json:"channel_location"`
	ChannelArt         *string `json:"channel_art"`
	LikeComment        string  `json:"like_comment"`
	DislikeComment     string  `json:"dislike_comment"`
	Subscribed         string  `json:"subscribed"`
	LikeVideo          *string `json:"like_video"`
	DislikeVideo       *string `json:"dislike_video"`
	LikePost           *string `json:"like_post"`
	DislikePost        *string `json:"dislike_post"`
}

type NewVideo struct {
	URL         string  `json:"url"`
	Title       string  `json:"title"`
	Likes       *int    `json:"likes"`
	Dislikes    *int    `json:"dislikes"`
	Description *string `json:"description"`
	Thumbnail   *string `json:"thumbnail"`
	UserID      *string `json:"userId"`
	Views       *int    `json:"views"`
	PlaylistID  int     `json:"playlist_id"`
	Category    *string `json:"category"`
	Audience    *string `json:"audience"`
	Visibility  *string `json:"visibility"`
	Premium     *string `json:"premium"`
	Date        string  `json:"date"`
}
