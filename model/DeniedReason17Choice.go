package model

// Choice of format for the denied reason.
type DeniedReason17Choice struct {

	// Specifies additional information about the processed instruction.
	Code *DeniedReason3Code `xml:"Cd"`

	// Specifies additional information about the processed instruction.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (d *DeniedReason17Choice) SetCode(value string) {
	d.Code = (*DeniedReason3Code)(&value)
}

func (d *DeniedReason17Choice) AddProprietary() *GenericIdentification30 {
	d.Proprietary = new(GenericIdentification30)
	return d.Proprietary
}
