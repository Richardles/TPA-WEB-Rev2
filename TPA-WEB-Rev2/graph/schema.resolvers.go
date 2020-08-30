package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Go_Go_Go/graph/generated"
	"Go_Go_Go/graph/model"
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.Secuser, error) {
	user := model.Secuser{
		ID:                 input.ID,
		Name:               input.Name,
		ProfilePicture:     input.ProfilePicture,
		Subscriber:         input.Subscriber,
		Email:              input.Email,
		Location:           input.Location,
		Premium:            input.Premium,
		Restriction:        input.Restriction,
		PremiumDate:        time.Now().Format("2006-01-02 15:04:05"),
		ChannelIcon:        input.ChannelIcon,
		ChannelDescription: input.ChannelDescription,
		ChannelJoinDate:    time.Now().Format("2006-01-02 15:04:05"),
		ChannelViews:       input.ChannelViews,
		ChannelLocation:    input.ChannelLocation,
		ChannelArt:         input.ChannelArt,
	}

	_, err := r.DB.Model(&user).Insert()

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input *model.NewUser) (*model.Secuser, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateSubscriber(ctx context.Context, userID string, targetID string) (*model.Secuser, error) {
	var user model.Secuser
	var target model.Secuser

	err := r.DB.Model(&user).Where("id = ?", userID).First()
	if err != nil {
		return nil, errors.New("failed to retrieve channel")
	}

	cerr := r.DB.Model(&target).Where("id = ?", targetID).First()
	if cerr != nil {
		return nil, errors.New("failed to retrieve channel")
	}

	arr := strings.Split(user.Subscribed, ",")

	if arr[0] == "" {
		user.Subscribed = targetID
		target.Subscriber = target.Subscriber + 1
	} else {
		var flag = false

		for idx, i := range arr {
			if i == targetID {
				if len(arr) == 1 {
					arr = nil
				} else {
					arr = append(arr[:idx], arr[idx+1:]...)
				}
				flag = true
				target.Subscriber = target.Subscriber - 1
				break
			}
		}

		if flag == false {
			arr = append(arr, targetID)
			target.Subscriber = target.Subscriber + 1
		}
		res := strings.Join(arr, ",")

		user.Subscribed = res
	}
	_, updateErr := r.DB.Model(&user).Where("id = ?", userID).Update()
	if updateErr != nil {
		return nil, errors.New("failed to update view")
	}

	_, updateCom := r.DB.Model(&target).Where("id = ?", targetID).Update()
	if updateCom != nil {
		return nil, errors.New("failed to update view")
	}
	return &user, nil
}

func (r *mutationResolver) UpdateLikeCommentInUser(ctx context.Context, userID string, commentID int) (*model.Secuser, error) {
	var user model.Secuser
	var comment model.Comment

	err := r.DB.Model(&user).Where("id = ?", userID).First()
	if err != nil {
		return nil, errors.New("failed to retrieve channel")
	}

	cerr := r.DB.Model(&comment).Where("id = ?", commentID).First()
	if cerr != nil {
		return nil, errors.New("failed to retrieve channel")
	}

	arr := strings.Split(user.LikeComment, ",")

	if arr[0] == "" {
		user.LikeComment = strconv.Itoa(commentID)
		comment.Likes = comment.Likes + 1
	} else {
		var flag = false

		for idx, i := range arr {
			if i == strconv.Itoa(commentID) {
				if len(arr) == 1 {
					arr = nil
				} else {
					arr = append(arr[:idx], arr[idx+1:]...)
				}
				flag = true
				comment.Likes = comment.Likes - 1
				break
			}
		}

		if flag == false {
			arr = append(arr, strconv.Itoa(commentID))
			comment.Likes = comment.Likes + 1
		}
		res := strings.Join(arr, ",")

		user.LikeComment = res
	}
	_, updateErr := r.DB.Model(&user).Where("id = ?", userID).Update()
	if updateErr != nil {
		return nil, errors.New("failed to update view")
	}

	_, updateCom := r.DB.Model(&comment).Where("id = ?", commentID).Update()
	if updateCom != nil {
		return nil, errors.New("failed to update view")
	}
	return &user, nil
}

func (r *mutationResolver) UpdateDislikeCommentInUser(ctx context.Context, userID string, commentID int) (*model.Secuser, error) {
	var user model.Secuser
	var comment model.Comment

	err := r.DB.Model(&user).Where("id = ?", userID).First()
	if err != nil {
		return nil, errors.New("failed to retrieve channel")
	}

	cerr := r.DB.Model(&comment).Where("id = ?", commentID).First()
	if cerr != nil {
		return nil, errors.New("failed to retrieve channel")
	}

	arr := strings.Split(user.DislikeComment, ",")

	if arr[0] == "" {
		user.DislikeComment = strconv.Itoa(commentID)
		comment.Dislikes = comment.Dislikes + 1
	} else {
		var flag = false

		for idx, i := range arr {
			if i == strconv.Itoa(commentID) {
				if len(arr) == 1 {
					arr = nil
				} else {
					arr = append(arr[:idx], arr[idx+1:]...)
				}
				flag = true
				comment.Dislikes = comment.Dislikes - 1
				break
			}
		}

		if flag == false {
			arr = append(arr, strconv.Itoa(commentID))
			comment.Dislikes = comment.Dislikes + 1
		}
		res := strings.Join(arr, ",")

		user.DislikeComment = res
	}
	_, updateErr := r.DB.Model(&user).Where("id = ?", userID).Update()
	if updateErr != nil {
		return nil, errors.New("failed to update view")
	}

	_, updateCom := r.DB.Model(&comment).Where("id = ?", commentID).Update()
	if updateCom != nil {
		return nil, errors.New("failed to update view")
	}
	return &user, nil
}

func (r *mutationResolver) UpdateChannelDescription(ctx context.Context, userID string, newDesc string) (*model.Secuser, error) {
	var user model.Secuser

	err := r.DB.Model(&user).Where("id = ?", userID).Select()
	if err != nil {
		return nil, errors.New("failed to retrieve channel")
	}
	user.ChannelDescription = newDesc
	_, updateErr := r.DB.Model(&user).Where("id = ?", userID).Update()
	if updateErr != nil {
		return nil, errors.New("failed to update view")
	}

	return &user, nil
}

func (r *mutationResolver) UpdateChannelArt(ctx context.Context, userID string, newArt string) (*model.Secuser, error) {
	var user model.Secuser

	err := r.DB.Model(&user).Where("id = ?", userID).Select()
	if err != nil {
		return nil, errors.New("failed to retrieve channel")
	}
	user.ChannelArt = &newArt
	_, updateErr := r.DB.Model(&user).Where("id = ?", userID).Update()
	if updateErr != nil {
		return nil, errors.New("failed to update view")
	}

	return &user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateVideo(ctx context.Context, input *model.NewVideo) (*model.Secvid, error) {
	video := model.Secvid{
		URL:         input.URL,
		Title:       input.Title,
		Likes:       input.Likes,
		Dislikes:    input.Dislikes,
		Description: input.Description,
		Thumbnail:   input.Thumbnail,
		UserID:      input.UserID,
		Views:       input.Views,
		PlaylistID:  input.PlaylistID,
		Category:    input.Category,
		Audience:    input.Audience,
		Visibility:  input.Visibility,
		Premium:     input.Premium,
	}
	_, err := r.DB.Model(&video).Insert()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &video, nil
}

func (r *mutationResolver) UpdateVideo(ctx context.Context, id int, input *model.NewVideo) (*model.Secvid, error) {
	var video model.Secvid

	err := r.DB.Model(&video).Where("id = ?", id).Select()

	if err != nil {
		return nil, errors.New("failed to retrieve channel")
	}

	video.URL = input.URL
	video.Title = input.Title
	video.Likes = input.Likes
	video.Dislikes = input.Dislikes
	video.Description = input.Description
	video.Thumbnail = input.Thumbnail
	video.UserID = input.UserID
	video.Views = input.Views
	video.PlaylistID = input.PlaylistID
	video.Category = input.Category
	video.Audience = input.Audience
	video.Visibility = input.Visibility
	video.Premium = input.Premium

	_, updateErr := r.DB.Model(&video).Where("id = ?", id).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("failed to update view")
	}

	return &video, nil
}

func (r *mutationResolver) UpdateVideoView(ctx context.Context, id int, view *int) (*model.Secvid, error) {
	var video model.Secvid

	err := r.DB.Model(&video).Where("id = ?", id).Select()
	if err != nil {
		return nil, errors.New("failed to retrieve channel")
	}
	video.Views = view
	_, updateErr := r.DB.Model(&video).Where("id = ?", id).Update()
	if updateErr != nil {
		return nil, errors.New("failed to update view")
	}
	return &video, nil
}

func (r *mutationResolver) DeleteVideo(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePlaylist(ctx context.Context, input *model.NewPlaylist) (*model.Playlist, error) {
	playlist := model.Playlist{
		Title:       input.Title,
		TotalVideos: input.TotalVideos,
		Views:       input.Views,
		LastUpdated: time.Now().Format("2006-01-02 15:04:05"),
		ViewType:    input.ViewType,
		Description: *input.Description,
		UserID:      input.UserID,
		VideosID:    input.VideosID,
	}

	_, err := r.DB.Model(&playlist).Insert()

	if err != nil {
		return nil, err
	}

	return &playlist, nil
}

func (r *mutationResolver) UpdatePlaylist(ctx context.Context, id int, input *model.NewPlaylist) (*model.Playlist, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePlaylistTitle(ctx context.Context, id int, newTitle string) (*model.Playlist, error) {
	var playlist model.Playlist

	err := r.DB.Model(&playlist).Where("id = ?", id).Select()
	if err != nil {
		return nil, errors.New("failed to retrieve channel")
	}
	playlist.Title = newTitle
	playlist.LastUpdated = time.Now().Format("2006-01-02 15:04:05")
	_, uperr := r.DB.Model(&playlist).Where("id = ?", id).Update()
	if uperr != nil {
		return nil, errors.New("failed to update view")
	}
	return &playlist, nil
}

func (r *mutationResolver) UpdatePlaylistDescription(ctx context.Context, id int, newDesc string) (*model.Playlist, error) {
	var playlist model.Playlist

	err := r.DB.Model(&playlist).Where("id = ?", id).Select()
	if err != nil {
		return nil, errors.New("failed to retrieve channel")
	}
	playlist.Description = newDesc
	playlist.LastUpdated = time.Now().Format("2006-01-02 15:04:05")
	_, uperr := r.DB.Model(&playlist).Where("id = ?", id).Update()
	if uperr != nil {
		return nil, errors.New("failed to update view")
	}
	return &playlist, nil
}

func (r *mutationResolver) UpdatePlaylistVisibility(ctx context.Context, id int, newVis string) (*model.Playlist, error) {
	var playlist model.Playlist

	err := r.DB.Model(&playlist).Where("id = ?", id).Select()
	if err != nil {
		return nil, errors.New("failed to retrieve channel")
	}
	playlist.ViewType = newVis
	playlist.LastUpdated = time.Now().Format("2006-01-02 15:04:05")
	_, uperr := r.DB.Model(&playlist).Where("id = ?", id).Update()
	if uperr != nil {
		return nil, errors.New("failed to update view")
	}
	return &playlist, nil
}

func (r *mutationResolver) UpdateVideoInPlaylist(ctx context.Context, playlistID int, videoID int) (*model.Playlist, error) {
	var playlist model.Playlist

	err := r.DB.Model(&playlist).Where("id = ?", playlistID).First()
	if err != nil {
		return nil, errors.New("failed to retrieve channel")
	}

	arr := strings.Split(playlist.VideosID, ",")

	if arr[0] == "" {
		playlist.VideosID = strconv.Itoa(videoID)
	} else {
		var flag = false

		for idx, i := range arr {
			if i == strconv.Itoa(videoID) {
				if len(arr) == 1 {
					arr = nil
				} else {
					arr = append(arr[:idx], arr[idx+1:]...)
				}
				flag = true
				break
			}
		}

		if flag == false {
			arr = append(arr, strconv.Itoa(videoID))
		}
		res := strings.Join(arr, ",")

		playlist.VideosID = res
	}
	playlist.LastUpdated = time.Now().Format("2006-01-02 15:04:05")
	_, updateErr := r.DB.Model(&playlist).Where("id = ?", playlistID).Update()
	if updateErr != nil {
		return nil, errors.New("failed to update view")
	}
	return &playlist, nil
}

func (r *mutationResolver) EmptyPlaylist(ctx context.Context, playlistID int) (*model.Playlist, error) {
	var playlist model.Playlist

	err := r.DB.Model(&playlist).Where("id = ?", playlistID).Select()
	if err != nil {
		return nil, errors.New("failed to retrieve channel")
	}
	playlist.VideosID = ""
	_, uperr := r.DB.Model(&playlist).Where("id = ?", playlistID).Update()
	if uperr != nil {
		return nil, errors.New("failed to update view")
	}
	return &playlist, nil
}

func (r *mutationResolver) DeletePlaylist(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateComment(ctx context.Context, input *model.NewComment) (*model.Comment, error) {
	comment := model.Comment{
		Comment:   input.Comment,
		UserID:    input.UserID,
		Date:      time.Now().Format("2006-01-02 15:04:05"),
		Likes:     input.Likes,
		Dislikes:  input.Dislikes,
		RepliesID: input.RepliesID,
		VideoID:   input.VideoID,
	}

	_, err := r.DB.Model(&comment).Insert()

	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (r *mutationResolver) UpdateComment(ctx context.Context, id int, input *model.NewComment) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteComment(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePost(ctx context.Context, input *model.NewPost) (*model.Post, error) {
	post := model.Post{
		PictureURL:  input.PictureURL,
		Likes:       input.Likes,
		Dislikes:    input.Dislikes,
		Date:        time.Now().Format("2006-01-02 15:04:05"),
		Title:       input.Title,
		Description: input.Description,
		ChannelID:   input.ChannelID,
	}

	_, err := r.DB.Model(&post).Insert()

	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id int, input *model.NewPost) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePost(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.Secuser, error) {
	var user []*model.Secuser

	err := r.DB.Model(&user).Select()

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *queryResolver) Videos(ctx context.Context) ([]*model.Secvid, error) {
	var video []*model.Secvid

	err := r.DB.Model(&video).Select()

	if err != nil {
		return nil, err
	}

	return video, nil
}

func (r *queryResolver) Playlists(ctx context.Context) ([]*model.Playlist, error) {
	var playlist []*model.Playlist

	err := r.DB.Model(&playlist).Select()

	if err != nil {
		return nil, err
	}

	return playlist, nil
}

func (r *queryResolver) Comments(ctx context.Context) ([]*model.Comment, error) {
	var comment []*model.Comment

	err := r.DB.Model(&comment).Select()

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetVideo(ctx context.Context, id int) (*model.Secvid, error) {
	var video model.Secvid

	err := r.DB.Model(&video).Where("id = ?", id).First()

	if err != nil {
		return nil, err
	}

	return &video, nil
}

func (r *queryResolver) GetLastVideo(ctx context.Context, id string) (*model.Secvid, error) {
	var video model.Secvid

	err := r.DB.Model(&video).Where("user_id = ?", id).Last()

	if err != nil {
		return nil, err
	}

	return &video, nil
}

func (r *queryResolver) GetVideosByUserID(ctx context.Context, userID string) ([]*model.Secvid, error) {
	var video []*model.Secvid

	err := r.DB.Model(&video).Where("user_id = ?", userID).Select()

	if err != nil {
		return nil, err
	}

	return video, nil
}

func (r *queryResolver) GetVideosByCategory(ctx context.Context, category string) ([]*model.Secvid, error) {
	var videos []*model.Secvid

	err := r.DB.Model(&videos).Where("category = ?", category).Select()

	if err != nil {
		return nil, err
	}

	return videos, nil
}

func (r *queryResolver) GetVideosByName(ctx context.Context, name string) ([]*model.Secvid, error) {
	var videos []*model.Secvid

	err := r.DB.Model(&videos).Where("LOWER(title) LIKE LOWER(?)", "%"+name+"%").Select()

	if err != nil {
		return nil, err
	}

	return videos, nil
}

func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.Secuser, error) {
	var user model.Secuser

	err := r.DB.Model(&user).Where("id = ?", id).First()

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *queryResolver) GetUsersByName(ctx context.Context, name string) ([]*model.Secuser, error) {
	var users []*model.Secuser

	err := r.DB.Model(&users).Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%").Select()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *queryResolver) GetPlaylist(ctx context.Context, id int) (*model.Playlist, error) {
	var playlist model.Playlist

	err := r.DB.Model(&playlist).Where("id = ?", id).First()

	if err != nil {
		return nil, err
	}

	return &playlist, nil
}

func (r *queryResolver) GetPlaylistByUser(ctx context.Context, id string) ([]*model.Playlist, error) {
	var playlist []*model.Playlist

	err := r.DB.Model(&playlist).Where("user_id = ?", id).Select()

	if err != nil {
		return nil, err
	}

	return playlist, nil
}

func (r *queryResolver) GetPlaylistsByName(ctx context.Context, name string) ([]*model.Playlist, error) {
	var playlists []*model.Playlist

	err := r.DB.Model(&playlists).Where("LOWER(title) LIKE LOWER(?)", "%"+name+"%").Select()

	if err != nil {
		return nil, err
	}

	return playlists, nil
}

func (r *queryResolver) GetComment(ctx context.Context, id int) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCommentByVideoID(ctx context.Context, vidID int) ([]*model.Comment, error) {
	var comments []*model.Comment

	err := r.DB.Model(&comments).Where("video_id = ?", vidID).Select()

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *queryResolver) GetRepliesByCommentID(ctx context.Context, commentID int) ([]*model.Comment, error) {
	var replies []*model.Comment

	err := r.DB.Model(&replies).Where("replies_id = ?", commentID).Select()

	if err != nil {
		return nil, err
	}

	return replies, nil
}

func (r *queryResolver) GetPost(ctx context.Context, id int) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPostByChannelID(ctx context.Context, channelID string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPostsByChannelID(ctx context.Context, channelID string) ([]*model.Post, error) {
	var posts []*model.Post

	err := r.DB.Model(&posts).Where("channel_id = ?", channelID).Select()

	if err != nil {
		return nil, err
	}

	return posts, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) GetVideoByCategory(ctx context.Context, category string) ([]*model.Secvid, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) UpdateVideoViews(ctx context.Context, videoID int) (*model.Secvid, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) UpdateChannelComment(ctx context.Context, userID string, newComment string) (*model.Secuser, error) {
	panic(fmt.Errorf("not implemented"))
}
