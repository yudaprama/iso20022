package model

// Data related to the authentication of the cardholder.
type CardholderAuthentication5 struct {

	// Method used to authenticate a cardholder.
	AuthenticationMethod *AuthenticationMethod2Code `xml:"AuthntcnMtd"`

	// Entity or object in charge of verifying the cardholder authenticity.
	AuthenticationEntity *AuthenticationEntity1Code `xml:"AuthntcnNtty,omitempty"`

	// Value used to authenticate the cardholder.
	AuthenticationValue *Max40Text `xml:"AuthntcnVal,omitempty"`

	// Encrypted personal identification number (PIN) and related information.
	CardholderOnLinePIN *OnLinePIN3 `xml:"CrdhldrOnLinePIN,omitempty"`

	// Identifies in electronic commerce transactions whether customer authentication is supported and data is available.
	AuthenticationCollectionIndicator *Max35Text `xml:"AuthntcnColltnInd,omitempty"`
}

func (c *CardholderAuthentication5) SetAuthenticationMethod(value string) {
	c.AuthenticationMethod = (*AuthenticationMethod2Code)(&value)
}

func (c *CardholderAuthentication5) SetAuthenticationEntity(value string) {
	c.AuthenticationEntity = (*AuthenticationEntity1Code)(&value)
}

func (c *CardholderAuthentication5) SetAuthenticationValue(value string) {
	c.AuthenticationValue = (*Max40Text)(&value)
}

func (c *CardholderAuthentication5) AddCardholderOnLinePIN() *OnLinePIN3 {
	c.CardholderOnLinePIN = new(OnLinePIN3)
	return c.CardholderOnLinePIN
}

func (c *CardholderAuthentication5) SetAuthenticationCollectionIndicator(value string) {
	c.AuthenticationCollectionIndicator = (*Max35Text)(&value)
}
