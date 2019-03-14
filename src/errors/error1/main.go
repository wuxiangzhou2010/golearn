package main

// 1. error() string interface
// 2. r := strings.Builder(),  r.WriteString(""), r.String()
// 3. string function is called when fmt an error
import (
	"fmt"
	"strings"
)

type multiError []error

func main() {

	// fmt.Println(Combine(errors.New("first error"), errors.New("second error")))
	var e error         // this is a nil error  print value is nil
	var i []int         // this is a nil []int  print value is []
	var errs multiError // this is a nil errs, print value is multierr
	if errs == nil {
		fmt.Println("errs is nil ", errs) // it equals to nil
	}
	errs2 := multiError{}

	if errs2 != nil {
		fmt.Println("errs2 is not nil", errs2) // it equals to nil
	}
	if i == nil {
		fmt.Println("i is nil ")
	}
	fmt.Println("e is ", e, "combine is ", Combine(e), "ints ", i)
	fmt.Printf("ints %+v\n", i)
}
func (e multiError) Error() string {
	// if e == nil {
	// 	return ""
	// }
	var r strings.Builder
	r.WriteString("multierr:")
	for _, err := range e {
		r.WriteString(err.Error())
		r.WriteString(" | ")
	}
	return r.String()
}

// Combine output multiple errors
func Combine(maybeError ...error) error {
	var errs multiError
	for _, err := range maybeError {
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) == 0 {
		fmt.Println("len(errs) == 0 errs = ", errs)
		return nil
	}
	return errs
}

// :::output:::

// errs is nil  multierr:
// errs2 is not nil multierr:
// len(errs) == 0 errs =  multierr:
// e is  <nil> combine is  <nil> ints  []
// ints []
