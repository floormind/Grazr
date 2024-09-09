package api

import (
	"GrazerCodingChallenge/db"
	"GrazerCodingChallenge/helper"
	"GrazerCodingChallenge/interfaces"
	"encoding/json"
	"log"
	"net/http"
)

var mockdb = db.MockUserDb{}
var matchEngine = helper.MatchEngine{}

// /users
func GetUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	users, err := interfaces.UserRepository.GetUsers(mockdb)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("Error getting users from database", err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// /user?id={user-id}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.URL.Query().Get("id")
	result, err := helper.IntParser(idStr)

	user, err := interfaces.UserRepository.GetUserById(mockdb, result)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("Error getting users from database", err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// /user/matches?id={user-id}
func GetUserMatches(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.URL.Query().Get("id")

	result, err := helper.IntParser(idStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("Error getting users from database", err)
	}

	matches, err := interfaces.UserRepository.GetMatches(mockdb, result)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(matches)
}

// /user/like?userId={user-id}&likedId={liked-id}
func LikeUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	userIdStr := query.Get("userid")
	likedIdStr := query.Get("likedid")

	userId, err1 := helper.IntParser(userIdStr)
	likedId, err2 := helper.IntParser(likedIdStr)

	if err1 != nil || err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatalf("Error getting users from database", err1, err2)
	}

	created, err := interfaces.UserRepository.LikeUser(mockdb, userId, likedId)

	if err != nil || !created {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Like created successfully"})
}

// /user/likes?id={userid}
func GetLikes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	userId, err := helper.IntParser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	likes, err := interfaces.UserRepository.GetLikes(mockdb, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(likes)
}

// /user/search?id={user-id}
// search for users based on preferences.
func Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userIdStr := r.URL.Query().Get("id")
	id, err := helper.IntParser(userIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatalf("Enter a valid id", err)
	}

	// This is not an efficient way to do this as it is making a call to the database more than once for this request,
	// it is also retrieving all users to then perform a filtering on the result
	// the best way is to do the filtering at the database level, and then return the result
	// this way the searching takes a shorter amount of time.
	// but for the sake of this test, I am doing it this way.
	filteredUsers, mainUser, err := interfaces.UserRepository.FilterUsers(mockdb, id)

	for _, user := range filteredUsers {
		prefs, _ := interfaces.UserRepository.GetUserPreferences(mockdb, user.ID)
		user.Preferences.Id = prefs.Id
		user.Preferences.UserId = prefs.UserId
		user.Preferences.Gender = prefs.Gender
		user.Preferences.AgeMin = prefs.AgeMin
		user.Preferences.AgeMax = prefs.AgeMax
		user.Preferences.DietType = prefs.DietType
		user.Preferences.Distance = prefs.Distance
	}

	usersMatchingPreferences := interfaces.MatchingEngineInterface.Search(matchEngine, *mainUser, filteredUsers)

	for _, userMatchingPreferences := range usersMatchingPreferences {
		pref, err := interfaces.UserRepository.GetUserPreferences(mockdb, id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatalf("Error getting users from database", err)
		}
		userMatchingPreferences.Preferences.UserId = pref.UserId
		userMatchingPreferences.Preferences.Gender = pref.Gender
		userMatchingPreferences.Preferences.DietType = pref.DietType
		userMatchingPreferences.Preferences.AgeMin = pref.AgeMin
		userMatchingPreferences.Preferences.AgeMax = pref.AgeMax
		userMatchingPreferences.Preferences.Distance = pref.Distance
	}

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Fatalf("Error getting users from database", err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usersMatchingPreferences)
}

// Preferences
// /user/preference?id={user-id}
func GetUserPreferences(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userIdStr := r.URL.Query().Get("id")
	id, err := helper.IntParser(userIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatalf("Enter a valid id", err)
	}

	pref, err := interfaces.UserRepository.GetUserPreferences(mockdb, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Fatalf("Error getting users from database", err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pref)
}
