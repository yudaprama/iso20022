package model

// Specifies additional information about the processed instruction.
type AcknowledgementReason13 struct {

	// Reason provided for the status.
	Code *AcknowledgementReason16Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcknowledgementReason13) AddCode() *AcknowledgementReason16Choice {
	a.Code = new(AcknowledgementReason16Choice)
	return a.Code
}

func (a *AcknowledgementReason13) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
