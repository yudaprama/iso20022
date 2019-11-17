package model

// Specifies the reason why the request or instruction was denied.
type DeniedReason16 struct {

	// Specifies the reason why the request has a denied status.
	Code *DeniedReason21Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (d *DeniedReason16) AddCode() *DeniedReason21Choice {
	d.Code = new(DeniedReason21Choice)
	return d.Code
}

func (d *DeniedReason16) SetAdditionalReasonInformation(value string) {
	d.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
