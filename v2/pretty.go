package pretty

import (
	"fmt"
	"reflect"
)

// Print prints an object recursively in an indented way. The indent size in spaces is specified as "indentSize"
func Print(resource interface{}, indentSize uint) {
	printRecursively(reflect.TypeOf(resource), reflect.ValueOf(resource), 0, indentSize)
}

func printRecursively(t reflect.Type, v reflect.Value, level, indentSize uint) {
	switch v.Kind() {

	case reflect.Struct:
		if _, hasStringer := t.MethodByName("String"); hasStringer {
			printIndented("%v\n", level, indentSize, v)
			return
		}
		for i := 0; i < v.NumField(); i++ {
			printIndented("%s:\n", level, indentSize, t.Field(i).Name)
			printRecursively(v.Field(i).Type(), v.Field(i), level+1, indentSize)
		}

	case reflect.Slice:
		count := v.Len()
		if count == 0 {
			printIndented("[no items]\n", level, indentSize)

		} else {
			for i := 0; i < count; i++ {
				printIndented("[%d]\n", level, indentSize, i)
				s := v.Index(i)
				printRecursively(s.Type(), s, level+1, indentSize)
			}
		}

	case reflect.Map:
		for _, key := range v.MapKeys() {
			printIndented("[%s]\n", level, indentSize, key)
			s := v.MapIndex(key)
			printRecursively(s.Type(), s, level+1, indentSize)
		}

	case reflect.Ptr, reflect.Interface:
		if v.IsNil() {
			printIndented("[empty]\n", level, indentSize)
		} else {
			printRecursively(v.Elem().Type(), v.Elem(), level, indentSize)
		}

	default:
		printIndented("%v\n", level, indentSize, v)
	}
}

func printIndented(format string, level, indentSize uint, args ...interface{}) {
	fmt.Printf("%*s", level*indentSize, "")
	fmt.Printf(format, args...)
}
