package db

import (
	"GrazerCodingChallenge/structs"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type MockUserDb struct {
	DB *sql.DB
}

func (m MockUserDb) InitDB() {
	os.Remove("../sqlite-database.db")
	file, err := os.Create("../sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")

	log.Println("Opening SQLite database")

	m.createUserTable()
	m.createMatchTable()
	m.createLikeTable()
	m.createUserPreference()

	m.InsertMockUserData()
	m.InsertMockUserPreferenceData()
}

func (m MockUserDb) InsertMockUserData() {
	m.DB, _ = sql.Open("sqlite3", "../sqlite-database.db")
	users := []structs.User{
		{Name: "Alex", Gender: "male", Location_Lng: 0.0, Location_Lat: 0.0, DietType: "vegan", Age: 25},
		{Name: "Bob", Gender: "male", Location_Lng: 40.7128, Location_Lat: -74.0060, DietType: "vegetarian", Age: 30},
		{Name: "Charlie", Gender: "female", Location_Lng: 34.0522, Location_Lat: -118.2437, DietType: "omnivore", Age: 28},
		{Name: "Diana", Gender: "female", Location_Lng: 48.8566, Location_Lat: 2.3522, DietType: "vegan", Age: 26},
		{Name: "Leanne", Gender: "female", Location_Lng: 0.0, Location_Lat: 0.0, DietType: "vegan", Age: 26},
	}
	defer m.DB.Close() // Defer Closing the database
	insertUserSQL := `INSERT INTO users(name, gender, location_lat, location_lng, diet_type, age) VALUES (?, ?, ?, ?, ?, ?)`

	for _, user := range users {
		m.DB.Exec(insertUserSQL, user.Name, user.Gender, user.Location_Lat, user.Location_Lng, user.DietType, user.Age)
	}
}

func (m MockUserDb) InsertMockUserPreferenceData() {
	m.DB, _ = sql.Open("sqlite3", "../sqlite-database.db")
	preferences := []structs.Preferences{
		structs.Preferences{
			UserId:   1,
			Gender:   "female",
			DietType: "vegan",
			AgeMin:   20,
			AgeMax:   30,
			Distance: 5,
		},
		structs.Preferences{
			UserId:   5,
			Gender:   "male",
			DietType: "vegan",
			AgeMin:   20,
			AgeMax:   30,
			Distance: 5,
		},
		structs.Preferences{
			UserId:   5,
			Gender:   "female",
			DietType: "vegan",
			AgeMin:   20,
			AgeMax:   30,
			Distance: 5,
		},
	}
	defer m.DB.Close() // Defer Closing the database
	insertPreferencesSQL := `INSERT INTO user_preferences(user_id, gender, diet_type, age_min, age_max, distance) VALUES (?, ?, ?, ?, ?, ?)`

	for _, pref := range preferences {
		m.DB.Exec(insertPreferencesSQL, pref.UserId, pref.Gender, pref.DietType, pref.AgeMin, pref.AgeMax, pref.Distance)
	}

}

func (m MockUserDb) createUserTable() {
	m.DB, _ = sql.Open("sqlite3", "../sqlite-database.db")
	createUserTableSQL := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		gender TEXT NOT NULL,
		location_lat TEXT NOT NULL,
		location_lng TEXT NOT NULL,
		diet_type TEXT NOT NULL,
        age INTEGER NOT NULL
    );`

	defer m.DB.Close() // Defer Closing the database
	// Execute the SQL statement to create the table
	_, err := m.DB.Exec(createUserTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	fmt.Println("Table created successfully.")
}

func (m MockUserDb) createMatchTable() {
	m.DB, _ = sql.Open("sqlite3", "../sqlite-database.db")
	createMatchTableSQL := `CREATE TABLE IF NOT EXISTS user_matches(
    	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    	user_id INTEGER NOT NULL,
    	matched_user_id INTEGER NOT NULL,
    	created_at DATETIME NOT NULL
	);`
	defer m.DB.Close() // Defer Closing the database
	// Execute the SQL statement to create the table
	_, err := m.DB.Exec(createMatchTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	fmt.Println("Table created successfully.")
}

func (m MockUserDb) createUserPreference() {
	m.DB, _ = sql.Open("sqlite3", "../sqlite-database.db")
	createUserPreferenceSQL := `CREATE TABLE IF NOT EXISTS user_preferences(
    	id	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    	user_id INTEGER NOT NULL,
    	gender   TEXT NOT NULL,
		diet_type TEXT NOT NULL,
		age_min            TEXT NOT NULL,
		age_max				TEXT NOT NULL,
		distance           INTEGER NOT NULL
	);`
	defer m.DB.Close() // Defer Closing the database
	_, err := m.DB.Exec(createUserPreferenceSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	fmt.Println("Table created successfully.")
}

func (m MockUserDb) createLikeTable() {
	m.DB, _ = sql.Open("sqlite3", "../sqlite-database.db")
	createLikeTableSQL := `CREATE TABLE IF NOT EXISTS user_likes(
    	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    	user_id INTEGER NOT NULL,
    	liked_user_id INTEGER NOT NULL,
    	created_at DATETIME NOT NULL
	);`
	defer m.DB.Close() // Defer Closing the database
	// Execute the SQL statement to create the table
	_, err := m.DB.Exec(createLikeTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	fmt.Println("Table created successfully.")
}

// Get user preferences
func (m MockUserDb) GetUserPreferences(id int32) (*structs.Preferences, error) {
	m.DB, _ = openDb()
	defer m.DB.Close()
	query := `SELECT id, user_id, gender, diet_type, age_min, age_max, distance FROM user_preferences WHERE user_id = ?`
	var preference structs.Preferences
	defer m.DB.Close() // Defer Closing the database
	err := m.DB.QueryRow(query, id).Scan(&preference.Id, &preference.UserId, &preference.Gender, &preference.DietType, &preference.AgeMin, &preference.AgeMax, &preference.Distance)
	if err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	}
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("Failed to create table: %v", err)
		return nil, err
	}
	return &preference, nil
}

// Get users by ID
func (m MockUserDb) GetUserById(id int32) (*structs.User, error) {
	m.DB, _ = openDb()
	defer m.DB.Close()
	query := `SELECT id, name, gender, location_lat, location_lng, diet_type, age FROM users WHERE id = ?`

	var user structs.User
	defer m.DB.Close() // Defer Closing the database
	err := m.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Gender, &user.Location_Lat, &user.Location_Lng, &user.DietType, &user.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with ID %d not found", id)
		}
		return nil, err
	}

	// Return the retrieved user
	return &user, nil
}

// Get All Users
func (m MockUserDb) GetUsers() ([]structs.User, error) {
	m.DB, _ = openDb()
	defer m.DB.Close()
	query := `SELECT id, name, gender, location_lat, location_lng, diet_type, age FROM users`
	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Slice to hold the users
	var users []structs.User
	defer m.DB.Close() // Defer Closing the database
	for rows.Next() {
		var user structs.User
		err := rows.Scan(&user.ID, &user.Name, &user.Gender, &user.Location_Lat, &user.Location_Lng, &user.DietType, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	// Check for any error that might have occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Get user matches
func (m MockUserDb) GetMatches(id int32) ([]structs.User, error) {
	var users []structs.User
	m.DB, _ = openDb()
	defer m.DB.Close()
	query := `SELECT id, userId, macthed_user_id, created_at FROM user_matches WHERE userId = ?`

	//var match structs.Match

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return nil, err
	}

	// Slice to hold the matches
	var matches []structs.Match
	defer m.DB.Close() // Defer Closing the database
	for rows.Next() {
		var match structs.Match
		err := rows.Scan(&match.ID, &match.UserID, &match.MatchedUserID, &match.CreatedAt)
		if err != nil {
			return nil, err
		}
		matches = append(matches, match)
	}

	for _, match := range matches {
		user, err := m.GetUserById(match.MatchedUserID)
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}
	return users, nil
}

// assuming that userId is user1 and likedId is user2
// we check the Like table to see if user2 has a row for user1
// if so, we have a match,ß so we enter a new row into the database Match table
func (m MockUserDb) LikeUser(userId int32, likedId int32) (bool, error) {
	m.DB, _ = openDb()
	defer m.DB.Close()
	query := `INSERT INTO user_likes(user_id, liked_user_id, created_at) VALUES (?, ?, ?)`
	_, err := m.DB.Exec(query, userId, likedId, time.Now())

	// first we insert the new like record
	if err != nil {
		log.Fatalf("Failed inserting row into table", err.Error())
	}
	defer m.DB.Close() // Defer Closing the database
	// next we check if user2 also has a record for liking user1
	// if err is not nil, meaning there is an existing like from user2 to user1
	// we watch to create a match record
	existingLike := m.FindLike(likedId, userId)
	if existingLike != nil {
		_, _ = m.CreateMatch(existingLike.UserID, existingLike.LikedUserID)
	}

	return true, nil
}

func (m MockUserDb) FindLike(baseUserId int32, likeIdToFind int32) *structs.Like {
	m.DB, _ = openDb()
	defer m.DB.Close()
	var like structs.Like
	query := `SELECT user_id, liked_user_id, created_at FROM likes WHERE user_id = ? AND liked_user_id = ?`

	err := m.DB.QueryRow(query, baseUserId, likeIdToFind).Scan(&like.ID, &like.UserID, &like.LikedUserID, &like.CreatedAt)

	if err != sql.ErrNoRows {
		return nil
	} else {
		log.Printf("Like exist from the other user")
		defer m.DB.Close()
		return &like
	}
	return nil
}

func (m MockUserDb) GetLikes(userId int32) ([]structs.Like, error) {
	m.DB, _ = openDb()
	var likes []structs.Like
	query := `SELECT id, user_id, created_at FROM likes WHERE user_id = ?`
	rows, err := m.DB.Query(query, userId)

	if err != sql.ErrNoRows {
		return nil, sql.ErrNoRows
	}
	if err != nil {
		return nil, err
	}
	defer m.DB.Close() // Defer Closing the database
	for rows.Next() {
		var like structs.Like
		likes = append(likes, like)
	}
	return likes, nil
}

func (m MockUserDb) CreateMatch(userId int32, matchedUserID int32) (bool, error) {
	m.DB, _ = openDb()

	query := `INSERT INTO matches userId, matchedUserID, createdAt VALUES (?, ?, ?)`
	_, err := m.DB.Exec(query, userId, matchedUserID, time.Now())
	defer m.DB.Close() // Defer Closing the database
	if err != nil {
		log.Fatalf("Failed inserting row into table", err)
		return false, err
	}
	return true, nil
}
func (m MockUserDb) FilterUsers(userId int32) ([]structs.User, *structs.User, error) {
	m.DB, _ = openDb()
	//user we are filtering other users for.
	mainUser, err := m.GetUserById(userId)
	//user's preference
	mainUserPref, _ := m.GetUserPreferences(userId)
	query := `SELECT id, name, gender, location_lat, location_lng, diet_type, age FROM users WHERE gender = ? AND diet_type = ? AND age >= ? AND age <= ?`
	rows, err := m.DB.Query(query, mainUserPref.Gender, mainUserPref.DietType, mainUserPref.AgeMin, mainUserPref.AgeMax)

	if err != nil {
		return nil, nil, err
	}
	// Slice to hold the matching users based on search query
	var users []structs.User
	defer m.DB.Close() // Defer Closing the database
	for rows.Next() {
		var user structs.User
		err := rows.Scan(&user.ID, &user.Name, &user.Gender, &user.Location_Lat, &user.Location_Lng, &user.DietType, &user.Age)
		if err == sql.ErrNoRows {
			continue
		}
		users = append(users, user)
	}

	return users, mainUser, nil
}

func openDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "../sqlite-database.db")
	if err != nil {
		log.Fatalf("Failed to open SQLite database: %v", err)
	}
	return db, err
}
