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
		fmt.Println("1. default declearation 'var errs multiError' is nil, print value is ==> ", errs) // it equals to nil
	}
	errs2 := multiError{}

	if errs2 != nil {
		fmt.Println("2. empty declearation 'errs2 := multiError{}' is not nil, print value is ==> ", errs2) // it equals to nil
	}
	if i == nil {
		fmt.Println("3. default declearation of 'var i []int' i is nil, print value is  ", i)
	}
	fmt.Println("5. print value of \n default error e is ", e, "\n combine is ", Combine(e))

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
		fmt.Println("4. default value of len(errs) == 0 errs = ", errs)
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
