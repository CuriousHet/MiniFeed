package main

import (
	"fmt"
	"sort"

	"github.com/CuriousHet/MiniFeed/models"
)

func main() {
	// Simulate fetching posts for User 1
	userID := 1
	posts := getTimeline(userID)

	// Print out posts in reverse chronological order
	for _, post := range posts {
		fmt.Printf("Post by User %d: %s\n", post.UserID, post.Content)
	}
}

// getTimeline fetches the posts for the users that the given user follows.
func getTimeline(userID int) []models.Post {
	// Simulate a user following another user (hardcoded relationship)
	followedUsers := getFollowedUsers(userID)

	// Filter posts for followed users
	var timeline []models.Post
	for _, post := range models.Posts {
		for _, user := range followedUsers {
			if post.UserID == user.UserID {
				timeline = append(timeline, post)
			}
		}
	}

	// Sort posts by timestamp in reverse chronological order
	sort.Slice(timeline, func(i, j int) bool {
		return timeline[i].Timestamp.After(timeline[j].Timestamp)
	})

	// Return the latest 20 posts (or fewer if there aren't that many)
	if len(timeline) > 20 {
		timeline = timeline[:20]
	}
	return timeline
}

// getFollowedUsers returns the users that the given user is following.
func getFollowedUsers(userID int) []models.User {
	// For simplicity, User 1 follows User 2 (hardcoded)
	if userID == 1 {
		return []models.User{
			models.Users[1], // User 2
		}
	}

	// If a user is not User 1, they don't follow anyone
	return []models.User{}
}
