package model

// Reason for a disabled status.
type DisabledStatusReason1 struct {

	// Reason for the disabled account status.
	Code *DisabledStatusReason2Choice `xml:"Cd"`

	// Additional information about the reason for the disabled account status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (d *DisabledStatusReason1) AddCode() *DisabledStatusReason2Choice {
	d.Code = new(DisabledStatusReason2Choice)
	return d.Code
}

func (d *DisabledStatusReason1) SetAdditionalInformation(value string) {
	d.AdditionalInformation = (*Max350Text)(&value)
}
