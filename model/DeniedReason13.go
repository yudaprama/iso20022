package model

// Specifies the reason why the request or instruction was denied.
type DeniedReason13 struct {

	// Specifies the reason why the request has a denied status.
	Code *DeniedReason18Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (d *DeniedReason13) AddCode() *DeniedReason18Choice {
	d.Code = new(DeniedReason18Choice)
	return d.Code
}

func (d *DeniedReason13) SetAdditionalReasonInformation(value string) {
	d.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
