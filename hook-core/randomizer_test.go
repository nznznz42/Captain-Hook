/*
Copyright © 2024 nznznz42
*/
package hookcore_test

import (
	hookcore "hooktest/hook-core"
	"regexp"
	"testing"
)

func TestRandomString(t *testing.T) {
	length := 10
	randomStr := hookcore.RandomString(length)
	if len(randomStr) != length {
		t.Errorf("Random string length is not equal to %d", length)
	}
}

func TestRandomURL(t *testing.T) {
	randomURL := hookcore.RandomURL()
	re := regexp.MustCompile(`^https?://[a-zA-Z0-9]{1,15}\.com$`)
	if !re.MatchString(randomURL) {
		t.Errorf("Generated URL format is incorrect: %s", randomURL)
	}
}

func TestRandomizeJSON(t *testing.T) {
	originalData := map[string]interface{}{
		"key1": "https://example.com",
		"key2": "randomString",
		"key3": []interface{}{
			"https://example.com",
			"randomString",
		},
	}

	data := make(map[string]interface{})
	for key, value := range originalData {
		data[key] = value
	}

	hookcore.RandomizeJSON(data)

	for key, value := range data {
		originalValue, ok := originalData[key]
		if !ok {
			t.Errorf("Key %s does not exist in original data", key)
			continue
		}

		switch v := value.(type) {
		case string:
			if originalValue == value {
				t.Errorf("Field %s is not randomized: %s", key, v)
			}
		case []interface{}:
			originalSlice, ok := originalValue.([]interface{})
			if !ok {
				t.Errorf("Field %s in original data is not a slice", key)
				continue
			}

			if len(v) != len(originalSlice) {
				t.Errorf("Length of slice %s is different from original data", key)
				continue
			}
		default:
			t.Errorf("Unexpected data type for field %s", key)
		}
	}
}

func TestIsURL(t *testing.T) {
	testCases := map[string]bool{
		"http://example.com":  true,
		"https://example.com": true,
		"ftp://example.com":   true,
		"example.com":         false,
		"random string":       false,
	}

	for url, expected := range testCases {
		if hookcore.IsURL(url) != expected {
			t.Errorf("isURL returned incorrect result for %s", url)
		}
	}
}
