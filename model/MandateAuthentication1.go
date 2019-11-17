package model

// Specifies the transport authentication details related to the mandate.
type MandateAuthentication1 struct {

	// Specifies a piece of information used to authenticate a message, that is  to confirm that the message came from the stated sender (its authenticity) and has not been changed in transit (its integrity).
	MessageAuthenticationCode *Max16Text `xml:"MsgAuthntcnCd,omitempty"`

	// Date when the authentication was conducted.
	Date *ISODate `xml:"Dt,omitempty"`

	// Channel used to transmit the authentication information.
	Channel *AuthenticationChannel1Choice `xml:"Chanl,omitempty"`
}

func (m *MandateAuthentication1) SetMessageAuthenticationCode(value string) {
	m.MessageAuthenticationCode = (*Max16Text)(&value)
}

func (m *MandateAuthentication1) SetDate(value string) {
	m.Date = (*ISODate)(&value)
}

func (m *MandateAuthentication1) AddChannel() *AuthenticationChannel1Choice {
	m.Channel = new(AuthenticationChannel1Choice)
	return m.Channel
}
