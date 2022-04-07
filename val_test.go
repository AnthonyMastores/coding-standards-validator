package main

import (
	"testing"
	r "coding-standard-validator-part-2/validator/readMeError"
	tE "coding-standard-validator-part-2/validator/tabError"
	l1 "coding-standard-validator-part-2/validator/lineWrapError"
	l2 "coding-standard-validator-part-2/validator/licenseError"
)

func TestReadMeError(t *testing.T){
	cases := []struct {
        input          string
        expectedOutput int
    }{
        {"", 0},
        {"C:/Users/Anthony/go/APL/src/360-anthony-mastores/sprint-2", 0},
    
    }
	var r1 r.ReadMeError
    for _, c := range cases {
        if output := r1.HandleError(c.input); output != c.expectedOutput {
            t.Errorf("incorrect output for `%s`: expected `%d` but got `%d`", c.input, c.expectedOutput, output)
        }
    }
}
func TestLicenseError(t *testing.T){
	cases := []struct {
        input         string
        expectedOutput int
    }{
        {"", 0},
        {"C:/Users/Anthony/go/APL/src/360-anthony-mastores/sprint-2",0},
    }
	var L1 l1.LineWrapError
    for _, c := range cases {
        if output := L1.HandleError(c.input); output != c.expectedOutput {
            t.Errorf("incorrect output for `%s`: expected `%d` but got `%d`", c.input, c.expectedOutput, output)
        }
    }
}
func TestLineWrapError(t *testing.T){
	cases := []struct {
        input          string
        expectedOutput int
    }{
        {"", 1},
        {"C:/Users/Anthony/go/APL/src/360-anthony-mastores/sprint-2",1},
    }
	var L2 l2.LicenseError
    for _, c := range cases {
        if output := L2.HandleError(c.input); output != c.expectedOutput {
            t.Errorf("incorrect output for `%s`: expected `%d` but got `%d`", c.input, c.expectedOutput, output)
        }
    }
}
func TestTabError(t *testing.T){
	cases := []struct {
        input          string
        expectedOutput int
    }{
        {"", 0},
        {"C:/Users/Anthony/go/APL/src/360-anthony-mastores/sprint-2",0},
    }
	var t1 tE.TabError
    for _, c := range cases {
        if output := t1.HandleError(c.input); output != c.expectedOutput {
            t.Errorf("incorrect output for `%s`: expected `%d` but got `%d`", c.input, c.expectedOutput, output)
        }
    }
}