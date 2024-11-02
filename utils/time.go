package utils

import (
	"errors"
	"reflect"
	"time"
)

func Time(strct interface{}, new bool) error {
	t := reflect.TypeOf(strct)
	v := reflect.ValueOf(strct).Elem()

	// Validate if strct is a pointer and struct
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		return errors.New("strct must be a pointer to a struct")
	}

	// Validate if the datecreated and dateupdated fields are in strct and are of type time.Time
	dateCreatedField, has_created := t.Elem().FieldByName("DateCreated")
	dateUpdatedField, has_updated := t.Elem().FieldByName("DateUpdated")
	if !has_created || !has_updated {
		return errors.New("strct must have DateCreated and DateUpdated fields")
	}

	if dateCreatedField.Type.Kind() != reflect.TypeOf(time.Time{}).Kind() || dateUpdatedField.Type.Kind() != reflect.TypeOf(time.Time{}).Kind() {
		return errors.New("strct DateCreated and DateUpdated fields must be of type time.Time")
	}

	// Set the time for the fields based on new arguement value
	if new {
		// Set the "DateUpdated" field to current UTC time
		v.FieldByName("DateCreated").Set(reflect.ValueOf(time.Now().UTC()))
	}

	v.FieldByName("DateUpdated").Set(reflect.ValueOf(time.Now().UTC()))
	return nil

}
