package hookcore

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRandomString(t *testing.T) {
	length := 10
	randomStr := RandomString(length)
	if len(randomStr) != length {
		t.Errorf("Random string length is not equal to %d", length)
	}
}

func TestRandomURL(t *testing.T) {
	randomURL := RandomURL()
	re := regexp.MustCompile(`^https?://[a-zA-Z0-9]{5,15}\.com$`)
	if !re.MatchString(randomURL) {
		t.Errorf("Generated URL format is incorrect: %s", randomURL)
	}
}

func TestRandomizeJSON(t *testing.T) {
	data := map[string]interface{}{
		"key1": "https://example.com",
		"key2": "randomString",
		"key3": []interface{}{
			"https://example.com",
			"randomString",
		},
	}

	RandomizeJSON(data)
	fmt.Print(data)

	for _, value := range data {
		switch v := value.(type) {
		case string:
			if IsURL(v) {
				if len(v) < 10 || len(v) > 20 {
					t.Errorf("URL length is out of range: %s", v)
				}
			} else {
				if len(v) != 8 {
					t.Errorf("Randomized string length is not 8: %s", v)
				}
			}
		case []interface{}:
			for _, innerValue := range v {
				switch iv := innerValue.(type) {
				case string:
					if IsURL(iv) {
						if len(iv) < 10 || len(iv) > 20 {
							t.Errorf("URL length is out of range: %s", iv)
						}
					} else {
						if len(iv) != 8 {
							t.Errorf("Randomized string length is not 8: %s", iv)
						}
					}
				}
			}
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
		if IsURL(url) != expected {
			t.Errorf("isURL returned incorrect result for %s", url)
		}
	}
}
