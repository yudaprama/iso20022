package model

// Choice between a standard code or proprietary code to specify the eligibility of outturn resources.
type NonEligibleProceedsIndicator4Choice struct {

	// Standard code to specify information regarding the non eligibility of the outturn resources.
	Code *NonEligibleProceedsIndicator1Code `xml:"Cd"`

	// Proprietary identification to specify information regarding the non eligibility of the outturn resources.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (n *NonEligibleProceedsIndicator4Choice) SetCode(value string) {
	n.Code = (*NonEligibleProceedsIndicator1Code)(&value)
}

func (n *NonEligibleProceedsIndicator4Choice) AddProprietary() *GenericIdentification47 {
	n.Proprietary = new(GenericIdentification47)
	return n.Proprietary
}
