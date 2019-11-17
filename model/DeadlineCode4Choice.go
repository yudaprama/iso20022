package model

// Choice between a code or a proprietary code for a deadline code.
type DeadlineCode4Choice struct {

	// Standard code to specify the reference date of a corporate action.
	Code *CorporateActionDeadline1Code `xml:"Cd"`

	// Proprietary identification of the reference date of a corporate action.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (d *DeadlineCode4Choice) SetCode(value string) {
	d.Code = (*CorporateActionDeadline1Code)(&value)
}

func (d *DeadlineCode4Choice) AddProprietary() *GenericIdentification47 {
	d.Proprietary = new(GenericIdentification47)
	return d.Proprietary
}
