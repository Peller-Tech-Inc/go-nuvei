package nuvei

import (
	"go-nuvei/nuvei/model"
	"go-nuvei/nuvei/util"
)

type payment struct {
	*nuveiObject
	SessionToken   string               `json:"sessionToken"`
	UserTokenId    string               `json:"userTokenId,omitempty"`
	ClientUniqueId string               `json:"clientUniqueId,omitempty"`
	Amount         string               `json:"amount"`
	Currency       string               `json:"currency"`
	PaymentOption  model.PaymentOption  `json:"paymentOption"`
	DeviceDetails  model.DeviceDetails  `json:"deviceDetails"`
	BillingAddress model.BillingAddress `json:"billingAddress"`
}

func NewPayment(clientRequestId string, clientUniqueId string, currency string, amount string, userTokenId string,
	billingAddress model.BillingAddress, sessionToken string, paymentOption model.PaymentOption, details model.DeviceDetails) *payment {
	result := &payment{
		nuveiObject:    NewNuveiObject(clientRequestId),
		SessionToken:   sessionToken,
		UserTokenId:    userTokenId,
		ClientUniqueId: clientUniqueId,
		Amount:         amount,
		Currency:       currency,
		PaymentOption:  paymentOption,
		DeviceDetails:  details,
		BillingAddress: billingAddress,
	}
	result.nuveiObject.Checksum = util.SHA256(result.MerchantId + result.MerchantSiteId + result.ClientRequestId +
		result.Amount + result.Currency + result.TimeStamp + instance.secret)
	return result
}
