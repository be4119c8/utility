package signature

import (
	"fmt"
	"math/rand"
	my_sha1 "github.com/be4119c8/utility/security/signature/hmac/sha1"
	my_sha256 "github.com/be4119c8/utility/security/signature/hmac/sha256"
	"github.com/be4119c8/utility/security/signature/qcloud"
	"strconv"
	"time"
)

func qcloud_example() {
	var params = make( map[string]string )
	params["Region"] = "all"
	params["Timestamp"] = strconv.Itoa(int( time.Now().Unix()))
	rand.Seed(time.Now().UnixNano())
	params["Nonce"] = strconv.Itoa(rand.Int())
	params["SecretId"] = "secret_id"
	sign := qcloud.New(my_sha256.New,"key","prefix")

	fmt.Println( sign.GetSignature( params, ))

	sign = qcloud.New(my_sha1.New,"key", "prefix")
	fmt.Println(sign.GetSignature(params))

}
