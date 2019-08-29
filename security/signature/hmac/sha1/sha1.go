package sha1

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"github.com/be4119c8/utility/security/signature/encoder"
)

type HmacSha1 struct {
	data []byte
	key string
}



func (h *HmacSha1) Encode( ) string {

	hmac := hmac.New(sha1.New, []byte(h.key))
	hmac.Write( h.data )
	bs := hmac.Sum(nil)
	return base64.StdEncoding.EncodeToString(bs)
}

func (h *HmacSha1) Init(data []byte, key string ){
	h.data = data
	h.key = key
}


func New() encoder.Encoder{
	return new(HmacSha1)
}