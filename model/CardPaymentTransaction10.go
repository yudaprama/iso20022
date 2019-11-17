package model

// Cancellation response from the acquirer.
type CardPaymentTransaction10 struct {

	// Outcome of the authorisation, and actions to perform.
	AuthorisationResult *AuthorisationResult1 `xml:"AuthstnRslt"`

	// Set of actions to be performed by the POI (Point Of Interaction) system.
	Action []*Action1 `xml:"Actn,omitempty"`
}

func (c *CardPaymentTransaction10) AddAuthorisationResult() *AuthorisationResult1 {
	c.AuthorisationResult = new(AuthorisationResult1)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction10) AddAction() *Action1 {
	newValue := new(Action1)
	c.Action = append(c.Action, newValue)
	return newValue
}
