package json

import (
	"encoding/json"
	"errors"
	"github.com/be4119c8/utility/conf"
	"io/ioutil"
)

type JsonConf struct {
	 path string
}


func (jsConf JsonConf) GetConf(t interface{}) (interface{},error) {
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

func New(path string) (conf.Conf,error){
	if path == "" {
		return nil,errors.New("Path is empty!")
	}

	jsonConf := JsonConf{
		path,
	}
	return jsonConf,nil
}