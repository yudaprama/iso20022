package model

// Choice between a standard code or proprietary code to specify the eligibility of outturn resources.
type NonEligibleProceedsIndicator3Choice struct {

	// Standard code to specify information regarding the non eligibility of the outturn resources.
	Code *NonEligibleProceedsIndicator1Code `xml:"Cd"`

	// Proprietary identification to specify information regarding the non eligibility of the outturn resources.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (n *NonEligibleProceedsIndicator3Choice) SetCode(value string) {
	n.Code = (*NonEligibleProceedsIndicator1Code)(&value)
}

func (n *NonEligibleProceedsIndicator3Choice) AddProprietary() *GenericIdentification30 {
	n.Proprietary = new(GenericIdentification30)
	return n.Proprietary
}
