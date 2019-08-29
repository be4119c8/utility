package tag

import (
	"fmt"
	"reflect"
	"strings"
)

const TAGNAME = "be4119c8"

func CheckStructByTag(obj interface{}, tagname string) error {
	if tagname == "" {
		tagname = TAGNAME
	}
	t := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		val = val.Elem()
	}

	for i := 0; i< t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := val.Field(i)
		fieldType := field.Type.String()
		tag := field.Tag.Get(tagname)
		ary := strings.Split(tag, ",")
		fmt.Println(fieldType,tag,ary,fieldVal)
		for _,v := range ary{
			idx := strings.Index(v,"=")
			var key,value string
			if idx != -1 {
				key = v[:idx]
				value = v[idx+1:]
			} else {
				key = v
				value = ""
			}
			err := checkTagValue(fieldType,key,value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func checkTagValue(fieldType,key,value string) error {
	var err error
	switch key {
	case "required":
	case "defalut":
	case "max":
	case "min":
	default:

	}

	if err != nil {
		return err
	}
	return nil
}

//func checkIsNull(key,value string) error {
//	switch fieldType {
//	case "int","uint","int8","uint8","int16","uint16","int32","uint32","int64","uint64","rune","byte":
//		err = checkInt(key,value)
//	case "float32","float64","complex64","complex128":
//		err = checkFloat(key,value)
//	case "string":
//		err = checkString(key,value)
//	case "bool":
//		err = checkBool(key,value)
//	default:
//		err = checkDefault(key,value)
//
//	}
//	return nil
//}

func checkFloat(key,value string) error {
	return nil
}

func checkString(key,value string) error {
	return nil
}

func checkBool(key,value string) error {
	return nil
}

func checkDefault(key,value string) error {
	return nil
}
