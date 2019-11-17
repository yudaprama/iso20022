package model

// Modification reasons.
type ModificationReason2 struct {

	// Specifies the reason why the transaction is modified.
	Code *ModificationReason2Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (m *ModificationReason2) AddCode() *ModificationReason2Choice {
	m.Code = new(ModificationReason2Choice)
	return m.Code
}

func (m *ModificationReason2) SetAdditionalReasonInformation(value string) {
	m.AdditionalReasonInformation = (*Max210Text)(&value)
}
