package utils

import (
	"log"
	"reflect"
)

// CreateNewInstance creates a new instance of a type passed as a reflect.Type.
func CreateNewInstance(t reflect.Type) interface{} {
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		log.Fatalf("CreateNewInstance expects a pointer to a struct, got %v", t)
	}
	return reflect.New(t.Elem()).Interface()
}
