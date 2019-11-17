package model

// Specifies additional information about the processed instruction.
type AcknowledgementReason19 struct {

	// Reason provided for the status.
	Code *AcknowledgementReason22Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcknowledgementReason19) AddCode() *AcknowledgementReason22Choice {
	a.Code = new(AcknowledgementReason22Choice)
	return a.Code
}

func (a *AcknowledgementReason19) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
