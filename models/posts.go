package models

import "time"

type Post struct {
	UserID    int
	PostID    int
	Content   string
	Timestamp time.Time
}

var Posts = []Post{
	{UserID: 2, PostID: 201, Content: "Hello from user-2", Timestamp: time.Now().Add(-1 * time.Hour)},
	{UserID: 3, PostID: 301, Content: "Hello from user-3", Timestamp: time.Now().Add(-2 * time.Hour)},
	{UserID: 2, PostID: 202, Content: "Another post from user-2", Timestamp: time.Now().Add(-3 * time.Hour)},
	{UserID: 4, PostID: 401, Content: "User 4 here!", Timestamp: time.Now().Add(-30 * time.Minute)},
	{UserID: 5, PostID: 501, Content: "Nina’s thoughts today", Timestamp: time.Now().Add(-45 * time.Minute)},
}
