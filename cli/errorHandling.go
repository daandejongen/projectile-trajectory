package cli

import (
	"fmt"
	"os"
)

type errorAction = func(error)

func printErrorToStOut(err error) {
	fmt.Println(err)
}

func exitWithErrorStatus(error) {
	os.Exit(1)
}

func handleErrors(action errorAction, possibleErrors ...error) {
	hasError := false
	errors := []error{}

	for _, possibleErr := range possibleErrors {
		if possibleErr != nil {
			hasError = true
		}
		errors = append(errors, possibleErr)
	}

	if hasError {
		for _, err := range errors {
			action(err)
		}
	}
}