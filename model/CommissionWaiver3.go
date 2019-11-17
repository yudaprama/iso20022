package model

// Non-enforcement of the right to all or part of a commission by the party entitled to the commission.
type CommissionWaiver3 struct {

	// Form of the rebate, eg, cash.
	InstructionBasis *WaivingInstruction1Code `xml:"InstrBsis"`

	// Form of the rebate, eg, cash.
	ExtendedInstructionBasis *Extended350Code `xml:"XtndedInstrBsis"`

	// Proportion of the commission that is waived, ie, if  the commission is 5% and half is waived, 2.5% should be stated in this field.
	WaivedRate *PercentageRate `xml:"WvdRate"`
}

func (c *CommissionWaiver3) SetInstructionBasis(value string) {
	c.InstructionBasis = (*WaivingInstruction1Code)(&value)
}

func (c *CommissionWaiver3) SetExtendedInstructionBasis(value string) {
	c.ExtendedInstructionBasis = (*Extended350Code)(&value)
}

func (c *CommissionWaiver3) SetWaivedRate(value string) {
	c.WaivedRate = (*PercentageRate)(&value)
}
