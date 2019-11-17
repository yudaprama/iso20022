package model

// Specifies corporate action quantities.
type CorporateActionQuantity8 struct {

	// The maximum number of securities the offeror/issuer is ready to purchase or redeem. This can be a number or the term "any and all".
	MaximumQuantity *FinancialInstrumentQuantity21Choice `xml:"MaxQty,omitempty"`

	// Minimum quantity of securities the offeror/issuer is ready to purchase or redeem under the terms of the event. This can be a number or the term "any and all".
	MinimumQuantitySought *FinancialInstrumentQuantity21Choice `xml:"MinQtySght,omitempty"`

	// Quantity of equity that makes up the new board lot.
	NewBoardLotQuantity *FinancialInstrumentQuantity22Choice `xml:"NewBrdLotQty,omitempty"`

	// New denomination of the equity following, for example, an increase or decrease in nominal value.
	NewDenominationQuantity *FinancialInstrumentQuantity22Choice `xml:"NewDnmtnQty,omitempty"`

	// Minimum integral amount of securities that each account owner must have remaining after the called amounts are applied.
	BaseDenomination *FinancialInstrumentQuantity22Choice `xml:"BaseDnmtn,omitempty"`

	// Amount used when the called amount is not met by running the lottery with the base denomination.
	IncrementalDenomination *FinancialInstrumentQuantity22Choice `xml:"IncrmtlDnmtn,omitempty"`
}

func (c *CorporateActionQuantity8) AddMaximumQuantity() *FinancialInstrumentQuantity21Choice {
	c.MaximumQuantity = new(FinancialInstrumentQuantity21Choice)
	return c.MaximumQuantity
}

func (c *CorporateActionQuantity8) AddMinimumQuantitySought() *FinancialInstrumentQuantity21Choice {
	c.MinimumQuantitySought = new(FinancialInstrumentQuantity21Choice)
	return c.MinimumQuantitySought
}

func (c *CorporateActionQuantity8) AddNewBoardLotQuantity() *FinancialInstrumentQuantity22Choice {
	c.NewBoardLotQuantity = new(FinancialInstrumentQuantity22Choice)
	return c.NewBoardLotQuantity
}

func (c *CorporateActionQuantity8) AddNewDenominationQuantity() *FinancialInstrumentQuantity22Choice {
	c.NewDenominationQuantity = new(FinancialInstrumentQuantity22Choice)
	return c.NewDenominationQuantity
}

func (c *CorporateActionQuantity8) AddBaseDenomination() *FinancialInstrumentQuantity22Choice {
	c.BaseDenomination = new(FinancialInstrumentQuantity22Choice)
	return c.BaseDenomination
}

func (c *CorporateActionQuantity8) AddIncrementalDenomination() *FinancialInstrumentQuantity22Choice {
	c.IncrementalDenomination = new(FinancialInstrumentQuantity22Choice)
	return c.IncrementalDenomination
}
