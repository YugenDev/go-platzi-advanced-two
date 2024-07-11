package models

import "time"

type Posts struct {
	Id          string    `json:"id"`
	PostContent string    `json:"post_content"`
	CreatedAt    time.Time `json:"create_at"`
	UserId      string    `json:"user_id"`
}
