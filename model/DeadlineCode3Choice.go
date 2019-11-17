package model

// Choice between a code or a proprietary code for a deadline code.
type DeadlineCode3Choice struct {

	// Standard code to specify the reference date of a corporate action.
	Code *CorporateActionDeadline1Code `xml:"Cd"`

	// Proprietary identification of the reference date of a corporate action.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (d *DeadlineCode3Choice) SetCode(value string) {
	d.Code = (*CorporateActionDeadline1Code)(&value)
}

func (d *DeadlineCode3Choice) AddProprietary() *GenericIdentification30 {
	d.Proprietary = new(GenericIdentification30)
	return d.Proprietary
}
