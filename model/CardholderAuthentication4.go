package model

// Data related to the authentication of the cardholder.
type CardholderAuthentication4 struct {

	// Method used to authenticate the cardholder.
	AuthenticationMethod *AuthenticationMethod2Code `xml:"AuthntcnMtd"`

	// Entity or object in charge of verifying the cardholder authenticity.
	AuthenticationEntity *AuthenticationEntity1Code `xml:"AuthntcnNtty"`
}

func (c *CardholderAuthentication4) SetAuthenticationMethod(value string) {
	c.AuthenticationMethod = (*AuthenticationMethod2Code)(&value)
}

func (c *CardholderAuthentication4) SetAuthenticationEntity(value string) {
	c.AuthenticationEntity = (*AuthenticationEntity1Code)(&value)
}
