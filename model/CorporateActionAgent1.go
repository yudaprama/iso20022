package model

// Provides information about the agent.
type CorporateActionAgent1 struct {

	// Identification of the agent.
	AgentIdentification *PartyIdentification2Choice `xml:"AgtId"`

	// Role played by the agent.
	AgentRole *AgentRole1FormatChoice `xml:"AgtRole"`

	// Contact person at the agent.
	ContactPerson *NameAndAddress5 `xml:"CtctPrsn,omitempty"`
}

func (c *CorporateActionAgent1) AddAgentIdentification() *PartyIdentification2Choice {
	c.AgentIdentification = new(PartyIdentification2Choice)
	return c.AgentIdentification
}

func (c *CorporateActionAgent1) AddAgentRole() *AgentRole1FormatChoice {
	c.AgentRole = new(AgentRole1FormatChoice)
	return c.AgentRole
}

func (c *CorporateActionAgent1) AddContactPerson() *NameAndAddress5 {
	c.ContactPerson = new(NameAndAddress5)
	return c.ContactPerson
}
