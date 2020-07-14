package models

import uuid "github.com/satori/go.uuid"

func GenerateID() string {
	return uuid.NewV4().String()
}

func stringInArray(str string, list []string) bool {
	for _, _str := range list {
		if _str == str {
			return true
		}
	}
	return false
}
