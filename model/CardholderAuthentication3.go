package model

// Data related to the authentication of the cardholder.
type CardholderAuthentication3 struct {

	// Method used to authenticate a cardholder.
	AuthenticationMethod *AuthenticationMethod2Code `xml:"AuthntcnMtd"`

	// Entity or object in charge of verifying the cardholder authenticity.
	AuthenticationEntity *AuthenticationEntity1Code `xml:"AuthntcnNtty,omitempty"`

	// Value used to authenticate the cardholder.
	AuthenticationValue *Max40Text `xml:"AuthntcnVal,omitempty"`

	// Encrypted personal identification number (PIN) and related information.
	CardholderOnLinePIN *OnLinePIN2 `xml:"CrdhldrOnLinePIN,omitempty"`

	// Identifies in electronic commerce transactions whether customer authentication is supported and data is available.
	AuthenticationCollectionIndicator *Max35Text `xml:"AuthntcnColltnInd,omitempty"`
}

func (c *CardholderAuthentication3) SetAuthenticationMethod(value string) {
	c.AuthenticationMethod = (*AuthenticationMethod2Code)(&value)
}

func (c *CardholderAuthentication3) SetAuthenticationEntity(value string) {
	c.AuthenticationEntity = (*AuthenticationEntity1Code)(&value)
}

func (c *CardholderAuthentication3) SetAuthenticationValue(value string) {
	c.AuthenticationValue = (*Max40Text)(&value)
}

func (c *CardholderAuthentication3) AddCardholderOnLinePIN() *OnLinePIN2 {
	c.CardholderOnLinePIN = new(OnLinePIN2)
	return c.CardholderOnLinePIN
}

func (c *CardholderAuthentication3) SetAuthenticationCollectionIndicator(value string) {
	c.AuthenticationCollectionIndicator = (*Max35Text)(&value)
}
