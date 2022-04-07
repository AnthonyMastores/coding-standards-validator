package readMeError

import (
	"errors"
	"fmt"
	"strings"
	g "coding-standard-validator-part-2/validator/getFiles"
)
type ReadMeError struct{
	//errorLine []string
	//errorsInfo []string
}
func (r ReadMeError) HandleError(folder string) int   {
	fNames := g.GetFiles(folder)
	passed := false
	var err error
	numE := 0
	for _, f := range fNames {
		if (strings.HasSuffix(f,"README.md")){ //checks if readme file in directory
			//check inside file
			passed = true
		}
	}
	if (passed){
		fmt.Printf("Passed readme file standard\n")
	}else{
		err = errors.New("no readMe file")
		fmt.Printf("error running program: %s \n", err.Error())
		numE = 1

	}
	return numE
}

func GetFiles(folder string) {
	panic("unimplemented")
}