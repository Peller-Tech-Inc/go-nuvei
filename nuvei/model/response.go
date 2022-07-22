package model

type Response struct {
	InternalRequestId int    `json:"internalRequestId"`
	Status            string `json:"status"`
	ErrCode           int    `json:"errCode"`
	Reason            string `json:"reason"`
	MerchantId        string `json:"merchantId"`
	MerchantSiteId    string `json:"merchantSiteId"`
	Version           string `json:"version"`
	ClientRequestId   string `json:"clientRequestId,omitempty"`
}

type SessionTokenResponse struct {
	Response
	SessionToken string `json:"sessionToken"`
}

type OpenOrderResponse struct {
	Response
	OrderId        int    `json:"orderId"`
	SessionToken   string `json:"sessionToken"`
	ClientUniqueId string `json:"clientUniqueId"`
}

type InitPaymentResponse struct {
	Response
	OrderId             string        `json:"orderId"`
	TransactionStatus   string        `json:"transactionStatus"`
	CustomData          string        `json:"customData"`
	TransactionId       string        `json:"transactionId"`
	TransactionType     string        `json:"transactionType"`
	GwExtendedErrorCode int           `json:"gwExtendedErrorCode"`
	GwErrorCode         int           `json:"gwErrorCode"`
	ClientUniqueId      string        `json:"clientUniqueId"`
	PaymentOption       PaymentOption `json:"paymentOption"`
	SessionToken        string        `json:"sessionToken"`
	UserTokenId         string        `json:"userTokenId"`
}

type PaymentResponse struct {
	Response
	OrderId               string        `json:"orderId"`
	TransactionStatus     string        `json:"transactionStatus"`
	CustomData            string        `json:"customData"`
	TransactionId         string        `json:"transactionId"`
	TransactionType       string        `json:"transactionType"`
	GwExtendedErrorCode   int           `json:"gwExtendedErrorCode"`
	GwErrorCode           int           `json:"gwErrorCode"`
	GwErrorReason         string        `json:"gwErrorReason"`
	ClientUniqueId        string        `json:"clientUniqueId"`
	PaymentOption         PaymentOption `json:"paymentOption"`
	SessionToken          string        `json:"sessionToken"`
	UserTokenId           string        `json:"userTokenId"`
	ExternalTransactionId string        `json:"externalTransactionId"`
	AuthCode              string        `json:"authCode"`
	FraudDetails          FraudDetails  `json:"fraudDetails"`
}

type PayoutResponse struct {
	Response
	UserTokenId              string          `json:"userTokenId"`
	ClientUniqueId           string          `json:"clientUniqueId"`
	TransactionId            string          `json:"transactionId"`
	ExternalTransactionId    string          `json:"externalTransactionId"`
	CardData                 CardData        `json:"cardData"`
	TransactionStatus        string          `json:"transactionStatus"`
	MerchantDetails          MerchantDetails `json:"merchantDetails"`
	UserPaymentOptionId      string          `json:"userPaymentOptionId"`
	PaymentMethodErrorCode   string          `json:"paymentMethodErrorCode"`
	PaymentMethodErrorReason string          `json:"paymentMethodErrorReason"`
	GwErrorCode              string          `json:"gwErrorCode"`
	GwErrorReason            string          `json:"gwErrorReason"`
	GwExtendedErrorCode      string          `json:"gwExtendedErrorCode"`
}

type PaymentStatusResponse struct {
	Response
	TransactionType   string        `json:"transactionType"`
	TransactionStatus string        `json:"transactionStatus"`
	TransactionId     string        `json:"transactionId"`
	PaymentOption     PaymentOption `json:"paymentOption"`
	Currency          string        `json:"currency"`
	Amount            string        `json:"amount"`
	SessionToken      string        `json:"sessionToken"`
}

type PayoutStatusResponse struct {
	Response
	UserTokenId         string `json:"userTokenId"`
	TransactionId       string `json:"transactionId"`
	Amount              string `json:"amount"`
	Currency            string `json:"currency"`
	TransactionStatus   string `json:"transactionStatus"`
	UserPaymentOptionId string `json:"userPaymentOptionId"`
	GwErrorCode         string `json:"gwErrorCode"`
	GwErrorReason       string `json:"gwErrorReason"`
	GwExtendedErrorCode string `json:"gwExtendedErrorCode"`
}

type CreateUserResponse struct {
	Response
	UserId string `json:"userId"`
}

type AddUPOResponse struct {
	Response
	SessionToken        string `json:"SessionToken"`
	UserTokenId         string `json:"userTokenId"`
	UserPaymentOptionId string `json:"userPaymentOptionId"`
}

type UPOsResponse struct {
	Response
	PaymentMethods []PaymentMethod `json:"paymentMethods"`
}
