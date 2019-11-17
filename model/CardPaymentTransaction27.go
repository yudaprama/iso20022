package model

// Cancellation response from the acquirer.
type CardPaymentTransaction27 struct {

	// Outcome of the authorisation, and actions to perform.
	AuthorisationResult *AuthorisationResult3 `xml:"AuthstnRslt"`

	// Set of actions to be performed by the POI (Point Of Interaction) system.
	Action []*Action1 `xml:"Actn,omitempty"`
}

func (c *CardPaymentTransaction27) AddAuthorisationResult() *AuthorisationResult3 {
	c.AuthorisationResult = new(AuthorisationResult3)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction27) AddAction() *Action1 {
	newValue := new(Action1)
	c.Action = append(c.Action, newValue)
	return newValue
}
