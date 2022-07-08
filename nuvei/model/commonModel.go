package model

type BillingAddress struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Address   string `json:"address,omitempty"`
	Cell      string `json:"cell,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Zip       string `json:"zip,omitempty"`
	City      string `json:"city,omitempty"`
	Country   string `json:"country"`
	State     string `json:"state,omitempty"`
	Email     string `json:"email"`
	County    string `json:"county,omitempty"`
}

type ExternalSchemeDetails struct {
	TransactionId string `json:"transactionId"`
	Brand         string `json:"brand"`
}

type SubMerchant struct {
	CountryCode string `json:"countryCode"`
	City        string `json:"city"`
	Id          string `json:"id"`
}

type PaymentOption struct {
	UserPaymentOptionId      string                   `json:"userPaymentOptionId,omitempty"`
	Card                     Card                     `json:"card,omitempty"`
	AlternativePaymentMethod AlternativePaymentMethod `json:"alternativePaymentMethod,omitempty"`
}

type Card struct {
	CardNumber      string `json:"cardNumber"`
	CardHolderName  string `json:"cardHolderName"`
	ExpirationMonth string `json:"expirationMonth"`
	ExpirationYear  string `json:"expirationYear"`
	CVV             string `json:"CVV"`
	ThreeD          ThreeD `json:"threeD"`
}

type ThreeD struct {
	BrowserDetails     BrowserDetails     `json:"browserDetails"`
	Version            string             `json:"version"`
	NotificationUrl    string             `json:"notificationUrl"`
	MerchantUrl        string             `json:"merchantUrl"`
	PlatformType       string             `json:"platformType"`
	V2AdditionalParams V2AdditionalParams `json:"v2AdditionalParams"`
}

type V2AdditionalParams struct {
	ChallengeWindowSize string `json:"challengeWindowSize"`
}

type DeviceDetails struct {
	IP string `json:"ipAddress"`
}

type AmountDetails struct {
	TotalShipping string `json:"totalShipping"`
	TotalHandling string `json:"totalHandling"`
	TotalDiscount string `json:"totalDiscount"`
	TotalTax      string `json:"totalTax"`
}

type Item struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}

type BrowserDetails struct {
	AcceptHeader      string `json:"acceptHeader"`
	Ip                string `json:"ip"`
	JavaEnabled       string `json:"javaEnabled"`
	JavaScriptEnabled string `json:"javaScriptEnabled"`
	Language          string `json:"language"`
	ColorDepth        string `json:"colorDepth"`
	ScreenHeight      string `json:"screenHeight"`
	ScreenWidth       string `json:"screenWidth"`
	TimeZone          string `json:"timeZone"`
	UserAgent         string `json:"userAgent"`
}

type FraudDetails struct {
	FinalDecision string `json:"finalDecision"`
}

type UserPaymentOption struct {
	UserPaymentOptionId string `json:"userPaymentOptionId"`
}

type CardData struct {
	AcquirerId string `json:"acquirerId"`
	VisaDirect string `json:"visaDirect"`
}

type MerchantDetails struct {
	CustomField1 string `json:"customField1"`
}

type SessionToken struct {
	SessionToken string `json:"sessionToken"`
}

type AlternativePaymentMethod struct {
	ExternalAccountID     string `json:"externalAccountID"`
	ExternalTransactionId string `json:"externalTransactionId"`
	APMReferenceID        string `json:"APMReferenceID"`
	OrderTransactionId    string `json:"orderTransactionId"`
	ApmPayerInfo          string `json:"apmPayerInfo"`
	PaymentMethod         string `json:"paymentMethod"`
}

type PaymentMethod struct {
	UserPaymentOptionId string         `json:"userPaymentOptionId"`
	PaymentMethodName   string         `json:"paymentMethodName"`
	UpoName             string         `json:"upoName"`
	UpoRegistrationDate string         `json:"upoRegistrationDate"`
	UpoStatus           string         `json:"upoStatus,omitempty"`
	ExpiryDate          string         `json:"expiryDate"`
	BillingAddress      BillingAddress `json:"billingAddress"`
	UpoData             UpoData        `json:"upoData"`
	UpoStatus1          string         `json:"upoStatus ,omitempty"`
}

type UpoData struct {
	CardType         string `json:"cardType,omitempty"`
	CcCardNumber     string `json:"ccCardNumber,omitempty"`
	CcCardNumberHash string `json:"ccCardNumberHash,omitempty"`
	CcExpMonth       string `json:"ccExpMonth,omitempty"`
	CcExpYear        string `json:"ccExpYear,omitempty"`
	CcNameOnCard     string `json:"ccNameOnCard,omitempty"`
	CcToken          string `json:"ccToken,omitempty"`
	Brand            string `json:"brand,omitempty"`
	UniqueCC         string `json:"uniqueCC,omitempty"`
	Bin              string `json:"bin,omitempty"`
	Last4Digits      string `json:"last4Digits,omitempty"`
	AccountId        string `json:"account_id,omitempty"`
}
