package main_test

import (
	"reflect"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

func unsetUntestedFields(item interface{}) {
	val := reflect.Indirect(reflect.ValueOf(item))
	if val.Kind() != reflect.Struct {
		return
	}

	strFldNames := []string{"Id"}
	for _, name := range strFldNames {
		fv := val.FieldByName(name)
		if fv.IsValid() && fv.Kind() == reflect.String && fv.CanSet() {
			fv.SetString("")
		}
	}

	timeFldNames := []string{
		"LastLogin",
		"Created",
		"Updated",
		"Deleted",
		"Blocked",
	}
	t := new(timestamp.Timestamp)
	tt := reflect.TypeOf(t)

	for _, name := range timeFldNames {
		fv := val.FieldByName(name)
		if fv.IsValid() && fv.Type() == tt && fv.CanSet() {
			fv.Set(reflect.Zero(tt))
		}
	}
}
