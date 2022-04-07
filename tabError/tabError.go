package tabError

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	g "coding-standard-validator-part-2/validator/getFiles"

)
type TabError struct{
	ErrorsInfo []string 
	
}
 
func (t TabError) HandleError(folder string) int {
	fNames := g.GetFiles(folder)
	passed := true
	counter := 1
	numE := 0
	for _, f := range fNames {
		file, err := os.Open(f)
		if err != nil {
			break
		}
		defer file.Close()
		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			if (strings.HasPrefix(line, "          ")){
				passed = false
				errM:= errors.New("tab: empty spaces ")
				fmt.Printf("error running program: %s\n\t, file: %s, line: %d \n", errM.Error(), f, counter)
				strLine := strconv.Itoa(counter)
				fNameInfo := "File name: " + f
				t.ErrorsInfo = append(t.ErrorsInfo,strLine,fNameInfo)
				numE ++
			}
			counter ++
		}
		
	}
	if (passed){
		fmt.Println("Passed proper tab standard")
	}
	return numE	
}


