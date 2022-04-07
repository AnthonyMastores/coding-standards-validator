package main

import (
	l2 "coding-standard-validator-part-2/validator/licenseError"
	l1 "coding-standard-validator-part-2/validator/lineWrapError"
	r "coding-standard-validator-part-2/validator/readMeError"
	s "coding-standard-validator-part-2/validator/standards"
	t "coding-standard-validator-part-2/validator/tabError"
	"fmt"
	"os"
)

func main() {
	var (
		//root string = "C:/Users/Anthony/go/APL/src/360-anthony-mastores/sprint-2"
		folder string
		//errors []error
	)
	if (len(os.Args) > 1) && os.Args[1] == "help" { //val help
		q := "This program validates a go project based on four coding standards: " +
			"Contains a readme file, Contains a license file, " +
			"follows proper tab standards, and wraps text at most to 100 characters."
		fmt.Println(q)
		//os.Exit(0)
	}

	var rE r.ReadMeError
	var tE t.TabError
	var l1E l1.LineWrapError
	var l2E l2.LicenseError
	var numER, numET, numEL1, numEL2 int

	numER = ValidateStandard(rE, folder)
	numET = ValidateStandard(tE, folder)
	numEL1 = ValidateStandard(l1E, folder)
	numEL2 = ValidateStandard(l2E, folder)

	if len(os.Args) > 1 && os.Args[1] == "detail" { //val detail: file name, line number error
		fmt.Printf("Number of read me errors: %d\nNumber of tab errors: %d\n", numER, numET)
		fmt.Printf("Number of line wrap errors: %d\nNumber of license errors: %d", numEL1, numEL2)
		//fmt.Println(tE.ErrorsInfo)
		//fmt.Println(l1E.ErrorsInfo)
	}

}

/*func GetFiles( folder string) ([]string) {
	var files []string
    err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
        files = append(files, path)
        return nil
    })
    if err != nil {
        log.Fatal(err)
    }
	return files

    }
*/
func ValidateStandard(s s.Standards, folder string) int {
	numE := s.HandleError(folder)
	return numE
}
