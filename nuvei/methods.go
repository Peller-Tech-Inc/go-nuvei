package nuvei

import (
	"encoding/json"
	"github.com/Peller-Tech-Inc/go-nuvei/nuvei/model"
	"log"
)

func GetSessionToken(clientRequestId string) (*model.SessionTokenResponse, error) {
	body := GenerateNuveiObject(clientRequestId)
	resp, err := sendRequest(body, model.GetSessionToken, &model.SessionTokenResponse{})
	if err != nil {
		return nil, err
	}
	return resp.(*model.SessionTokenResponse), nil
}

func OpenOrder(clientRequestId string, clientUniqueId string, currency string, amount string, userTokenId string,
	billingAddress model.BillingAddress, isRebilling int, details model.ExternalSchemeDetails, authOnlyType string) (*model.OpenOrderResponse, error) {
	body := NewOrder(clientRequestId, clientUniqueId, currency, amount, userTokenId, billingAddress, isRebilling, details, authOnlyType)
	resp, err := sendRequest(body, model.OpenOrder, &model.OpenOrderResponse{})
	if err != nil {
		return nil, err
	}
	return resp.(*model.OpenOrderResponse), nil
}

func InitPayment(clientRequestId string, clientUniqueId string, currency string, amount string, userTokenId string,
	billingAddress model.BillingAddress, sessionToken string, paymentOption model.PaymentOption, details model.DeviceDetails) (*model.InitPaymentResponse, error) {
	body := NewPayment(clientRequestId, clientUniqueId, currency, amount, userTokenId, billingAddress, sessionToken, paymentOption, details)
	resp, err := sendRequest(body, model.InitPayment, &model.InitPaymentResponse{})
	if err != nil {
		return nil, err
	}
	return resp.(*model.InitPaymentResponse), nil
}

func Pay(clientRequestId string, clientUniqueId string, currency string, amount string, userTokenId string,
	billingAddress model.BillingAddress, sessionToken string, paymentOption model.PaymentOption, details model.DeviceDetails) (*model.PaymentResponse, error) {
	body := NewPayment(clientRequestId, clientUniqueId, currency, amount, userTokenId, billingAddress, sessionToken, paymentOption, details)
	resp, err := sendRequest(body, model.Pay, &model.PaymentResponse{})
	if err != nil {
		return nil, err
	}
	return resp.(*model.PaymentResponse), nil
}

func Payout(clientRequestId string, userTokenId string, clientUniqueId string, amount string, currency string,
	details model.DeviceDetails, uop model.UserPaymentOption) (*model.PayoutResponse, error) {
	body := NewPayout(clientRequestId, userTokenId, clientUniqueId, amount, currency, details, uop)
	resp, err := sendRequest(body, model.Payout, &model.PayoutResponse{})
	if err != nil {
		return nil, err
	}
	return resp.(*model.PayoutResponse), nil
}

func GetPaymentStatus(sessionToken string) (*model.PaymentStatusResponse, error) {
	body := model.SessionToken{SessionToken: sessionToken}
	resp, err := sendRequest(body, model.GetPaymentStatus, &model.PaymentStatusResponse{})
	if err != nil {
		return nil, err
	}
	return resp.(*model.PaymentStatusResponse), nil
}

func GetPayoutStatus(clientRequestId string) (*model.PayoutStatusResponse, error) {
	body := NewNuveiObject(clientRequestId)
	resp, err := sendRequest(body, model.GetPayoutStatus, &model.PayoutStatusResponse{})
	if err != nil {
		return nil, err
	}
	return resp.(*model.PayoutStatusResponse), nil
}

func CreateUser(clientRequestId string, userTokenId string, firstName string, lastName string, address string, state string,
	city string, zip string, countryCode string, phone string, locale string, email string, dob string, county string) (*model.CreateUserResponse, error) {
	body := NewUser(clientRequestId, userTokenId, firstName, lastName, address, state, city, zip, countryCode, phone,
		locale, email, dob, county)
	resp, err := sendRequest(body, model.CreateUser, &model.CreateUserResponse{})
	if err != nil {
		return nil, err
	}
	return resp.(*model.CreateUserResponse), nil
}

func AddUPOByTempToken(clientRequestId string, sessionToken string, userTokenId string, ccTempToken string,
	address model.BillingAddress) (*model.AddUPOResponse, error) {
	body := NewUPO(clientRequestId, sessionToken, userTokenId, ccTempToken, address)
	resp, err := sendRequest(body, model.AddUPOByTempToken, &model.AddUPOResponse{})
	if err != nil {
		return nil, err
	}
	return resp.(*model.AddUPOResponse), nil
}

func GetUPOs(userTokenId string) (*model.UPOsResponse, error) {
	UPOsRequest := NewUPO("", "", userTokenId, "", model.BillingAddress{})
	resp, err := sendRequest(UPOsRequest, model.GetUPOs, &model.UPOsResponse{})
	if err != nil {
		return nil, err
	}
	return resp.(*model.UPOsResponse), nil
}

func sendRequest(model interface{}, action model.NuveiAction, responseType interface{}) (interface{}, error) {
	body, err := json.Marshal(model)
	if err != nil {
		log.Printf("Can't create payment request body: %v", err)
		return nil, err
	}
	resp, err := CreateRequest(body, action, responseType)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
