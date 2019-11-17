package model

// Choice of formats for an in repair status.
type InRepairStatusReason5Choice struct {

	// In repair reason expressed as a code.
	Code *InRepairStatusReason1Code `xml:"Cd"`

	// In repair reason expressed as a proprietary code.
	Proprietary *GenericIdentification1 `xml:"Prtry"`
}

func (i *InRepairStatusReason5Choice) SetCode(value string) {
	i.Code = (*InRepairStatusReason1Code)(&value)
}

func (i *InRepairStatusReason5Choice) AddProprietary() *GenericIdentification1 {
	i.Proprietary = new(GenericIdentification1)
	return i.Proprietary
}
