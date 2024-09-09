package main

import (
	"GrazerCodingChallenge/api"
	db2 "GrazerCodingChallenge/db"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {

	mockUserDb := db2.MockUserDb{}
	mockUserDb.InitDB()

	http.HandleFunc("/users", api.GetUsers)
	http.HandleFunc("/user", api.GetUser)
	http.HandleFunc("/user/like/", api.LikeUser)
	http.HandleFunc("/user/search", api.Search)
	http.HandleFunc("/user/matches", api.GetUserMatches)
	http.HandleFunc("/user/preferences", api.GetUserPreferences)
	http.HandleFunc("/user/likes", api.GetLikes)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8000", nil))

	/**
	//// Find matches for User 1
	//var fullUserList, err = mockUserDb.GetUsers()
	//
	//if err != nil {
	//	log.Fatalf("Failed to get users %v", err)
	//}
	//
	//log.Println("The users", fullUserList)
	//
	//matchEngine := MatchEngine{}
	//user1 := fullUserList[0]
	//matchedUsers := matchEngine.Search(user1,)
	//
	//// Display matched users
	//fmt.Println("Matched users for User 1:")
	//for _, match := range matchedUsers {
	//	fmt.Printf("- %s (ID: %d)\n", match.Name, match.ID)
	//}
	//
	//// User 1 likes User 2
	//matchEngine.LikeUser(1, 2)
	//
	//// User 2 likes User 1
	//matchEngine.LikeUser(2, 1)
	//
	//// Check current matches
	//fmt.Println("Current matches:")
	//for _, match := range matchedUsers {
	//	fmt.Printf("User %d matched with User %d\n", match.ID, match.Name)
	//}
	*/
}
