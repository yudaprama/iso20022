package model

// The status of an instruction, advice or request.
type AcknowledgementReason3 struct {

	// Reason provided for the status.
	Code *AcknowledgementReason4Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcknowledgementReason3) AddCode() *AcknowledgementReason4Choice {
	a.Code = new(AcknowledgementReason4Choice)
	return a.Code
}

func (a *AcknowledgementReason3) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}
