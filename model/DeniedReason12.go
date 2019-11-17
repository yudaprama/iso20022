package model

// Specifies the reason why the request or instruction was denied.
type DeniedReason12 struct {

	// Specifies the reason why the request has a denied status.
	Code *DeniedReason17Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (d *DeniedReason12) AddCode() *DeniedReason17Choice {
	d.Code = new(DeniedReason17Choice)
	return d.Code
}

func (d *DeniedReason12) SetAdditionalReasonInformation(value string) {
	d.AdditionalReasonInformation = (*Max210Text)(&value)
}
