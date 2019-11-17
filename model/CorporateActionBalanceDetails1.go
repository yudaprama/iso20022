package model

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails1 struct {

	// Total quantity of financial instruments of the balance.
	TotalEligibleBalance *Quantity3Choice `xml:"TtlElgblBal,omitempty"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *BalanceFormat1Choice `xml:"BlckdBal,omitempty"`

	// Quantity of securities in the sub-balance.
	BorrowedBalance *BalanceFormat1Choice `xml:"BrrwdBal,omitempty"`

	// Quantity of securities in the sub-balance.
	CollateralInBalance *BalanceFormat1Choice `xml:"CollInBal,omitempty"`

	// Quantity of securities in the sub-balance.
	CollateralOutBalance *BalanceFormat1Choice `xml:"CollOutBal,omitempty"`

	// Quantity of securities in the sub-balance.
	OnLoanBalance *BalanceFormat1Choice `xml:"OnLnBal,omitempty"`

	// Quantity of securities in the sub-balance.
	PendingDeliveryBalance []*BalanceFormat1Choice `xml:"PdgDlvryBal,omitempty"`

	// Quantity of securities in the sub-balance.
	PendingReceiptBalance []*BalanceFormat1Choice `xml:"PdgRctBal,omitempty"`

	// Quantity of securities in the sub-balance.
	OutForRegistrationBalance *BalanceFormat1Choice `xml:"OutForRegnBal,omitempty"`

	// Quantity of securities in the sub-balance.
	SettlementPositionBalance *BalanceFormat1Choice `xml:"SttlmPosBal,omitempty"`

	// Quantity of securities in the sub-balance.
	StreetPositionBalance *BalanceFormat1Choice `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *BalanceFormat1Choice `xml:"TradDtPosBal,omitempty"`

	// Quantity of securities in the sub-balance.
	InTransshipmentBalance *BalanceFormat1Choice `xml:"InTrnsShipmntBal,omitempty"`

	// Quantity of securities in the sub-balance.
	RegisteredBalance *BalanceFormat1Choice `xml:"RegdBal,omitempty"`

	// Position that account holders should return to the account servicer to participate in the event or to fulfil their obligation for the event to be complete, for example, return of securities for late announced drawing.
	ObligatedBalance *BalanceFormat1Choice `xml:"OblgtdBal,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *BalanceFormat1Choice `xml:"UinstdBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *BalanceFormat1Choice `xml:"InstdBal,omitempty"`

	// Balance that has been affected by the process run through the event.
	AffectedBalance *BalanceFormat1Choice `xml:"AfctdBal,omitempty"`

	// Balance that has not been affected by the process run through the event.
	UnaffectedBalance *BalanceFormat1Choice `xml:"UafctdBal,omitempty"`
}

func (c *CorporateActionBalanceDetails1) AddTotalEligibleBalance() *Quantity3Choice {
	c.TotalEligibleBalance = new(Quantity3Choice)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails1) AddBlockedBalance() *BalanceFormat1Choice {
	c.BlockedBalance = new(BalanceFormat1Choice)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails1) AddBorrowedBalance() *BalanceFormat1Choice {
	c.BorrowedBalance = new(BalanceFormat1Choice)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails1) AddCollateralInBalance() *BalanceFormat1Choice {
	c.CollateralInBalance = new(BalanceFormat1Choice)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails1) AddCollateralOutBalance() *BalanceFormat1Choice {
	c.CollateralOutBalance = new(BalanceFormat1Choice)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails1) AddOnLoanBalance() *BalanceFormat1Choice {
	c.OnLoanBalance = new(BalanceFormat1Choice)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails1) AddPendingDeliveryBalance() *BalanceFormat1Choice {
	newValue := new(BalanceFormat1Choice)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails1) AddPendingReceiptBalance() *BalanceFormat1Choice {
	newValue := new(BalanceFormat1Choice)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails1) AddOutForRegistrationBalance() *BalanceFormat1Choice {
	c.OutForRegistrationBalance = new(BalanceFormat1Choice)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails1) AddSettlementPositionBalance() *BalanceFormat1Choice {
	c.SettlementPositionBalance = new(BalanceFormat1Choice)
	return c.SettlementPositionBalance
}

func (c *CorporateActionBalanceDetails1) AddStreetPositionBalance() *BalanceFormat1Choice {
	c.StreetPositionBalance = new(BalanceFormat1Choice)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails1) AddTradeDatePositionBalance() *BalanceFormat1Choice {
	c.TradeDatePositionBalance = new(BalanceFormat1Choice)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails1) AddInTransshipmentBalance() *BalanceFormat1Choice {
	c.InTransshipmentBalance = new(BalanceFormat1Choice)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails1) AddRegisteredBalance() *BalanceFormat1Choice {
	c.RegisteredBalance = new(BalanceFormat1Choice)
	return c.RegisteredBalance
}

func (c *CorporateActionBalanceDetails1) AddObligatedBalance() *BalanceFormat1Choice {
	c.ObligatedBalance = new(BalanceFormat1Choice)
	return c.ObligatedBalance
}

func (c *CorporateActionBalanceDetails1) AddUninstructedBalance() *BalanceFormat1Choice {
	c.UninstructedBalance = new(BalanceFormat1Choice)
	return c.UninstructedBalance
}

func (c *CorporateActionBalanceDetails1) AddInstructedBalance() *BalanceFormat1Choice {
	c.InstructedBalance = new(BalanceFormat1Choice)
	return c.InstructedBalance
}

func (c *CorporateActionBalanceDetails1) AddAffectedBalance() *BalanceFormat1Choice {
	c.AffectedBalance = new(BalanceFormat1Choice)
	return c.AffectedBalance
}

func (c *CorporateActionBalanceDetails1) AddUnaffectedBalance() *BalanceFormat1Choice {
	c.UnaffectedBalance = new(BalanceFormat1Choice)
	return c.UnaffectedBalance
}
