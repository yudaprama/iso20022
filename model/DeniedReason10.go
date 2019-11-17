package model

// Specifies the reason why the request or instruction was denied.
type DeniedReason10 struct {

	// Specifies the reason why the request has a denied status.
	Code *DeniedReason15Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (d *DeniedReason10) AddCode() *DeniedReason15Choice {
	d.Code = new(DeniedReason15Choice)
	return d.Code
}

func (d *DeniedReason10) SetAdditionalReasonInformation(value string) {
	d.AdditionalReasonInformation = (*Max210Text)(&value)
}
