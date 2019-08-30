package conf

import(
	"github.com/be4119c8/utility/conf/common"
)

func New(e func(string)(common.Conf,error),path string) (common.Conf,error){
	conf,err := e(path)
	if err != nil {
		return nil,err
	}
	return conf,nil
}
