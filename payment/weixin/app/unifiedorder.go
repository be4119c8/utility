package app

import (
	"github.com/be4119c8/utility/payment/weixin"
	"github.com/be4119c8/utility/tag"
	"math/rand"
	"strconv"
	"time"
)

type UnifiedorderRequest struct {
	appid          string `be4119c8:"appid,required,max=32"`
	mchid          string `be4119c8:"mch_id,required,max=32"`
	deviceInfo     string `be4119c8:"device_info,max=32"`
	nonceStr       string `be4119c8:"nonce_str,required,max=32"`
	sign           string `be4119c8:"sign,max=32"`
	signType       string `be4119c8:"sign_type,default=HMAC-SHA256,max=32"`
	body           string `be4119c8:"body,required,max=128"`
	detail         string `be4119c8:"detail,max=8192"`
	attach         string `be4119c8:"accach,max=127"`
	outTraceNo     string `be4119c8:"out_trade_no,max=32,required"`
	feeType        string `be4119c8:"fee_type,max=16,default=CNY"`
	totalFee       int    `be4119c8:"total_fee,required"`
	spbillCreateIp string `be4119c8:"spbill_create_ip,max=64,required"`
	timeStart      string `be4119c8:"time_start,max=14"`
	timeExpire     string `be4119c8:"time_expire,max=14"`
	goodsTag       string `be4119c8:"goods_tag,max=32"`
	notifyUrl      string `be4119c8:"notify_url,max=256,required"`
	tradeType      string `be4119c8:"trade_type,max=16,required"`
	limitPay       string `be4119c8:"limit_pey,max=32"`
	receipt        string `be4119c8:"receipt,max=8"`
}

type UnifiedOrderResponse struct {
	returnCode string `xml:"return_code"`
	returnMsg  string `xml:"return_msg"`
	appid      string `xml:"appid"`
	mchid      string `xml:"mch_id"`
	deviceInfo string `xml:"device_info"`
	nonceStr   string `xml:"nonce_str"`
	sign       string
	resultCode string `xml:"result_code"`
	errCode    string `xml:"err_code"`
	errCodeDes string `xml:"err_code_des"`
	tradeType  string `xml:"trade_type"`
	prepayId   string `xml:"prepay_id"`
}

func NewUnifiedorderRequest() *UnifiedorderRequest {
	return &UnifiedorderRequest{}
}

func (req *UnifiedorderRequest) SetRequiredParams(
	appid,
	mchid,
	body,
	totalFee,
	outTradeNo,
	spbillCreateIp,
	notifyUrl,
	tradeType string) error {

	if appid == "" || mchid == "" || body == "" || totalFee == "" || outTradeNo == "" || spbillCreateIp == "" || notifyUrl == "" {
		return weixin.ErrParamError
	}
	if tradeType == "" {
		req.tradeType = "APP"
	} else {
		req.tradeType = tradeType
	}
	req.appid = appid
	req.mchid = mchid
	req.body = body
	req.totalFee,_ = strconv.Atoi(totalFee)
	req.outTraceNo = outTradeNo
	req.spbillCreateIp = spbillCreateIp
	req.notifyUrl = notifyUrl
	rand.Seed(time.Now().UnixNano())
	req.nonceStr = strconv.Itoa(rand.Int())
	return nil
}

func (req *UnifiedorderRequest) SetExtParams(
	deviceInfo,
	signType,
	detail,
	attach,
	feeType,
	timeStart,
	timeExpire,
	goodsTag,
	limitPay,
	receipt string) {
	if deviceInfo != "" {
		req.deviceInfo = deviceInfo
	} else {
		req.deviceInfo = "WEB"
	}

	if signType != "" {
		req.signType = signType
	} else {
		req.signType = "HMAC-SHA256"
	}

	if detail != "" {
		req.detail = detail
	}

	if attach != "" {
		req.attach = attach
	}

	if feeType != "" {
		req.feeType = feeType
	} else {
		req.feeType = "CNY"
	}

	if timeStart != "" {
		req.timeStart = timeStart
	}

	if timeExpire != "" {
		req.timeExpire = timeExpire
	}

	if goodsTag != "" {
		req.goodsTag = goodsTag
	}

	if limitPay != "" {
		req.limitPay = limitPay
	}

	if receipt != "" {
		req.receipt = receipt
	}
}

func (req *UnifiedorderRequest) CheckRequestParams () error {
	err := tag.CheckStructByTag(req,tag.TAGNAME)
	if err != nil {
		return err
	}
	return nil
}
