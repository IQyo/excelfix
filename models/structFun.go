package models

import (
	"reflect"
	"fmt"
)

//Map2Struct
func
SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		fmt.Println("structFieldType:",structFieldType)
		fmt.Println("val.Type():",val.Type())
		return fmt.Errorf("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}


//StructFunc
func (s *NanChangSrc) FillStruct(mapData map[string]interface{}) error {
	for k, v := range mapData {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *GuangZhouSrc) FillStruct(mapData map[string]interface{}) error {
	for k, v := range mapData {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *CalBianMa) FillStruct(mapData map[string]interface{}) error {
	for k, v := range mapData {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}


func (s *CalCompany) FillStruct(mapData map[string]interface{}) error {
	for k, v := range mapData {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

