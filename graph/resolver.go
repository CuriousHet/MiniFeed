package graph

import (
	"context"
	"sort"
	"strconv"
	"time"

	"github.com/CuriousHet/MiniFeed/models"
)

type Resolver struct{}

type postResolver struct {
	p models.Post
}

func (r *Resolver) GetTimeline(ctx context.Context, args struct{ UserID string }) ([]*postResolver, error) {
	uid, err := strconv.Atoi(args.UserID)
	if err != nil {
		return nil, err
	}

	posts := getTimeline(uid)
	var result []*postResolver
	for _, p := range posts {
		result = append(result, &postResolver{p: p})
	}
	return result, nil
}

func (r *postResolver) UserID() int32 {
	return int32(r.p.UserID)
}

func (r *postResolver) PostID() int32 {
	return int32(r.p.PostID)
}

func (r *postResolver) Content() string {
	return r.p.Content
}

func (r *postResolver) Timestamp() string {
	return r.p.Timestamp.Format(time.RFC3339)
}

func getTimeline(userID int) []models.Post {
	followedUsers := getFollowedUsers(userID)

	var timeline []models.Post
	for _, post := range models.Posts {
		for _, user := range followedUsers {
			if post.UserID == user.UserID {
				timeline = append(timeline, post)
			}
		}
	}

	sort.Slice(timeline, func(i, j int) bool {
		return timeline[i].Timestamp.After(timeline[j].Timestamp)
	})

	if len(timeline) > 20 {
		timeline = timeline[:20]
	}
	return timeline
}

func getFollowedUsers(userID int) []models.User {
	if userID == 1 {
		return []models.User{
			models.Users[1],
		}
	}
	return []models.User{}
}
