package model

// Specifies the role of the issuer agent.
type IssuerAgent2 struct {

	// Identifies issuer agent.
	Identification *PartyIdentification40Choice `xml:"Id"`

	// Specifies the role of the issuer agent.
	Role *AgentRole1Code `xml:"Role,omitempty"`
}

func (i *IssuerAgent2) AddIdentification() *PartyIdentification40Choice {
	i.Identification = new(PartyIdentification40Choice)
	return i.Identification
}

func (i *IssuerAgent2) SetRole(value string) {
	i.Role = (*AgentRole1Code)(&value)
}
