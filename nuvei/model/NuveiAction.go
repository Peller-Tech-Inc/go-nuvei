package model

type NuveiAction string

const (
	GetSessionToken   NuveiAction = "/getSessionToken.do"
	OpenOrder         NuveiAction = "/openOrder.do"
	InitPayment       NuveiAction = "/initPayment.do"
	Pay               NuveiAction = "/payment.do"
	Payout            NuveiAction = "/payout.do"
	GetPaymentStatus  NuveiAction = "/getPaymentStatus.do"
	GetPayoutStatus   NuveiAction = "/getPayoutStatus.do"
	CreateUser        NuveiAction = "/createUser.do"
	AddUPOByTempToken NuveiAction = "/addUPOCreditCardByTempToken.do"
	GetUPOs           NuveiAction = "/getUserUPOs.do"
	DeleteUPO         NuveiAction = "/deleteUPO.do"
)
