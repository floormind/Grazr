package interfaces

import "GrazerCodingChallenge/structs"

type MatchingEngineInterface interface {
	Search(user structs.User, users []structs.User) []structs.User
}
