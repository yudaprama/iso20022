package model

// Specifies additional information about the processed instruction.
type AcknowledgementReason11 struct {

	// Choice of format for the acknowledgement reason.
	Code *AcknowledgementReason14Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcknowledgementReason11) AddCode() *AcknowledgementReason14Choice {
	a.Code = new(AcknowledgementReason14Choice)
	return a.Code
}

func (a *AcknowledgementReason11) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}
