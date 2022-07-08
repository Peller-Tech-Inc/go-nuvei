package nuvei

import (
	"github.com/Peller-Tech-Inc/go-nuvei/nuvei/util"
)

type user struct {
	*nuveiObject
	UserTokenId string `json:"userTokenId"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Address     string `json:"address"`
	State       string `json:"state"`
	City        string `json:"city"`
	Zip         string `json:"zip"`
	CountryCode string `json:"countryCode"`
	Phone       string `json:"phone"`
	Locale      string `json:"locale"`
	Email       string `json:"email"`
	DateOfBirth string `json:"dateOfBirth"`
	County      string `json:"county"`
}

func NewUser(clientRequestId string, userTokenId string, firstName string, lastName string, address string, state string,
	city string, zip string, countryCode string, phone string, locale string, email string, dob string, county string) *user {
	result := &user{
		nuveiObject: NewNuveiObject(clientRequestId),
		UserTokenId: userTokenId,
		FirstName:   firstName,
		LastName:    lastName,
		Address:     address,
		State:       state,
		City:        city,
		Zip:         zip,
		CountryCode: countryCode,
		Phone:       phone,
		Locale:      locale,
		Email:       email,
		DateOfBirth: dob,
		County:      county,
	}
	result.nuveiObject.Checksum = util.SHA256(result.MerchantId + result.MerchantSiteId + result.UserTokenId + result.ClientRequestId +
		firstName + lastName + address + state + city + zip + countryCode + phone + locale + email + county + result.TimeStamp + instance.secret)
	return result
}
