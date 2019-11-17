package model

// Specifies additional information about the processed instruction.
type AcknowledgementReason15 struct {

	// Specifies additional information about the processed instruction.
	Code *AcknowledgementReason18Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcknowledgementReason15) AddCode() *AcknowledgementReason18Choice {
	a.Code = new(AcknowledgementReason18Choice)
	return a.Code
}

func (a *AcknowledgementReason15) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
