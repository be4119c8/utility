package common

type Conf interface {
	GetConf(t interface{}) error
}