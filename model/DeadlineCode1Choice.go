package model

// Choice between a code or a proprietary code for a deadline code.
type DeadlineCode1Choice struct {

	// Standard code to specify the reference date of a corporate action.
	Code *CorporateActionDeadline1Code `xml:"Cd"`

	// Proprietary identification of the reference date of a corporate action.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (d *DeadlineCode1Choice) SetCode(value string) {
	d.Code = (*CorporateActionDeadline1Code)(&value)
}

func (d *DeadlineCode1Choice) AddProprietary() *GenericIdentification20 {
	d.Proprietary = new(GenericIdentification20)
	return d.Proprietary
}
