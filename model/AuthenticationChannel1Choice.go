package model

// Choice of format for an authentication channel.
type AuthenticationChannel1Choice struct {

	// Authentication channel expressed as an external ISO 20022 code.
	Code *ExternalAuthenticationChannel1Code `xml:"Cd"`

	// Authentication channel expressed as a proprietary code.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (a *AuthenticationChannel1Choice) SetCode(value string) {
	a.Code = (*ExternalAuthenticationChannel1Code)(&value)
}

func (a *AuthenticationChannel1Choice) SetProprietary(value string) {
	a.Proprietary = (*Max35Text)(&value)
}
