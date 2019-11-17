package model

// The status of an instruction, advice or request.
type DeniedReason2 struct {

	// Specifies the reason why the request has a denied status.
	Code *DeniedReason2Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (d *DeniedReason2) AddCode() *DeniedReason2Choice {
	d.Code = new(DeniedReason2Choice)
	return d.Code
}

func (d *DeniedReason2) SetAdditionalReasonInformation(value string) {
	d.AdditionalReasonInformation = (*Max210Text)(&value)
}
