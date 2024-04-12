package utils

import "reflect"

func GetFieldsTagsByName(name string, field reflect.StructField) string {
	tag, found := field.Tag.Lookup(name)

	if !found {
		return ""
	}
	return tag
}

func GetStructTagVals(name string, val any) []string {
	rv := reflect.TypeOf(val)

	if rv.Kind() != reflect.Struct {
		return nil
	}

	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	ret := []string{}
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		value := GetFieldsTagsByName(name, field)

		if value != "" {
			ret = append(ret, value)
		}
	}

	return ret
}