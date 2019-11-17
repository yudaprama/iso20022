package model

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails12 struct {

	// Total quantity of financial instruments of the balance.
	TotalEligibleBalance *Quantity3Choice `xml:"TtlElgblBal,omitempty"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *BalanceFormat1Choice `xml:"BlckdBal,omitempty"`

	// Balance of financial instruments that have been borrowed from another party.
	BorrowedBalance *BalanceFormat1Choice `xml:"BrrwdBal,omitempty"`

	// Balance of securities that belong to a third party and that are held for the purpose of collateralisation.
	CollateralInBalance *BalanceFormat1Choice `xml:"CollInBal,omitempty"`

	// Balance of securities that belong to the safekeeping account indicated within this message, and are deposited with a third party for the purpose of collateralisation.
	CollateralOutBalance *BalanceFormat1Choice `xml:"CollOutBal,omitempty"`

	// Balance of financial instruments that have been loaned to a third party.
	OnLoanBalance *BalanceFormat1Choice `xml:"OnLnBal,omitempty"`

	// Balance of financial instruments that are pending delivery.
	PendingDeliveryBalance []*BalanceFormat1Choice `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*BalanceFormat1Choice `xml:"PdgRctBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *BalanceFormat1Choice `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance *BalanceFormat1Choice `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *BalanceFormat1Choice `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *BalanceFormat1Choice `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *BalanceFormat1Choice `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
	RegisteredBalance *BalanceFormat1Choice `xml:"RegdBal,omitempty"`
}

func (c *CorporateActionBalanceDetails12) AddTotalEligibleBalance() *Quantity3Choice {
	c.TotalEligibleBalance = new(Quantity3Choice)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails12) AddBlockedBalance() *BalanceFormat1Choice {
	c.BlockedBalance = new(BalanceFormat1Choice)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails12) AddBorrowedBalance() *BalanceFormat1Choice {
	c.BorrowedBalance = new(BalanceFormat1Choice)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails12) AddCollateralInBalance() *BalanceFormat1Choice {
	c.CollateralInBalance = new(BalanceFormat1Choice)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails12) AddCollateralOutBalance() *BalanceFormat1Choice {
	c.CollateralOutBalance = new(BalanceFormat1Choice)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails12) AddOnLoanBalance() *BalanceFormat1Choice {
	c.OnLoanBalance = new(BalanceFormat1Choice)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails12) AddPendingDeliveryBalance() *BalanceFormat1Choice {
	newValue := new(BalanceFormat1Choice)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails12) AddPendingReceiptBalance() *BalanceFormat1Choice {
	newValue := new(BalanceFormat1Choice)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails12) AddOutForRegistrationBalance() *BalanceFormat1Choice {
	c.OutForRegistrationBalance = new(BalanceFormat1Choice)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails12) AddSettlementPositionBalance() *BalanceFormat1Choice {
	c.SettlementPositionBalance = new(BalanceFormat1Choice)
	return c.SettlementPositionBalance
}

func (c *CorporateActionBalanceDetails12) AddStreetPositionBalance() *BalanceFormat1Choice {
	c.StreetPositionBalance = new(BalanceFormat1Choice)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails12) AddTradeDatePositionBalance() *BalanceFormat1Choice {
	c.TradeDatePositionBalance = new(BalanceFormat1Choice)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails12) AddInTransshipmentBalance() *BalanceFormat1Choice {
	c.InTransshipmentBalance = new(BalanceFormat1Choice)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails12) AddRegisteredBalance() *BalanceFormat1Choice {
	c.RegisteredBalance = new(BalanceFormat1Choice)
	return c.RegisteredBalance
}
