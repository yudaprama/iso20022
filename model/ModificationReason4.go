package model

// Modification reasons.
type ModificationReason4 struct {

	// Specifies the reason why the transaction is modified.
	Code *ModificationReason4Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (m *ModificationReason4) AddCode() *ModificationReason4Choice {
	m.Code = new(ModificationReason4Choice)
	return m.Code
}

func (m *ModificationReason4) SetAdditionalReasonInformation(value string) {
	m.AdditionalReasonInformation = (*Max210Text)(&value)
}
