package helper

import (
	"GrazerCodingChallenge/structs"
	"math"
)

// var users []structs.User
var likes []structs.Like
var matches []structs.Match

type MatchEngine struct{}

func (m MatchEngine) Search(user structs.User, users []structs.User) []structs.User {

	var usersMatchingPreferences []structs.User
	for _, other := range users {
		if other.ID == user.ID {
			continue
		}

		userLocation := structs.Location{
			Lat: user.Location_Lat,
			Lng: user.Location_Lng,
		}

		otherLocation := structs.Location{
			Lat: other.Location_Lat,
			Lng: other.Location_Lng,
		}

		ageRange := structs.AgeRange{
			Min: user.Preferences.AgeMin,
			Max: user.Preferences.AgeMax,
		}

		if other.Gender == user.Preferences.Gender &&
			other.DietType == user.Preferences.DietType &&
			ageInRange(other.Age, ageRange) &&
			m.calculateDistance(userLocation, otherLocation) <= user.Preferences.Distance {
			usersMatchingPreferences = append(usersMatchingPreferences, other)
		}
	}
	return usersMatchingPreferences
}

// This could be done in a much better way
// however that could take a long time.
// but the idea  is to use a geospacial algorithm that divides the earth into smaller blocks
// with each block having it's own identifies and blocks nearer to it will be considered close.
// can explain this better when we next chat.
func (m MatchEngine) calculateDistance(userLocation structs.Location, otherLocation structs.Location) float64 {
	const earthRadius = 3958.8

	latDiff := (otherLocation.Lat - userLocation.Lat) * (math.Pi / 180.0)
	lngDiff := (otherLocation.Lng - userLocation.Lng) * (math.Pi / 180.0)

	a := math.Sin(latDiff/2)*math.Sin(latDiff/2) +
		math.Cos(userLocation.Lat*(math.Pi/180.0))*math.Cos(otherLocation.Lat*(math.Pi/180.0))*
			math.Sin(lngDiff/2)*math.Sin(lngDiff/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return earthRadius * c
}

// Helper function to check if age is within the range
func ageInRange(age int32, ageRange structs.AgeRange) bool {
	return age >= ageRange.Min && age <= ageRange.Max
}
