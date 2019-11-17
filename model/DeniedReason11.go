package model

// Specifies the reason why the request or instruction was denied.
type DeniedReason11 struct {

	// Specifies the reason why the request has a denied status.
	Code *DeniedReason16Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (d *DeniedReason11) AddCode() *DeniedReason16Choice {
	d.Code = new(DeniedReason16Choice)
	return d.Code
}

func (d *DeniedReason11) SetAdditionalReasonInformation(value string) {
	d.AdditionalReasonInformation = (*Max210Text)(&value)
}
