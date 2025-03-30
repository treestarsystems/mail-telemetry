package utils

import (
	"math/rand"
	"reflect"
	"regexp"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Removes all HTML tags from the given string.
func RemoveHTMLTags(text string) string {
	// Define a regex pattern to match HTML tags
	re := regexp.MustCompile(`<[^>]*>`)
	// Replace all HTML tags with an empty string
	cleanText := re.ReplaceAllString(text, "")
	return cleanText
}

// Deduplicates a slice of any type and converts all strings to lowercase.
func DeduplicateSliceContents(slice interface{}) interface{} {
	// Use reflection to get the value and type of the input slice
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return []string{}
	}

	// Create a map to track unique elements
	uniqueMap := make(map[interface{}]bool)
	uniqueSlice := reflect.MakeSlice(v.Type(), 0, v.Len())

	// Iterate over the input slice and add unique elements to the result slice
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i).Interface()

		// Convert strings to lowercase
		if str, ok := elem.(string); ok {
			elem = strings.ToLower(str)
		}

		if !uniqueMap[elem] {
			uniqueMap[elem] = true
			uniqueSlice = reflect.Append(uniqueSlice, reflect.ValueOf(elem))
		}
	}

	return uniqueSlice.Interface()
}

// RandomString generates a random string of the specified length.
func RandomAplhaNumericString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
