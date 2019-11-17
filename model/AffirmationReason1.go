package model

// The status of an instruction, advice or request.
type AffirmationReason1 struct {

	// Specifies the reason why the instruction/request has a unaffirmed status.
	Code *UnaffirmedReason2Choice `xml:"Cd"`

	// Provides additional information about the reason in narrative form.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AffirmationReason1) AddCode() *UnaffirmedReason2Choice {
	a.Code = new(UnaffirmedReason2Choice)
	return a.Code
}

func (a *AffirmationReason1) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}
