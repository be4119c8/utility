package weixin

import "errors"

var (
	ErrOrderNotExist = errors.New("Weixin::Order not exist!")
	ErrSystemError	= errors.New("Weixin::System error,please try !")
	ErrInvalidRequest = errors.New("Weixin::invalid request params!")
	ErrNoAuth = errors.New( "Weixin::no auth!")
	ErrNotEnough = errors.New("Weixin::not enough!")
	ErrOrderPaid = errors.New("Weixin::order paid!")
	ErrOrderClosed = errors.New("Weixin::order closed!")
	ErrAppidNotExist = errors.New("Weixin::appid not exist!")
	ErrMchidNotExist = errors.New("Weixin::mchid not exist!")
	ErrAppIdMchidNotMatch = errors.New("Weixin::appid mchid not match!")
	ErrLackParams = errors.New("Weixin::lack params!")
	ErrOutTradeNoUses = errors.New("Weixin::out trade no used!")
	ErrSignError = errors.New("Weixin::sign error!")
	ErrXmlFormatError = errors.New("Weixin::xml format error!")
	ErrPostDataEmpty = errors.New("Weixin::post data empty!")
	ErrRequestPostMethod = errors.New("Weixin::request post method!")
	ErrNotUtf8 = errors.New("Weixin::not utf8!")
	ErrParamError = errors.New("Weixin::param error!")
	ErrTimeExpire = errors.New("Weixin::time expire!")
	ErrNoComment = errors.New("Weixin::no comment!")
	ErrNoBillExist = errors.New( "Weixin::no bill exist!")
	ErrBillCreating = errors.New("Weixin::bill creating!")
	)
