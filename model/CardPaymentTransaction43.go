package model

// Cancellation response from the acquirer.
type CardPaymentTransaction43 struct {

	// Outcome of the authorisation, and actions to perform.
	AuthorisationResult *AuthorisationResult6 `xml:"AuthstnRslt"`

	// Set of actions to be performed by the POI (Point Of Interaction) system.
	Action []*Action3 `xml:"Actn,omitempty"`
}

func (c *CardPaymentTransaction43) AddAuthorisationResult() *AuthorisationResult6 {
	c.AuthorisationResult = new(AuthorisationResult6)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction43) AddAction() *Action3 {
	newValue := new(Action3)
	c.Action = append(c.Action, newValue)
	return newValue
}
