package model

// Specifies additional information about the processed instruction.
type AcknowledgementReason9 struct {

	// Reason provided for the status.
	Code *AcknowledgementReason12Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcknowledgementReason9) AddCode() *AcknowledgementReason12Choice {
	a.Code = new(AcknowledgementReason12Choice)
	return a.Code
}

func (a *AcknowledgementReason9) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}
