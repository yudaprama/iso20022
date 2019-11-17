package model

// Choice between a standard code or proprietary code to specify the eligibility of outturn resources.
type NonEligibleProceedsIndicator1Choice struct {

	// Standard code to specify information regarding the non eligibility of the outturn resources.
	Code *NonEligibleProceedsIndicator1Code `xml:"Cd"`

	// Proprietary identification to specify information regarding the non eligibility of the outturn resources.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (n *NonEligibleProceedsIndicator1Choice) SetCode(value string) {
	n.Code = (*NonEligibleProceedsIndicator1Code)(&value)
}

func (n *NonEligibleProceedsIndicator1Choice) AddProprietary() *GenericIdentification20 {
	n.Proprietary = new(GenericIdentification20)
	return n.Proprietary
}
