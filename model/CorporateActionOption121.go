package model

// Provides information about the corporate action option.
type CorporateActionOption121 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption22Choice `xml:"OptnTp"`

	// Party that owns the account.
	AccountOwner *PartyIdentification103Choice `xml:"AcctOwnr,omitempty"`

	// Account where financial instruments are maintained.
	SafekeepingAccount *RestrictedFINXMax35Text `xml:"SfkpgAcct,omitempty"`

	// Account on which a securities entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat11Choice `xml:"SfkpgPlc,omitempty"`

	// Identifies the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId,omitempty"`

	// Total balance of securities eligible for this corporate action event. The entitlement calculation is based on this balance.
	TotalEligibleBalance *SignedQuantityFormat8 `xml:"TtlElgblBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *SignedQuantityFormat8 `xml:"InstdBal,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *SignedQuantityFormat8 `xml:"UinstdBal,omitempty"`

	// Quantity of securities that has been assigned the status indicated.
	StatusQuantity *Quantity10Choice `xml:"StsQty"`
}

func (c *CorporateActionOption121) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption121) AddOptionType() *CorporateActionOption22Choice {
	c.OptionType = new(CorporateActionOption22Choice)
	return c.OptionType
}

func (c *CorporateActionOption121) AddAccountOwner() *PartyIdentification103Choice {
	c.AccountOwner = new(PartyIdentification103Choice)
	return c.AccountOwner
}

func (c *CorporateActionOption121) SetSafekeepingAccount(value string) {
	c.SafekeepingAccount = (*RestrictedFINXMax35Text)(&value)
}

func (c *CorporateActionOption121) AddCashAccount() *CashAccountIdentification6Choice {
	c.CashAccount = new(CashAccountIdentification6Choice)
	return c.CashAccount
}

func (c *CorporateActionOption121) AddSafekeepingPlace() *SafekeepingPlaceFormat11Choice {
	c.SafekeepingPlace = new(SafekeepingPlaceFormat11Choice)
	return c.SafekeepingPlace
}

func (c *CorporateActionOption121) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return c.FinancialInstrumentIdentification
}

func (c *CorporateActionOption121) AddTotalEligibleBalance() *SignedQuantityFormat8 {
	c.TotalEligibleBalance = new(SignedQuantityFormat8)
	return c.TotalEligibleBalance
}

func (c *CorporateActionOption121) AddInstructedBalance() *SignedQuantityFormat8 {
	c.InstructedBalance = new(SignedQuantityFormat8)
	return c.InstructedBalance
}

func (c *CorporateActionOption121) AddUninstructedBalance() *SignedQuantityFormat8 {
	c.UninstructedBalance = new(SignedQuantityFormat8)
	return c.UninstructedBalance
}

func (c *CorporateActionOption121) AddStatusQuantity() *Quantity10Choice {
	c.StatusQuantity = new(Quantity10Choice)
	return c.StatusQuantity
}
