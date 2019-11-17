package model

// Non-enforcement of the right to all or part of a commission by the party entitled to the commission.
type CommissionWaiver4 struct {

	// Form of the rebate, for example, cash.
	InstructionBasis *WaivingInstruction1Choice `xml:"InstrBsis"`

	// Proportion of the commission that is waived, for example, if  the commission is 5% and half is waived, 2.5% should be stated in this field.
	WaivedRate *PercentageRate `xml:"WvdRate"`
}

func (c *CommissionWaiver4) AddInstructionBasis() *WaivingInstruction1Choice {
	c.InstructionBasis = new(WaivingInstruction1Choice)
	return c.InstructionBasis
}

func (c *CommissionWaiver4) SetWaivedRate(value string) {
	c.WaivedRate = (*PercentageRate)(&value)
}
