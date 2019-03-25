package main

// 1. ToString
// 2. Concat

import (
	"errors"
	"fmt"
	"strings"
)

// ToString trun a value into string
func ToString(v interface{}) string {
	if v == nil {
		return " "
	}
	switch value := v.(type) {
	case string:
		return value
	case *string:
		return *value
	case fmt.Stringer:
		return value.String()
	case error:
		return value.Error()
	default:
		return fmt.Sprintf("%+v", value)
	}
}

// Concat concatenate a variable length of values
func Concat(v ...interface{}) string {
	builder := strings.Builder{}

	for _, value := range v {
		builder.WriteString(ToString(value))
	}
	return builder.String()
}

func main() {
	a := []interface{}{"string", errors.New("error"), 32}

	fmt.Println(Concat(a...))
	fmt.Println(Concat(a))
}
