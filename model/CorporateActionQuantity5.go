package model

// Specifies corporate action quantities.
type CorporateActionQuantity5 struct {

	// Minimum quantity (or lot) of financial instrument that may be exercised or tendered.
	MinimumExercisableQuantity *FinancialInstrumentQuantity19Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity (or lot) of financial instrument that  may be exercised or tendered.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity20Choice `xml:"MinExrcblMltplQty,omitempty"`

	// The maximum number of securities the offeror/issuer is ready to purchase or redeem. This can be a number or the term "any and all".
	MaximumQuantity *FinancialInstrumentQuantity19Choice `xml:"MaxQty,omitempty"`

	// Minimum quantity of securities the offeror/issuer is ready to purchase or redeem under the terms of the event. This can be a number or the term "any and all".
	MinimumQuantitySought *FinancialInstrumentQuantity19Choice `xml:"MinQtySght,omitempty"`

	// Quantity of equity that makes up the new board lot.
	NewBoardLotQuantity *FinancialInstrumentQuantity20Choice `xml:"NewBrdLotQty,omitempty"`

	// New denomination of the equity following, for example, an increase or decrease in nominal value.
	NewDenominationQuantity *FinancialInstrumentQuantity20Choice `xml:"NewDnmtnQty,omitempty"`

	// Minimum integral amount of securities that each account owner must have remaining after the called amounts are applied.
	BaseDenomination *FinancialInstrumentQuantity20Choice `xml:"BaseDnmtn,omitempty"`

	// Amount used when the called amount is not met by running the lottery with the base denomination.
	IncrementalDenomination *FinancialInstrumentQuantity20Choice `xml:"IncrmtlDnmtn,omitempty"`
}

func (c *CorporateActionQuantity5) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity19Choice {
	c.MinimumExercisableQuantity = new(FinancialInstrumentQuantity19Choice)
	return c.MinimumExercisableQuantity
}

func (c *CorporateActionQuantity5) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity20Choice {
	c.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity20Choice)
	return c.MinimumExercisableMultipleQuantity
}

func (c *CorporateActionQuantity5) AddMaximumQuantity() *FinancialInstrumentQuantity19Choice {
	c.MaximumQuantity = new(FinancialInstrumentQuantity19Choice)
	return c.MaximumQuantity
}

func (c *CorporateActionQuantity5) AddMinimumQuantitySought() *FinancialInstrumentQuantity19Choice {
	c.MinimumQuantitySought = new(FinancialInstrumentQuantity19Choice)
	return c.MinimumQuantitySought
}

func (c *CorporateActionQuantity5) AddNewBoardLotQuantity() *FinancialInstrumentQuantity20Choice {
	c.NewBoardLotQuantity = new(FinancialInstrumentQuantity20Choice)
	return c.NewBoardLotQuantity
}

func (c *CorporateActionQuantity5) AddNewDenominationQuantity() *FinancialInstrumentQuantity20Choice {
	c.NewDenominationQuantity = new(FinancialInstrumentQuantity20Choice)
	return c.NewDenominationQuantity
}

func (c *CorporateActionQuantity5) AddBaseDenomination() *FinancialInstrumentQuantity20Choice {
	c.BaseDenomination = new(FinancialInstrumentQuantity20Choice)
	return c.BaseDenomination
}

func (c *CorporateActionQuantity5) AddIncrementalDenomination() *FinancialInstrumentQuantity20Choice {
	c.IncrementalDenomination = new(FinancialInstrumentQuantity20Choice)
	return c.IncrementalDenomination
}
