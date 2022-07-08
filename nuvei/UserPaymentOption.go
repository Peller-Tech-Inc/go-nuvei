package nuvei

import (
	"github.com/Peller-Tech-Inc/go-nuvei/nuvei/model"
)

type userPaymentOption struct {
	*nuveiObject
	SessionToken   string               `json:"sessionToken,omitempty"`
	UserTokenId    string               `json:"userTokenId"`
	CcTempToken    string               `json:"ccTempToken,omitempty"`
	BillingAddress model.BillingAddress `json:"billingAddress,omitempty"`
}

func NewUPO(clientRequestId string, sessionToken string, userTokenId string, ccTempToken string, address model.BillingAddress) *userPaymentOption {
	result := &userPaymentOption{
		nuveiObject:    GenerateNuveiObject(clientRequestId),
		SessionToken:   sessionToken,
		UserTokenId:    userTokenId,
		CcTempToken:    ccTempToken,
		BillingAddress: address,
	}
	return result
}
