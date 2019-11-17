package model

// The status of an instruction, advice or request.
type AcknowledgementReason7 struct {

	// Choice of format for the acknowledgement reason.
	Code *AcknowledgementReason9Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcknowledgementReason7) AddCode() *AcknowledgementReason9Choice {
	a.Code = new(AcknowledgementReason9Choice)
	return a.Code
}

func (a *AcknowledgementReason7) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}
