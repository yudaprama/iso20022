package model

// Choice of format for the denied reason.
type DeniedReason3Choice struct {

	// Specifies additional information about the processed instruction.
	Code *DeniedReason3Code `xml:"Cd"`

	// Specifies additional information about the processed instruction.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (d *DeniedReason3Choice) SetCode(value string) {
	d.Code = (*DeniedReason3Code)(&value)
}

func (d *DeniedReason3Choice) AddProprietary() *GenericIdentification20 {
	d.Proprietary = new(GenericIdentification20)
	return d.Proprietary
}
