package model

// Data related to the authentication of the cardholder.
type CardholderAuthentication1 struct {

	// Method used to authenticate a cardholder.
	AuthenticationMethod *AuthenticationMethod1Code `xml:"AuthntcnMtd"`

	// Entity or object in charge of verifying the cardholder authenticity.
	AuthenticationEntity *AuthenticationEntity1Code `xml:"AuthntcnNtty"`

	// Value used to authenticate the cardholder.
	AuthenticationValue *Max40Text `xml:"AuthntcnVal,omitempty"`

	// Encrypted personal identification number (PIN) and related information.
	CardholderOnLinePIN *OnLinePIN1 `xml:"CrdhldrOnLinePIN,omitempty"`

	// Identifies in electronic commerce transactions whether customer authentication is supported and data is available.
	AuthenticationCollectionIndicator *Max35Text `xml:"AuthntcnColltnInd,omitempty"`
}

func (c *CardholderAuthentication1) SetAuthenticationMethod(value string) {
	c.AuthenticationMethod = (*AuthenticationMethod1Code)(&value)
}

func (c *CardholderAuthentication1) SetAuthenticationEntity(value string) {
	c.AuthenticationEntity = (*AuthenticationEntity1Code)(&value)
}

func (c *CardholderAuthentication1) SetAuthenticationValue(value string) {
	c.AuthenticationValue = (*Max40Text)(&value)
}

func (c *CardholderAuthentication1) AddCardholderOnLinePIN() *OnLinePIN1 {
	c.CardholderOnLinePIN = new(OnLinePIN1)
	return c.CardholderOnLinePIN
}

func (c *CardholderAuthentication1) SetAuthenticationCollectionIndicator(value string) {
	c.AuthenticationCollectionIndicator = (*Max35Text)(&value)
}
