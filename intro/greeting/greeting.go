package greeting

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos return a map that associate each of the named people
// with greeting message
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with message
	messages := make(map[string]string)

	// loop through the received slice of names, calling
	// the Hello function to get greeting message for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// init sets initial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met",
	}

	return formats[rand.Intn(len(formats))]
}
