package reflection

import (
	"fmt"
	"reflect"
)

type order struct {
	id         int
	customerId int
}

type customer struct {
	id      int
	name    string
	address string
}

func buildQuery(object interface{}) (query string, err error) {
	switch reflect.ValueOf(object).Kind() {
	case reflect.Struct:
		return buildQueryForStruct(object)
	default:
		return "", fmt.Errorf("unsupported kind of object: %v", reflect.ValueOf(object).Kind())
	}
}

func buildQueryForStruct(object interface{}) (query string, err error) {
	v := reflect.ValueOf(object)
	if v.NumField() < 1 {
		return "", fmt.Errorf("no fields in struct")
	}

	// Append 1st field
	name := reflect.TypeOf(object).Name()
	part, err := buildValuePart(v.Field(0))
	if err != nil {
		return "", err
	}
	query = fmt.Sprintf("INSERT INTO %s VALUES(%s", name, part)

	// Append remaining fields
	for i := 1; i < v.NumField(); i++ {
		p, err := buildValuePart(v.Field(i))
		if err != nil {
			return "", err
		}
		query = query + ", " + p
	}

	// Close parenthesis
	query = query + ")"
	return query, nil
}

func buildValuePart(field reflect.Value) (part string, err error) {
	switch field.Kind() {
	case reflect.Int:
		return fmt.Sprintf("%d", field.Int()), nil
	case reflect.String:
		return field.String(), nil
	default:
		return "", fmt.Errorf("unsupported kind of field: %v", field.Kind())
	}
}
