package model

// Specifies additional information about the processed instruction.
type AcknowledgementReason18 struct {

	// Choice of format for the acknowledgement reason.
	Code *AcknowledgementReason21Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcknowledgementReason18) AddCode() *AcknowledgementReason21Choice {
	a.Code = new(AcknowledgementReason21Choice)
	return a.Code
}

func (a *AcknowledgementReason18) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
