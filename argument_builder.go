package gofreerdp

import (
	"fmt"
	"reflect"
	"strings"
)

type ArgumentBuilder interface {
	argumentBuild() string
}

func argumentBuild(obj interface{}) string {
	var parts []string
	val := reflect.ValueOf(obj).Elem() // Get the value of the struct (dereference the pointer)
	typ := val.Type()                  // Get the type of the struct

	// Get the struct name (type name)
	structName := typ.Name()

	// Iterate through the fields of the struct
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := val.Type().Field(i).Name // Get the field name
		fieldValue := field.Interface()       // Get the field value

		// Convert field name to lowercase
		fieldName = strings.ToLower(fieldName)

		// Check if the field value is non-empty (string check for simplicity)
		if strVal, ok := fieldValue.(string); ok && strVal != "" {
			// Append the field name and value to parts
			parts = append(parts, fmt.Sprintf("%s:'%s'", fieldName, strVal))
		}
	}

	// Join all parts into a single string, separated by commas
	// Return the struct name as the prefix (e.g., /app:params)
	return fmt.Sprintf("/%s:%s", strings.ToLower(structName), strings.Join(parts, ", "))
}
