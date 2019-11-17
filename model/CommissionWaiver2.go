package model

// Non-enforcement of the right to all or part of a commission by the party entitled to the commission.
type CommissionWaiver2 struct {

	// Form of the rebate, eg, cash.
	InstructionBasis *WaivingType1 `xml:"InstrBsis"`

	// Proportion of the commission that is waived, ie, if  the commission is 5% and half is waived, 2.5% should be stated in this field.
	WaivedRate *PercentageRate `xml:"WvdRate"`
}

func (c *CommissionWaiver2) AddInstructionBasis() *WaivingType1 {
	c.InstructionBasis = new(WaivingType1)
	return c.InstructionBasis
}

func (c *CommissionWaiver2) SetWaivedRate(value string) {
	c.WaivedRate = (*PercentageRate)(&value)
}
