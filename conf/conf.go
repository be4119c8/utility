package conf

type Conf interface {
	GetConf(t interface{}) error
}

func New(e func(string)(Conf,error),path string) (Conf,error){
	conf,err := e(path)
	if err != nil {
		return nil,err
	}
	return conf,nil
}