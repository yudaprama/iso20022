package model

// Modification reasons.
type ModificationReason5 struct {

	// Specifies the reason why the transaction is modified.
	Code *ModificationReason5Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (m *ModificationReason5) AddCode() *ModificationReason5Choice {
	m.Code = new(ModificationReason5Choice)
	return m.Code
}

func (m *ModificationReason5) SetAdditionalReasonInformation(value string) {
	m.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
