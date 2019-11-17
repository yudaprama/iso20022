package model

// Choice of formats to express the role of the agent.
type AgentRole1FormatChoice struct {

	// Standard code to specify the role of the agent.
	Code *AgentRole2Code `xml:"Cd"`

	// Proprietary code to express the role of the agent.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (a *AgentRole1FormatChoice) SetCode(value string) {
	a.Code = (*AgentRole2Code)(&value)
}

func (a *AgentRole1FormatChoice) AddProprietary() *GenericIdentification13 {
	a.Proprietary = new(GenericIdentification13)
	return a.Proprietary
}
