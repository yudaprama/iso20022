package model

// Specifies corporate action quantities.
type CorporateActionQuantity3 struct {

	// Minimum quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Maximum number of securities the offeror is requesting to complete the event.
	MaximumQuantity *FinancialInstrumentQuantity16Choice `xml:"MaxQty,omitempty"`

	// Minimum quantity of securities the offeror/issuer will purchase or redeem under the terms of the event. This can be a number or the term "any and all".
	MinimumQuantitySought *FinancialInstrumentQuantity16Choice `xml:"MinQtySght,omitempty"`

	// Quantity of equity that makes up the new board lot.
	NewBoardLotQuantity *FinancialInstrumentQuantity1Choice `xml:"NewBrdLotQty,omitempty"`

	// New denomination of the equity following, for example, an increase or decrease in nominal value.
	NewDenominationQuantity *FinancialInstrumentQuantity1Choice `xml:"NewDnmtnQty,omitempty"`

	// Minimum integral amount of securities that each account owner must have remaining after the called amounts are applied.
	BaseDenomination *FinancialInstrumentQuantity1Choice `xml:"BaseDnmtn,omitempty"`

	// Amount used when the called amount is not met by running the lottery with the base denomination.
	IncrementalDenomination *FinancialInstrumentQuantity1Choice `xml:"IncrmtlDnmtn,omitempty"`
}

func (c *CorporateActionQuantity3) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	c.MinimumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return c.MinimumExercisableQuantity
}

func (c *CorporateActionQuantity3) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity1Choice {
	c.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity1Choice)
	return c.MinimumExercisableMultipleQuantity
}

func (c *CorporateActionQuantity3) AddMaximumQuantity() *FinancialInstrumentQuantity16Choice {
	c.MaximumQuantity = new(FinancialInstrumentQuantity16Choice)
	return c.MaximumQuantity
}

func (c *CorporateActionQuantity3) AddMinimumQuantitySought() *FinancialInstrumentQuantity16Choice {
	c.MinimumQuantitySought = new(FinancialInstrumentQuantity16Choice)
	return c.MinimumQuantitySought
}

func (c *CorporateActionQuantity3) AddNewBoardLotQuantity() *FinancialInstrumentQuantity1Choice {
	c.NewBoardLotQuantity = new(FinancialInstrumentQuantity1Choice)
	return c.NewBoardLotQuantity
}

func (c *CorporateActionQuantity3) AddNewDenominationQuantity() *FinancialInstrumentQuantity1Choice {
	c.NewDenominationQuantity = new(FinancialInstrumentQuantity1Choice)
	return c.NewDenominationQuantity
}

func (c *CorporateActionQuantity3) AddBaseDenomination() *FinancialInstrumentQuantity1Choice {
	c.BaseDenomination = new(FinancialInstrumentQuantity1Choice)
	return c.BaseDenomination
}

func (c *CorporateActionQuantity3) AddIncrementalDenomination() *FinancialInstrumentQuantity1Choice {
	c.IncrementalDenomination = new(FinancialInstrumentQuantity1Choice)
	return c.IncrementalDenomination
}
