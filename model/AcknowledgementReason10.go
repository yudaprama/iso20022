package model

// Specifies additional information about the processed instruction.
type AcknowledgementReason10 struct {

	// Specifies additional information about the processed instruction.
	Code *AcknowledgementReason13Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcknowledgementReason10) AddCode() *AcknowledgementReason13Choice {
	a.Code = new(AcknowledgementReason13Choice)
	return a.Code
}

func (a *AcknowledgementReason10) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}
