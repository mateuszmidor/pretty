package pretty

import (
	"fmt"
	"reflect"
)

// Print prints an object recursively in an indented way. The indent is 2 spaces
func Print(resource interface{}) {
	printRecursively(reflect.TypeOf(resource), reflect.ValueOf(resource), 0)
}

func printRecursively(t reflect.Type, v reflect.Value, level int) {
	switch v.Kind() {

	case reflect.Struct:
		if _, hasStringer := t.MethodByName("String"); hasStringer {
			printIndented("%v\n", level, v)
			return
		}
		for i := 0; i < v.NumField(); i++ {
			printIndented("%s:\n", level, t.Field(i).Name)
			printRecursively(v.Field(i).Type(), v.Field(i), level+1)
		}

	case reflect.Slice:
		count := v.Len()
		if count == 0 {
			printIndented("[no items]\n", level)

		} else {
			for i := 0; i < count; i++ {
				printIndented("[%d]\n", level, i)
				s := v.Index(i)
				printRecursively(s.Type(), s, level+1)
			}
		}

	case reflect.Map:
		for _, key := range v.MapKeys() {
			printIndented("[%s]\n", level, key)
			s := v.MapIndex(key)
			printRecursively(s.Type(), s, level+1)
		}

	case reflect.Ptr, reflect.Interface:
		if v.IsNil() {
			printIndented("[empty]\n", level)
		} else {
			printRecursively(v.Elem().Type(), v.Elem(), level)
		}

	default:
		printIndented("%v\n", level, v)
	}
}

func printIndented(format string, level int, args ...interface{}) {
	fmt.Printf("%*s", level*2, "")
	fmt.Printf(format, args...)
}
