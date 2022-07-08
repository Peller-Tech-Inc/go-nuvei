package nuvei

import (
	"go-nuvei/nuvei/model"
	"go-nuvei/nuvei/util"
)

// TODO add remain fields here and in constructor
type order struct {
	*nuveiObject
	ClientUniqueId  string                      `json:"clientUniqueId,omitempty"`
	Currency        string                      `json:"currency"`
	Amount          string                      `json:"amount"`
	UserTokenId     string                      `json:"userTokenId"`
	BillingAddress  model.BillingAddress        `json:"billingAddress"`
	IsRebilling     int                         `json:"isRebilling,omitempty"`
	Details         model.ExternalSchemeDetails `json:"externalSchemeDetails,omitempty"`
	AuthOnlyType    string                      `json:"authenticationOnlyType,omitempty"`
	SubMerchant     model.SubMerchant           `json:"subMerchant,omitempty"`
	TransactionType string                      `json:"transactionType,omitempty"`
	PaymentOption   model.PaymentOption         `json:"paymentOption,omitempty"`
	DeviceDetails   model.DeviceDetails         `json:"deviceDetails,omitempty"`
	AmountDetails   model.AmountDetails         `json:"amountDetails,omitempty"`
	Items           []model.Item                `json:"items,omitempty"`
}

func NewOrder(clientRequestId string, clientUniqueId string, currency string, amount string, userTokenId string,
	billingAddress model.BillingAddress, isRebilling int, details model.ExternalSchemeDetails, authOnlyType string) *order {
	result := &order{
		nuveiObject:    NewNuveiObject(clientRequestId),
		ClientUniqueId: clientUniqueId,
		Amount:         amount,
		Currency:       currency,
		BillingAddress: billingAddress,
		UserTokenId:    userTokenId,
		IsRebilling:    isRebilling,
		Details:        details,
		AuthOnlyType:   authOnlyType,
	}
	result.nuveiObject.Checksum = util.SHA256(result.MerchantId + result.MerchantSiteId + result.ClientRequestId +
		result.Amount + result.Currency + result.TimeStamp + instance.secret)
	return result
}
