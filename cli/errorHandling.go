package cli

import (
	"fmt"
	"os"
)

func handleErrors(exitOnError bool, possibleErrors ...error) {
	hasError := false
	errors := []error{}

	for _, possibleErr := range possibleErrors {
		if possibleErr != nil {
			hasError = true
			errors = append(errors, possibleErr)
		}
	}

	if hasError {
		for _, err := range errors {
			fmt.Println(err.Error())
		}
		if exitOnError {
			os.Exit(1)
		}
	}
}
