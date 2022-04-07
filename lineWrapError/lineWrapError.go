package lineWrapError

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	g "coding-standard-validator-part-2/validator/getFiles"

)
type LineWrapError struct{
	//lineNumber int
	//errors []string
	ErrorsInfo []string
}

func (l LineWrapError) HandleError(folder string) int {
	fNames := g.GetFiles(folder)
	counter := 0
	passed := true
	numE := 0
	for _, f := range fNames {
		file, err := os.Open(f)
		if err != nil {
			break
		}
		defer file.Close()
		reader := bufio.NewReader(file)
		for {
			counter ++
			line, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			if (len(line) > 100){
				errM:= errors.New("lines wrapped over 100 characters")
				fmt.Printf("error running program: %s\n\t, file: %s, line: %d \n", errM.Error(), f, counter)
				str1 := strconv.Itoa(counter)
				fNameInfo := "File name: " +f
				l.ErrorsInfo = append(l.ErrorsInfo,fNameInfo , " line: " , str1)
				numE ++
			}
		}
		
	}
	if passed {
		fmt.Println("Passed proper line wrapping standard")
	}
	return numE
}
