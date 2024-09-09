package interfaces

import (
	"GrazerCodingChallenge/structs"
)

type UserRepository interface {
	InsertMockUserData()
	InsertMockUserPreferenceData()
	GetUserPreferences(id int32) (*structs.Preferences, error)
	GetUserById(id int32) (*structs.User, error)
	GetUsers() ([]structs.User, error)
	GetMatches(id int32) ([]structs.User, error)
	LikeUser(userId int32, likedId int32) (bool, error)
	GetLikes(userId int32) ([]structs.Like, error)
	FindLike(baseUserId int32, likeIdToFind int32) *structs.Like
	CreateMatch(userId int32, matchedUserID int32) (bool, error)
	FilterUsers(userId int32) ([]structs.User, *structs.User, error)
}
