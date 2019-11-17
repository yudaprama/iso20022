package model

// Cancellation response from the acquirer.
type CardPaymentTransaction58 struct {

	// Outcome of the authorisation, and actions to perform.
	AuthorisationResult *AuthorisationResult12 `xml:"AuthstnRslt"`

	// Set of actions to be performed by the POI (Point Of Interaction) system.
	Action []*Action6 `xml:"Actn,omitempty"`
}

func (c *CardPaymentTransaction58) AddAuthorisationResult() *AuthorisationResult12 {
	c.AuthorisationResult = new(AuthorisationResult12)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction58) AddAction() *Action6 {
	newValue := new(Action6)
	c.Action = append(c.Action, newValue)
	return newValue
}
