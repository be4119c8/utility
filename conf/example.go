package conf

import (
	"fmt"
	"github.com/be4119c8/utility/conf/yaml"
)

type RedisConf struct {
	Host string
	Port uint
	Auth string
}

type SysConf struct {
	SrcRedis RedisConf `yaml:"src-redis"`
	DestRedis []RedisConf `yaml:"dest-redis"`
}


func Conf_Example(){
	conf,err := New(yaml.New,"~/Documents/github/redis_data_migration/conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var cf SysConf
	err = conf.GetConf(cf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(cf.SrcRedis.Host)
}
