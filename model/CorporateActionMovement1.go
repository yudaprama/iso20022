package model

// Provides information about the movement instruction.
type CorporateActionMovement1 struct {

	// Type of movement instruction.
	OrderType *DistributionInstructionType1Code `xml:"OrdrTp"`

	// Indicates whether the movement is a high priority or not.
	// Meaning when true: High priority
	// Meaning when false: Standard
	HighPriorityIndicator *YesNoIndicator `xml:"HghPrtyInd"`

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb,omitempty"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption1FormatChoice `xml:"OptnTp,omitempty"`

	// Date at which the distribution movement must be executed.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt"`

	// Identification of the party that owns the account.
	AccountOwnerIdentification *PartyIdentification2Choice `xml:"AcctOwnrId,omitempty"`

	// Identification of the account.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Quantity of securities in the confirmed balance, ie, the balance to which the credit of the outturned resources applies.
	ConfirmedBalanceSecuritiesQuantity *UnitOrFaceAmount1Choice `xml:"ConfdBalSctiesQty,omitempty"`
}

func (c *CorporateActionMovement1) SetOrderType(value string) {
	c.OrderType = (*DistributionInstructionType1Code)(&value)
}

func (c *CorporateActionMovement1) SetHighPriorityIndicator(value string) {
	c.HighPriorityIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionMovement1) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionMovement1) AddOptionType() *CorporateActionOption1FormatChoice {
	c.OptionType = new(CorporateActionOption1FormatChoice)
	return c.OptionType
}

func (c *CorporateActionMovement1) SetRequestedExecutionDate(value string) {
	c.RequestedExecutionDate = (*ISODate)(&value)
}

func (c *CorporateActionMovement1) AddAccountOwnerIdentification() *PartyIdentification2Choice {
	c.AccountOwnerIdentification = new(PartyIdentification2Choice)
	return c.AccountOwnerIdentification
}

func (c *CorporateActionMovement1) SetAccountIdentification(value string) {
	c.AccountIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionMovement1) AddConfirmedBalanceSecuritiesQuantity() *UnitOrFaceAmount1Choice {
	c.ConfirmedBalanceSecuritiesQuantity = new(UnitOrFaceAmount1Choice)
	return c.ConfirmedBalanceSecuritiesQuantity
}
