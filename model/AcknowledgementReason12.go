package model

// Specifies additional information about the processed instruction.
type AcknowledgementReason12 struct {

	// Reason provided for the status.
	Code *AcknowledgementReason15Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcknowledgementReason12) AddCode() *AcknowledgementReason15Choice {
	a.Code = new(AcknowledgementReason15Choice)
	return a.Code
}

func (a *AcknowledgementReason12) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}
