package qcloud

import (
	"bytes"
	"github.com/be4119c8/utility/security/signature"
	"sort"
)

type Signature struct {
	securityKey string
	prefix string
	En signature.Encoder
}


func ( sign *Signature )GetSignature( params map[string]string) string {
	var sbuf bytes.Buffer
	var keys []string

	for k,_ := range params {
		keys = append( keys, k )
	}
	sort.Strings(keys)

	if sign.prefix != "" {
		sbuf.WriteString(sign.prefix)
	}

	for _,v := range keys {
		sbuf.WriteString( v + "=" + params[v] + "&")
	}
	data := bytes.TrimSuffix(sbuf.Bytes(), []byte("&"))

	sign.En.Init(data,sign.securityKey)
	return sign.En.Encode()
}

func (sign *Signature)CheckSignature(signature string, params map[string]string) bool {
	if signature == sign.GetSignature(params ) {
		return true
	}
	return false
}

func New( e func() signature.Encoder, key string, prefix string ) *Signature {
	sign := new( Signature )
	sign.En = e()
	sign.securityKey = key
	sign.prefix = prefix
	return sign
}