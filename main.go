package main

import (
	"fmt"
	"github.com/be4119c8/utility/payment/weixin/app"
)

func main(){
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
