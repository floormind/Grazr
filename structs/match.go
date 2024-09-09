package structs

import "time"

type Match struct {
	ID            int32
	UserID        int32
	MatchedUserID int32
	CreatedAt     time.Time
}
