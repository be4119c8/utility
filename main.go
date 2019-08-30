package main

import (
	"fmt"
	"github.com/be4119c8/utility/conf"
	"github.com/be4119c8/utility/conf/yaml"
	"github.com/be4119c8/utility/payment/weixin/app"
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
	conf,err := conf.New(yaml.New,"/Users/seca/Documents/github/redis_data_migration/conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	cf := new(SysConf)
	err = conf.GetConf(cf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(cf)
}

func main(){

	Conf_Example()
	return

	unifiedorde := app.NewUnifiedorderRequest();
	err := unifiedorde.SetRequiredParams("wx5dc57154b4af07f0","1273119401","测试","1","test_001","ip","url","APP")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println( unifiedorde )
	err = unifiedorde.CheckRequestParams()
	if err != nil {
		fmt.Println(err.Error())
	}

	//signature.Qcloud_example();
}
