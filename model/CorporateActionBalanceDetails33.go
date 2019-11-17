package model

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails33 struct {

	// Total quantity of financial instruments of the balance.
	TotalEligibleBalance *Quantity22Choice `xml:"TtlElgblBal"`

	// Quantity of securities in the sub-balance.
	UninstructedBalance *BalanceFormat7Choice `xml:"UinstdBal"`

	// Provides information about the total instructed balance.
	TotalInstructedBalanceDetails *InstructedBalanceDetails6 `xml:"TtlInstdBalDtls"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *SignedQuantityFormat9 `xml:"BlckdBal,omitempty"`

	// Balance of financial instruments that have been borrowed from another party.
	BorrowedBalance *SignedQuantityFormat9 `xml:"BrrwdBal,omitempty"`

	// Balance of securities that belong to a third party and that are held for the purpose of collateralisation.
	CollateralInBalance *SignedQuantityFormat9 `xml:"CollInBal,omitempty"`

	// Balance of securities that belong to the safekeeping account indicated within this message, and are deposited with a third party for the purpose of collateralisation.
	CollateralOutBalance *SignedQuantityFormat9 `xml:"CollOutBal,omitempty"`

	// Balance of financial instruments that have been loaned to a third party.
	OnLoanBalance *SignedQuantityFormat9 `xml:"OnLnBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *SignedQuantityFormat9 `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance *SignedQuantityFormat9 `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *SignedQuantityFormat9 `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *SignedQuantityFormat9 `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *SignedQuantityFormat9 `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
	RegisteredBalance *SignedQuantityFormat9 `xml:"RegdBal,omitempty"`

	// Position that account holders should return to the account servicer to participate in the event or to fulfil their obligation for the event to be complete, for example, return of securities for late announced drawing.
	ObligatedBalance *SignedQuantityFormat9 `xml:"OblgtdBal,omitempty"`

	// Balance of financial instruments that are pending delivery.
	PendingDeliveryBalance []*PendingBalance4 `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*PendingBalance4 `xml:"PdgRctBal,omitempty"`
}

func (c *CorporateActionBalanceDetails33) AddTotalEligibleBalance() *Quantity22Choice {
	c.TotalEligibleBalance = new(Quantity22Choice)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails33) AddUninstructedBalance() *BalanceFormat7Choice {
	c.UninstructedBalance = new(BalanceFormat7Choice)
	return c.UninstructedBalance
}

func (c *CorporateActionBalanceDetails33) AddTotalInstructedBalanceDetails() *InstructedBalanceDetails6 {
	c.TotalInstructedBalanceDetails = new(InstructedBalanceDetails6)
	return c.TotalInstructedBalanceDetails
}

func (c *CorporateActionBalanceDetails33) AddBlockedBalance() *SignedQuantityFormat9 {
	c.BlockedBalance = new(SignedQuantityFormat9)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails33) AddBorrowedBalance() *SignedQuantityFormat9 {
	c.BorrowedBalance = new(SignedQuantityFormat9)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails33) AddCollateralInBalance() *SignedQuantityFormat9 {
	c.CollateralInBalance = new(SignedQuantityFormat9)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails33) AddCollateralOutBalance() *SignedQuantityFormat9 {
	c.CollateralOutBalance = new(SignedQuantityFormat9)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails33) AddOnLoanBalance() *SignedQuantityFormat9 {
	c.OnLoanBalance = new(SignedQuantityFormat9)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails33) AddOutForRegistrationBalance() *SignedQuantityFormat9 {
	c.OutForRegistrationBalance = new(SignedQuantityFormat9)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails33) AddSettlementPositionBalance() *SignedQuantityFormat9 {
	c.SettlementPositionBalance = new(SignedQuantityFormat9)
	return c.SettlementPositionBalance
}

func (c *CorporateActionBalanceDetails33) AddStreetPositionBalance() *SignedQuantityFormat9 {
	c.StreetPositionBalance = new(SignedQuantityFormat9)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails33) AddTradeDatePositionBalance() *SignedQuantityFormat9 {
	c.TradeDatePositionBalance = new(SignedQuantityFormat9)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails33) AddInTransshipmentBalance() *SignedQuantityFormat9 {
	c.InTransshipmentBalance = new(SignedQuantityFormat9)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails33) AddRegisteredBalance() *SignedQuantityFormat9 {
	c.RegisteredBalance = new(SignedQuantityFormat9)
	return c.RegisteredBalance
}

func (c *CorporateActionBalanceDetails33) AddObligatedBalance() *SignedQuantityFormat9 {
	c.ObligatedBalance = new(SignedQuantityFormat9)
	return c.ObligatedBalance
}

func (c *CorporateActionBalanceDetails33) AddPendingDeliveryBalance() *PendingBalance4 {
	newValue := new(PendingBalance4)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails33) AddPendingReceiptBalance() *PendingBalance4 {
	newValue := new(PendingBalance4)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}
