package model

// Choice of format for the denied reason.
type DeniedReason18Choice struct {

	// Specifies additional information about the processed instruction.
	Code *DeniedReason3Code `xml:"Cd"`

	// Specifies additional information about the processed instruction.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (d *DeniedReason18Choice) SetCode(value string) {
	d.Code = (*DeniedReason3Code)(&value)
}

func (d *DeniedReason18Choice) AddProprietary() *GenericIdentification47 {
	d.Proprietary = new(GenericIdentification47)
	return d.Proprietary
}
