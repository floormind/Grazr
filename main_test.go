package main

import (
	"GrazerCodingChallenge/interfaces"
	"database/sql"
	"testing"
)

// Test function for inserting a user
func TestInsertUser(t *testing.T) {

	mockSqlDb := &sql.DB{}
	mockDB := &interfaces.UserRepository.InsertMockUserData(mockSqlDb)

	// Test InsertUser with mock database
	mockDB.InsertMockUserData(mockSqlDb)
}
