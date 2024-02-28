package helper

import (
	"errors"
	"reflect"
)

func ValidateRequiredFields(model interface{}) error {
	value := reflect.ValueOf(model)
	typ := reflect.TypeOf(model)

	if typ.Kind() == reflect.Ptr {
		value = value.Elem() // Obtém o valor subjacente da estrutura
		typ = typ.Elem()     // Obtém o tipo subjacente (a estrutura) do ponteiro
	}

	if typ.Kind() != reflect.Struct {
		return errors.New("model is not a struct")
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := value.Field(i)

		bindingTag, ok := field.Tag.Lookup("binding")
		if ok && bindingTag == "required" {
			if reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(field.Type).Interface()) {
				return errors.New("Required field is empty: " + field.Name)
			}
		}
	}

	return nil
}
