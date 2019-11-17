package model

// Data related to the authentication of the cardholder.
type CardholderAuthentication2 struct {

	// Method used to authenticate the cardholder.
	AuthenticationMethod *AuthenticationMethod1Code `xml:"AuthntcnMtd"`

	// Entity or object in charge of verifying the cardholder authenticity.
	AuthenticationEntity *AuthenticationEntity1Code `xml:"AuthntcnNtty"`
}

func (c *CardholderAuthentication2) SetAuthenticationMethod(value string) {
	c.AuthenticationMethod = (*AuthenticationMethod1Code)(&value)
}

func (c *CardholderAuthentication2) SetAuthenticationEntity(value string) {
	c.AuthenticationEntity = (*AuthenticationEntity1Code)(&value)
}
