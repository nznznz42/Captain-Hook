/*
Copyright © 2024 nznznz42
*/
package hookcore

import (
	"math/rand"
	"reflect"
	"regexp"
	"time"
)

func RandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomURL() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := seededRand.Intn(10) + 5
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return "https://" + string(b) + ".com"
}

func RandomizeJSON(data interface{}) {
	switch val := data.(type) {
	case map[string]interface{}:
		for key, value := range val {
			RandomizeJSON(value)
			if reflect.TypeOf(value).Kind() == reflect.String {
				if IsURL(value.(string)) {
					val[key] = RandomURL()
				} else {
					val[key] = RandomString(8)
				}
			}
		}
	case []interface{}:
		for _, value := range val {
			RandomizeJSON(value)
		}
	}
}

func IsURL(s string) bool {
	re := regexp.MustCompile(`(?i)\b((?:https?|ftp)://\S+)\b`)
	return re.MatchString(s)
}
