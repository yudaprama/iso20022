package model

// Data related to the authentication of the card and the cardholder.
type CardholderAuthentication8 struct {

	// Method and data intended to be used for this transaction to authenticate the customer or its card.
	AuthenticationMethod *AuthenticationMethod7Code `xml:"AuthntcnMtd"`

	// True if an authentication token is requested to the host. This token will be provided to the ATM for further authentication.
	TokenRequested *TrueFalseIndicator `xml:"TknReqd,omitempty"`

	// Value or token to be used for customer or card authentication.
	AuthenticationValue *Max5000Binary `xml:"AuthntcnVal,omitempty"`

	// Protection of the authentication value.
	ProtectedAuthenticationValue *ContentInformationType10 `xml:"PrtctdAuthntcnVal,omitempty"`

	// Encrypted personal identification number (PIN) and related information.
	CardholderOnLinePIN *OnLinePIN5 `xml:"CrdhldrOnLinePIN,omitempty"`
}

func (c *CardholderAuthentication8) SetAuthenticationMethod(value string) {
	c.AuthenticationMethod = (*AuthenticationMethod7Code)(&value)
}

func (c *CardholderAuthentication8) SetTokenRequested(value string) {
	c.TokenRequested = (*TrueFalseIndicator)(&value)
}

func (c *CardholderAuthentication8) SetAuthenticationValue(value string) {
	c.AuthenticationValue = (*Max5000Binary)(&value)
}

func (c *CardholderAuthentication8) AddProtectedAuthenticationValue() *ContentInformationType10 {
	c.ProtectedAuthenticationValue = new(ContentInformationType10)
	return c.ProtectedAuthenticationValue
}

func (c *CardholderAuthentication8) AddCardholderOnLinePIN() *OnLinePIN5 {
	c.CardholderOnLinePIN = new(OnLinePIN5)
	return c.CardholderOnLinePIN
}
