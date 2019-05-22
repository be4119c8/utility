package weixin

type Weixin struct{
	Appid string
	Mchid string
	SecuretKey string
	SignType string
}

func New(appid,mchid,securetKey,signtype string ) (*Weixin, error) {
	if appid == "" || mchid == "" || securetKey == "" {
		return nil,ErrParamError
	}
	if signtype == "" {
		signtype = "sha256"
	}
	return &Weixin {
		Appid:appid,
		Mchid:mchid,
		SecuretKey:securetKey,
		SignType:signtype,
	},nil
}