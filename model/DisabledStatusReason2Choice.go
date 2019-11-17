package model

// Choice of formats for a disabled reason code.
type DisabledStatusReason2Choice struct {

	// Reason for the disabled account status expressed as a code.
	Code *DisabledReason2Code `xml:"Cd"`

	// Reason for the disabled account status expressed as a proprietary code.
	Proprietary *GenericIdentification36 `xml:"Prtry"`
}

func (d *DisabledStatusReason2Choice) SetCode(value string) {
	d.Code = (*DisabledReason2Code)(&value)
}

func (d *DisabledStatusReason2Choice) AddProprietary() *GenericIdentification36 {
	d.Proprietary = new(GenericIdentification36)
	return d.Proprietary
}
