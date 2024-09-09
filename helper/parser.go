package helper

import (
	"log"
	"strconv"
)

func IntParser(data string) (int32, error) {
	parsedInt, err := strconv.ParseInt(data, 10, 32)
	result := int32(parsedInt)

	if err != nil {
		log.Fatalf("Error getting users from database", err)
		return 0, err
	}
	return result, nil
}
