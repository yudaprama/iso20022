package model

// Provides information about the corporate action option.
type CorporateActionOption41 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption13Choice `xml:"OptnTp"`

	// Party that owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Account on which a securities entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Identifies the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`

	// Total balance of securities eligible for this corporate action event. The entitlement calculation is based on this balance.
	TotalEligibleBalance *SignedQuantityFormat1 `xml:"TtlElgblBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *SignedQuantityFormat1 `xml:"InstdBal,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *SignedQuantityFormat1 `xml:"UinstdBal,omitempty"`

	// Specifies whether the quantity of financial instrument is a status quantity or a quantity to receive.
	StatusQuantityOrQuantityToReceive *StatusOrQuantityToReceive1Choice `xml:"StsQtyOrQtyToRcv,omitempty"`
}

func (c *CorporateActionOption41) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption41) AddOptionType() *CorporateActionOption13Choice {
	c.OptionType = new(CorporateActionOption13Choice)
	return c.OptionType
}

func (c *CorporateActionOption41) AddAccountOwner() *PartyIdentification36Choice {
	c.AccountOwner = new(PartyIdentification36Choice)
	return c.AccountOwner
}

func (c *CorporateActionOption41) SetSafekeepingAccount(value string) {
	c.SafekeepingAccount = (*Max35Text)(&value)
}

func (c *CorporateActionOption41) AddCashAccount() *CashAccountIdentification5Choice {
	c.CashAccount = new(CashAccountIdentification5Choice)
	return c.CashAccount
}

func (c *CorporateActionOption41) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	c.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return c.SafekeepingPlace
}

func (c *CorporateActionOption41) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return c.FinancialInstrumentIdentification
}

func (c *CorporateActionOption41) AddTotalEligibleBalance() *SignedQuantityFormat1 {
	c.TotalEligibleBalance = new(SignedQuantityFormat1)
	return c.TotalEligibleBalance
}

func (c *CorporateActionOption41) AddInstructedBalance() *SignedQuantityFormat1 {
	c.InstructedBalance = new(SignedQuantityFormat1)
	return c.InstructedBalance
}

func (c *CorporateActionOption41) AddUninstructedBalance() *SignedQuantityFormat1 {
	c.UninstructedBalance = new(SignedQuantityFormat1)
	return c.UninstructedBalance
}

func (c *CorporateActionOption41) AddStatusQuantityOrQuantityToReceive() *StatusOrQuantityToReceive1Choice {
	c.StatusQuantityOrQuantityToReceive = new(StatusOrQuantityToReceive1Choice)
	return c.StatusQuantityOrQuantityToReceive
}
