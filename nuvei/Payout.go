package nuvei

import (
	"go-nuvei/nuvei/model"
	"go-nuvei/nuvei/util"
)

type payout struct {
	*nuveiObject
	UserTokenId       string                  `json:"userTokenId"`
	ClientUniqueId    string                  `json:"clientUniqueId"`
	Amount            string                  `json:"amount"`
	Currency          string                  `json:"currency"`
	DeviceDetails     model.DeviceDetails     `json:"deviceDetails"`
	UserPaymentOption model.UserPaymentOption `json:"userPaymentOption"`
}

func NewPayout(clientRequestId string, userTokenId string, clientUniqueId string, amount string, currency string,
	details model.DeviceDetails, uop model.UserPaymentOption) *payout {
	result := &payout{
		nuveiObject:       NewNuveiObject(clientRequestId),
		UserTokenId:       userTokenId,
		ClientUniqueId:    clientUniqueId,
		Amount:            amount,
		Currency:          currency,
		DeviceDetails:     details,
		UserPaymentOption: uop,
	}
	result.nuveiObject.Checksum = util.SHA256(result.MerchantId + result.MerchantSiteId + result.ClientRequestId +
		result.Amount + result.Currency + result.TimeStamp + instance.secret)
	return result
}
