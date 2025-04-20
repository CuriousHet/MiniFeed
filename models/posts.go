package models

import "time"

type Post struct {
	UserID    int
	PostID    int
	Content   string
	Timestamp time.Time
}

var Posts = []Post{
	{
		UserID:    1,
		PostID:    101,
		Content:   "This is the first post content by user-1.",
		Timestamp: time.Now().Add(-time.Hour * 1),
	},
	{
		UserID:    2,
		PostID:    102,
		Content:   "Another post with more information by user-2",
		Timestamp: time.Now().Add(-time.Hour * 2),
	},
	{
		UserID:    2,
		PostID:    103,
		Content:   "Older post from user-2.",
		Timestamp: time.Now().Add(-time.Hour * 2),
	},
}
