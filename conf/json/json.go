package json

import (
	"encoding/json"
	"errors"
	"github.com/be4119c8/utility/conf"
	"io/ioutil"
	"reflect"
)

type JsonConf struct {
	 path string
}


func (jsConf JsonConf) GetConf(t interface{}) error {
	v := reflect.ValueOf(t)
	var err error
	if v.Kind() == reflect.Ptr && !v.IsNil(){
		content,err := ioutil.ReadFile(jsConf.path)
		if err != nil {
			return nil,err
		}
		err = json.Unmarshal(content,t)
		if err != nil {
			return nil,err
		}
		return t,nil
	}
	err = errors.New("GetConf Params Error!!")
	return nil,err
}

func New(path string) (conf.Conf,error){
	if path == "" {
		return nil,errors.New("Path is empty!")
	}

	jsonConf := JsonConf{
		path,
	}
	return jsonConf,nil
}