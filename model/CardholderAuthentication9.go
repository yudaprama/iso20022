package model

// Data related to the authentication of the cardholder.
type CardholderAuthentication9 struct {

	// Method and data intended to be used for this transaction to authenticate the cardholder or its card.
	AuthenticationMethod *AuthenticationMethod5Code `xml:"AuthntcnMtd"`

	// Value used to authenticate the cardholder.
	AuthenticationValue *Max5000Binary `xml:"AuthntcnVal,omitempty"`

	// Protection of the authentication value.
	ProtectedAuthenticationValue *ContentInformationType10 `xml:"PrtctdAuthntcnVal,omitempty"`

	// Encrypted personal identification number (PIN) and related information.
	CardholderOnLinePIN *OnLinePIN4 `xml:"CrdhldrOnLinePIN,omitempty"`

	// Identification of the cardholder to verify.
	CardholderIdentification *PersonIdentification11 `xml:"CrdhldrId,omitempty"`

	// Numeric characters of the cardholder's billing or shipping address for verification.
	AddressVerification *AddressVerification1 `xml:"AdrVrfctn,omitempty"`
}

func (c *CardholderAuthentication9) SetAuthenticationMethod(value string) {
	c.AuthenticationMethod = (*AuthenticationMethod5Code)(&value)
}

func (c *CardholderAuthentication9) SetAuthenticationValue(value string) {
	c.AuthenticationValue = (*Max5000Binary)(&value)
}

func (c *CardholderAuthentication9) AddProtectedAuthenticationValue() *ContentInformationType10 {
	c.ProtectedAuthenticationValue = new(ContentInformationType10)
	return c.ProtectedAuthenticationValue
}

func (c *CardholderAuthentication9) AddCardholderOnLinePIN() *OnLinePIN4 {
	c.CardholderOnLinePIN = new(OnLinePIN4)
	return c.CardholderOnLinePIN
}

func (c *CardholderAuthentication9) AddCardholderIdentification() *PersonIdentification11 {
	c.CardholderIdentification = new(PersonIdentification11)
	return c.CardholderIdentification
}

func (c *CardholderAuthentication9) AddAddressVerification() *AddressVerification1 {
	c.AddressVerification = new(AddressVerification1)
	return c.AddressVerification
}
