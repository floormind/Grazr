package structs

import "time"

// Structure for Like
type Like struct {
	ID          int32
	UserID      int32
	LikedUserID int32
	CreatedAt   time.Time
}
