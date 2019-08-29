package yaml

import (
	"gopkg.in/yaml.v2"
	"errors"
	"github.com/be4119c8/utility/conf"
	"io/ioutil"
	"reflect"
)

type YamlConf struct {
	path string
}


func (yamlConf YamlConf) GetConf(t interface{}) error {
	v := reflect.ValueOf(t)
	var err error
	if v.Kind() == reflect.Ptr && !v.IsNil() {
		content, err := ioutil.ReadFile(yamlConf.path)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(content, t)
		if err != nil {
			return err
		}
		return nil
	}
	err = errors.New("GetConf Params Error!!")
	return err
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