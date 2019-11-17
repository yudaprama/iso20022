package model

// Data related to the authentication of the cardholder.
type CardholderAuthentication6 struct {

	// Method to authenticate the cardholder.
	AuthenticationMethod *AuthenticationMethod3Code `xml:"AuthntcnMtd"`

	// Value used to authenticate the cardholder.
	AuthenticationValue *Max5000Binary `xml:"AuthntcnVal,omitempty"`

	// Protection of the authentication value.
	ProtectedAuthenticationValue *ContentInformationType10 `xml:"PrtctdAuthntcnVal,omitempty"`

	// Encrypted personal identification number (PIN) and related information.
	CardholderOnLinePIN *OnLinePIN4 `xml:"CrdhldrOnLinePIN,omitempty"`

	// Numeric characters of the cardholder's billing or shipping address for verification.
	AddressVerification *AddressVerification1 `xml:"AdrVrfctn,omitempty"`
}

func (c *CardholderAuthentication6) SetAuthenticationMethod(value string) {
	c.AuthenticationMethod = (*AuthenticationMethod3Code)(&value)
}

func (c *CardholderAuthentication6) SetAuthenticationValue(value string) {
	c.AuthenticationValue = (*Max5000Binary)(&value)
}

func (c *CardholderAuthentication6) AddProtectedAuthenticationValue() *ContentInformationType10 {
	c.ProtectedAuthenticationValue = new(ContentInformationType10)
	return c.ProtectedAuthenticationValue
}

func (c *CardholderAuthentication6) AddCardholderOnLinePIN() *OnLinePIN4 {
	c.CardholderOnLinePIN = new(OnLinePIN4)
	return c.CardholderOnLinePIN
}

func (c *CardholderAuthentication6) AddAddressVerification() *AddressVerification1 {
	c.AddressVerification = new(AddressVerification1)
	return c.AddressVerification
}
