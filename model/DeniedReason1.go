package model

// The status of an instruction, advice or request.
type DeniedReason1 struct {

	// Specifies the reason why the request has a denied status.
	Code *DeniedReason3Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (d *DeniedReason1) AddCode() *DeniedReason3Choice {
	d.Code = new(DeniedReason3Choice)
	return d.Code
}

func (d *DeniedReason1) SetAdditionalReasonInformation(value string) {
	d.AdditionalReasonInformation = (*Max210Text)(&value)
}
