package model

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails9 struct {

	// Total quantity of financial instruments of the balance.
	TotalEligibleBalance *Quantity3Choice `xml:"TtlElgblBal"`

	// Quantity of securities in the sub-balance.
	UninstructedBalance *BalanceFormat1Choice `xml:"UinstdBal"`

	// Provides information about the total instructed balance.
	TotalInstructedBalanceDetails *InstructedBalanceDetails3 `xml:"TtlInstdBalDtls"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *SignedQuantityFormat2 `xml:"BlckdBal,omitempty"`

	// Balance of financial instruments that have been borrowed from another party.
	BorrowedBalance *SignedQuantityFormat2 `xml:"BrrwdBal,omitempty"`

	// Balance of securities that belong to a third party and that are held for the purpose of collateralisation.
	CollateralInBalance *SignedQuantityFormat2 `xml:"CollInBal,omitempty"`

	// Balance of securities that belong to the safekeeping account indicated within this message, and are deposited with a third party for the purpose of collateralisation.
	CollateralOutBalance *SignedQuantityFormat2 `xml:"CollOutBal,omitempty"`

	// Balance of financial instruments that have been loaned to a third party.
	OnLoanBalance *SignedQuantityFormat2 `xml:"OnLnBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *SignedQuantityFormat2 `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance *SignedQuantityFormat2 `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *SignedQuantityFormat2 `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *SignedQuantityFormat2 `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *SignedQuantityFormat2 `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
	RegisteredBalance *SignedQuantityFormat2 `xml:"RegdBal,omitempty"`

	// Position that account holders should return to the account servicer to participate in the event or to fulfil their obligation for the event to be complete, for example, return of securities for late announced drawing.
	ObligatedBalance *SignedQuantityFormat2 `xml:"OblgtdBal,omitempty"`

	// Balance of financial instruments that are pending delivery.
	PendingDeliveryBalance []*PendingBalance1 `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*PendingBalance1 `xml:"PdgRctBal,omitempty"`
}

func (c *CorporateActionBalanceDetails9) AddTotalEligibleBalance() *Quantity3Choice {
	c.TotalEligibleBalance = new(Quantity3Choice)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails9) AddUninstructedBalance() *BalanceFormat1Choice {
	c.UninstructedBalance = new(BalanceFormat1Choice)
	return c.UninstructedBalance
}

func (c *CorporateActionBalanceDetails9) AddTotalInstructedBalanceDetails() *InstructedBalanceDetails3 {
	c.TotalInstructedBalanceDetails = new(InstructedBalanceDetails3)
	return c.TotalInstructedBalanceDetails
}

func (c *CorporateActionBalanceDetails9) AddBlockedBalance() *SignedQuantityFormat2 {
	c.BlockedBalance = new(SignedQuantityFormat2)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails9) AddBorrowedBalance() *SignedQuantityFormat2 {
	c.BorrowedBalance = new(SignedQuantityFormat2)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails9) AddCollateralInBalance() *SignedQuantityFormat2 {
	c.CollateralInBalance = new(SignedQuantityFormat2)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails9) AddCollateralOutBalance() *SignedQuantityFormat2 {
	c.CollateralOutBalance = new(SignedQuantityFormat2)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails9) AddOnLoanBalance() *SignedQuantityFormat2 {
	c.OnLoanBalance = new(SignedQuantityFormat2)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails9) AddOutForRegistrationBalance() *SignedQuantityFormat2 {
	c.OutForRegistrationBalance = new(SignedQuantityFormat2)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails9) AddSettlementPositionBalance() *SignedQuantityFormat2 {
	c.SettlementPositionBalance = new(SignedQuantityFormat2)
	return c.SettlementPositionBalance
}

func (c *CorporateActionBalanceDetails9) AddStreetPositionBalance() *SignedQuantityFormat2 {
	c.StreetPositionBalance = new(SignedQuantityFormat2)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails9) AddTradeDatePositionBalance() *SignedQuantityFormat2 {
	c.TradeDatePositionBalance = new(SignedQuantityFormat2)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails9) AddInTransshipmentBalance() *SignedQuantityFormat2 {
	c.InTransshipmentBalance = new(SignedQuantityFormat2)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails9) AddRegisteredBalance() *SignedQuantityFormat2 {
	c.RegisteredBalance = new(SignedQuantityFormat2)
	return c.RegisteredBalance
}

func (c *CorporateActionBalanceDetails9) AddObligatedBalance() *SignedQuantityFormat2 {
	c.ObligatedBalance = new(SignedQuantityFormat2)
	return c.ObligatedBalance
}

func (c *CorporateActionBalanceDetails9) AddPendingDeliveryBalance() *PendingBalance1 {
	newValue := new(PendingBalance1)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails9) AddPendingReceiptBalance() *PendingBalance1 {
	newValue := new(PendingBalance1)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}
