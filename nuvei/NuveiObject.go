package nuvei

import (
	"go-nuvei/nuvei/util"
	"time"
)

type nuveiObject struct {
	MerchantId      string `json:"merchantId"`
	MerchantSiteId  string `json:"merchantSiteId"`
	ClientRequestId string `json:"clientRequestId,omitempty"`
	TimeStamp       string `json:"timeStamp"`
	Checksum        string `json:"checksum"`
}

func NewNuveiObject(clientRequestId string) *nuveiObject {
	return &nuveiObject{
		MerchantId:      instance.merchantId,
		MerchantSiteId:  instance.merchantSiteId,
		TimeStamp:       time.Now().Format("20060201150405"),
		ClientRequestId: clientRequestId,
	}
}

func GenerateNuveiObject(clientRequestId string) *nuveiObject {
	result := NewNuveiObject(clientRequestId)
	result.Checksum = util.SHA256(result.MerchantId + result.MerchantSiteId + result.ClientRequestId + result.TimeStamp + instance.secret)
	return result
}
