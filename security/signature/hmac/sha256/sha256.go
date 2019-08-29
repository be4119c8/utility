package sha256

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/be4119c8/utility/security/signature/encoder"
)

type HmacSha256 struct {
	data []byte
	key string
}

func (h *HmacSha256) Encode( ) string {

	hmac := hmac.New(sha256.New, []byte(h.key))
	hmac.Write( h.data )
	bs := hmac.Sum(nil)
	return base64.StdEncoding.EncodeToString(bs)
}

func( h *HmacSha256) Init( data []byte, key string ){
	h.data = data
	h.key = key
}


func New( ) encoder.Encoder {
	return new(HmacSha256)
}