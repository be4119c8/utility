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
	if t.Kind().String() == "ptr" {
		t = t.Elem()
	}
	for i := 0; i< t.NumField(); i++ {
		field := t.Field(i)
		fieldType := field.Type.String()
		tag := field.Tag.Get(tagname)
		ary := strings.Split(tag, ",")
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
	if fieldType != "string" && fieldType != "int" && fieldType != "uint" {
		return nil
	}
}