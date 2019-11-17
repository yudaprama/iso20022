package model

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails17 struct {

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
	PendingDeliveryBalance []*BalanceFormat3Choice `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*BalanceFormat3Choice `xml:"PdgRctBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *BalanceFormat1Choice `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance []*BalanceFormat3Choice `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *BalanceFormat1Choice `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *BalanceFormat1Choice `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *BalanceFormat1Choice `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
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

func (c *CorporateActionBalanceDetails17) AddTotalEligibleBalance() *Quantity3Choice {
	c.TotalEligibleBalance = new(Quantity3Choice)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails17) AddBlockedBalance() *BalanceFormat1Choice {
	c.BlockedBalance = new(BalanceFormat1Choice)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails17) AddBorrowedBalance() *BalanceFormat1Choice {
	c.BorrowedBalance = new(BalanceFormat1Choice)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails17) AddCollateralInBalance() *BalanceFormat1Choice {
	c.CollateralInBalance = new(BalanceFormat1Choice)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails17) AddCollateralOutBalance() *BalanceFormat1Choice {
	c.CollateralOutBalance = new(BalanceFormat1Choice)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails17) AddOnLoanBalance() *BalanceFormat1Choice {
	c.OnLoanBalance = new(BalanceFormat1Choice)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails17) AddPendingDeliveryBalance() *BalanceFormat3Choice {
	newValue := new(BalanceFormat3Choice)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails17) AddPendingReceiptBalance() *BalanceFormat3Choice {
	newValue := new(BalanceFormat3Choice)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails17) AddOutForRegistrationBalance() *BalanceFormat1Choice {
	c.OutForRegistrationBalance = new(BalanceFormat1Choice)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails17) AddSettlementPositionBalance() *BalanceFormat3Choice {
	newValue := new(BalanceFormat3Choice)
	c.SettlementPositionBalance = append(c.SettlementPositionBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails17) AddStreetPositionBalance() *BalanceFormat1Choice {
	c.StreetPositionBalance = new(BalanceFormat1Choice)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails17) AddTradeDatePositionBalance() *BalanceFormat1Choice {
	c.TradeDatePositionBalance = new(BalanceFormat1Choice)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails17) AddInTransshipmentBalance() *BalanceFormat1Choice {
	c.InTransshipmentBalance = new(BalanceFormat1Choice)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails17) AddRegisteredBalance() *BalanceFormat1Choice {
	c.RegisteredBalance = new(BalanceFormat1Choice)
	return c.RegisteredBalance
}

func (c *CorporateActionBalanceDetails17) AddObligatedBalance() *BalanceFormat1Choice {
	c.ObligatedBalance = new(BalanceFormat1Choice)
	return c.ObligatedBalance
}

func (c *CorporateActionBalanceDetails17) AddUninstructedBalance() *BalanceFormat1Choice {
	c.UninstructedBalance = new(BalanceFormat1Choice)
	return c.UninstructedBalance
}

func (c *CorporateActionBalanceDetails17) AddInstructedBalance() *BalanceFormat1Choice {
	c.InstructedBalance = new(BalanceFormat1Choice)
	return c.InstructedBalance
}

func (c *CorporateActionBalanceDetails17) AddAffectedBalance() *BalanceFormat1Choice {
	c.AffectedBalance = new(BalanceFormat1Choice)
	return c.AffectedBalance
}

func (c *CorporateActionBalanceDetails17) AddUnaffectedBalance() *BalanceFormat1Choice {
	c.UnaffectedBalance = new(BalanceFormat1Choice)
	return c.UnaffectedBalance
}
