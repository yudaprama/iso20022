package model

// The status of an instruction, advice or request.
type DeniedReason5 struct {

	// Specifies the reason why the request has a denied status.
	Code *DeniedReason7Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (d *DeniedReason5) AddCode() *DeniedReason7Choice {
	d.Code = new(DeniedReason7Choice)
	return d.Code
}

func (d *DeniedReason5) SetAdditionalReasonInformation(value string) {
	d.AdditionalReasonInformation = (*Max210Text)(&value)
}
