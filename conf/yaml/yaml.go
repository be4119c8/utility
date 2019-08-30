package yaml

import (
	"errors"
	"fmt"
	"github.com/be4119c8/utility/conf/common"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
)

type YamlConf struct {
	path string
}


func (yamlConf YamlConf) GetConf(t interface{}) error {
	v := reflect.ValueOf(t)
	var err error
	fmt.Println(v.Kind().String())
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
	err = errors.New("GetConf Params Error!! Params require is ptr.")
	return err
}

func New(path string) (common.Conf,error){
	if path == "" {
		return nil,errors.New("Path is empty!")
	}

	yamlConf := YamlConf{
		path,
	}
	return yamlConf,nil
}