package LicenseError

import (
	"errors"
	"fmt"
	"strings"
	g "coding-standard-validator-part-2/validator/getFiles"

)

type LicenseError struct {
	//errors []string
}

func (l LicenseError) HandleError(folder string) int {
	fNames := g.GetFiles(folder)
	passed := false
	numE := 0
	for _, f := range fNames {
		if strings.HasSuffix(f, "license") { //checks if license file in directory
			passed = true
		}
	}
	if passed {
		fmt.Println("Passed license file standard")
	} else {
		err := errors.New("no license file")
		fmt.Printf("error running program: %s \n", err.Error())
		numE = 1
	}
	return numE
}