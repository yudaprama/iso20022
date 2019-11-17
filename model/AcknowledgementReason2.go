package model

// The status of an instruction, advice or request.
type AcknowledgementReason2 struct {

	// Specifies additional information about the processed instruction.
	Code *AcknowledgementReason3Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcknowledgementReason2) AddCode() *AcknowledgementReason3Choice {
	a.Code = new(AcknowledgementReason3Choice)
	return a.Code
}

func (a *AcknowledgementReason2) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}
