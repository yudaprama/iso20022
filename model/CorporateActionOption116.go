package model

// Provides information about the corporate action option.
type CorporateActionOption116 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption21Choice `xml:"OptnTp"`

	// Party that owns the account.
	AccountOwner *PartyIdentification92Choice `xml:"AcctOwnr,omitempty"`

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Account on which a securities entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat8Choice `xml:"SfkpgPlc,omitempty"`

	// Identifies the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId,omitempty"`

	// Total balance of securities eligible for this corporate action event. The entitlement calculation is based on this balance.
	TotalEligibleBalance *SignedQuantityFormat7 `xml:"TtlElgblBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *SignedQuantityFormat7 `xml:"InstdBal,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *SignedQuantityFormat7 `xml:"UinstdBal,omitempty"`

	// Quantity of securities that has been assigned the status indicated.
	StatusQuantity *Quantity6Choice `xml:"StsQty"`
}

func (c *CorporateActionOption116) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption116) AddOptionType() *CorporateActionOption21Choice {
	c.OptionType = new(CorporateActionOption21Choice)
	return c.OptionType
}

func (c *CorporateActionOption116) AddAccountOwner() *PartyIdentification92Choice {
	c.AccountOwner = new(PartyIdentification92Choice)
	return c.AccountOwner
}

func (c *CorporateActionOption116) SetSafekeepingAccount(value string) {
	c.SafekeepingAccount = (*Max35Text)(&value)
}

func (c *CorporateActionOption116) AddCashAccount() *CashAccountIdentification5Choice {
	c.CashAccount = new(CashAccountIdentification5Choice)
	return c.CashAccount
}

func (c *CorporateActionOption116) AddSafekeepingPlace() *SafekeepingPlaceFormat8Choice {
	c.SafekeepingPlace = new(SafekeepingPlaceFormat8Choice)
	return c.SafekeepingPlace
}

func (c *CorporateActionOption116) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return c.FinancialInstrumentIdentification
}

func (c *CorporateActionOption116) AddTotalEligibleBalance() *SignedQuantityFormat7 {
	c.TotalEligibleBalance = new(SignedQuantityFormat7)
	return c.TotalEligibleBalance
}

func (c *CorporateActionOption116) AddInstructedBalance() *SignedQuantityFormat7 {
	c.InstructedBalance = new(SignedQuantityFormat7)
	return c.InstructedBalance
}

func (c *CorporateActionOption116) AddUninstructedBalance() *SignedQuantityFormat7 {
	c.UninstructedBalance = new(SignedQuantityFormat7)
	return c.UninstructedBalance
}

func (c *CorporateActionOption116) AddStatusQuantity() *Quantity6Choice {
	c.StatusQuantity = new(Quantity6Choice)
	return c.StatusQuantity
}
