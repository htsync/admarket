package models

type Profile struct {
	ID          int64    `json:"id"`
	UserID      int64    `json:"user_id"`
	DisplayName string   `json:"display_name"`
	Bio         string   `json:"bio"`
	Tags        []string `json:"tags"`
	CreatedAt   string   `json:"created_at"`
}
