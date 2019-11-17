package model

// The status of an instruction, advice or request.
type AcknowledgementReason1 struct {

	// Reason provided for the status.
	Code *AcknowledgementReason1Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcknowledgementReason1) AddCode() *AcknowledgementReason1Choice {
	a.Code = new(AcknowledgementReason1Choice)
	return a.Code
}

func (a *AcknowledgementReason1) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}
