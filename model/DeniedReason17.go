package model

// The status of an instruction, advice or request.
type DeniedReason17 struct {

	// Specifies the reason why the request has a denied status.
	Code *DeniedReason24Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (d *DeniedReason17) AddCode() *DeniedReason24Choice {
	d.Code = new(DeniedReason24Choice)
	return d.Code
}

func (d *DeniedReason17) SetAdditionalReasonInformation(value string) {
	d.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
