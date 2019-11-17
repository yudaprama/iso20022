package model

// Provides information about the failure of the settlement of a movement.
type CorporateActionMovementFailedStatus1 struct {

	// Identification of the agent account.
	AgentAccountIdentification *Max35Text `xml:"AgtAcctId"`

	// Identification of the client account.
	ClientAccountIdentification *Max35Text `xml:"ClntAcctId,omitempty"`

	// Identification of the party that owns the client account.
	AccountOwnerIdentification *PartyIdentification2Choice `xml:"AcctOwnrId,omitempty"`

	// Provides information about the resource movement that failed and the reason of the failure.
	ResourceDetails []*FailedMovement1 `xml:"RsrcDtls"`
}

func (c *CorporateActionMovementFailedStatus1) SetAgentAccountIdentification(value string) {
	c.AgentAccountIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionMovementFailedStatus1) SetClientAccountIdentification(value string) {
	c.ClientAccountIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionMovementFailedStatus1) AddAccountOwnerIdentification() *PartyIdentification2Choice {
	c.AccountOwnerIdentification = new(PartyIdentification2Choice)
	return c.AccountOwnerIdentification
}

func (c *CorporateActionMovementFailedStatus1) AddResourceDetails() *FailedMovement1 {
	newValue := new(FailedMovement1)
	c.ResourceDetails = append(c.ResourceDetails, newValue)
	return newValue
}
