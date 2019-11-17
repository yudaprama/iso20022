package model

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails30 struct {

	// Total quantity of financial instruments of the balance.
	TotalEligibleBalance *Quantity17Choice `xml:"TtlElgblBal"`

	// Quantity of securities in the sub-balance.
	UninstructedBalance *BalanceFormat5Choice `xml:"UinstdBal"`

	// Provides information about the total instructed balance.
	TotalInstructedBalanceDetails *InstructedBalanceDetails5 `xml:"TtlInstdBalDtls"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *SignedQuantityFormat6 `xml:"BlckdBal,omitempty"`

	// Balance of financial instruments that have been borrowed from another party.
	BorrowedBalance *SignedQuantityFormat6 `xml:"BrrwdBal,omitempty"`

	// Balance of securities that belong to a third party and that are held for the purpose of collateralisation.
	CollateralInBalance *SignedQuantityFormat6 `xml:"CollInBal,omitempty"`

	// Balance of securities that belong to the safekeeping account indicated within this message, and are deposited with a third party for the purpose of collateralisation.
	CollateralOutBalance *SignedQuantityFormat6 `xml:"CollOutBal,omitempty"`

	// Balance of financial instruments that have been loaned to a third party.
	OnLoanBalance *SignedQuantityFormat6 `xml:"OnLnBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *SignedQuantityFormat6 `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance *SignedQuantityFormat6 `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *SignedQuantityFormat6 `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *SignedQuantityFormat6 `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *SignedQuantityFormat6 `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
	RegisteredBalance *SignedQuantityFormat6 `xml:"RegdBal,omitempty"`

	// Position that account holders should return to the account servicer to participate in the event or to fulfil their obligation for the event to be complete, for example, return of securities for late announced drawing.
	ObligatedBalance *SignedQuantityFormat6 `xml:"OblgtdBal,omitempty"`

	// Balance of financial instruments that are pending delivery.
	PendingDeliveryBalance []*PendingBalance3 `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*PendingBalance3 `xml:"PdgRctBal,omitempty"`
}

func (c *CorporateActionBalanceDetails30) AddTotalEligibleBalance() *Quantity17Choice {
	c.TotalEligibleBalance = new(Quantity17Choice)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails30) AddUninstructedBalance() *BalanceFormat5Choice {
	c.UninstructedBalance = new(BalanceFormat5Choice)
	return c.UninstructedBalance
}

func (c *CorporateActionBalanceDetails30) AddTotalInstructedBalanceDetails() *InstructedBalanceDetails5 {
	c.TotalInstructedBalanceDetails = new(InstructedBalanceDetails5)
	return c.TotalInstructedBalanceDetails
}

func (c *CorporateActionBalanceDetails30) AddBlockedBalance() *SignedQuantityFormat6 {
	c.BlockedBalance = new(SignedQuantityFormat6)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails30) AddBorrowedBalance() *SignedQuantityFormat6 {
	c.BorrowedBalance = new(SignedQuantityFormat6)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails30) AddCollateralInBalance() *SignedQuantityFormat6 {
	c.CollateralInBalance = new(SignedQuantityFormat6)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails30) AddCollateralOutBalance() *SignedQuantityFormat6 {
	c.CollateralOutBalance = new(SignedQuantityFormat6)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails30) AddOnLoanBalance() *SignedQuantityFormat6 {
	c.OnLoanBalance = new(SignedQuantityFormat6)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails30) AddOutForRegistrationBalance() *SignedQuantityFormat6 {
	c.OutForRegistrationBalance = new(SignedQuantityFormat6)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails30) AddSettlementPositionBalance() *SignedQuantityFormat6 {
	c.SettlementPositionBalance = new(SignedQuantityFormat6)
	return c.SettlementPositionBalance
}

func (c *CorporateActionBalanceDetails30) AddStreetPositionBalance() *SignedQuantityFormat6 {
	c.StreetPositionBalance = new(SignedQuantityFormat6)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails30) AddTradeDatePositionBalance() *SignedQuantityFormat6 {
	c.TradeDatePositionBalance = new(SignedQuantityFormat6)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails30) AddInTransshipmentBalance() *SignedQuantityFormat6 {
	c.InTransshipmentBalance = new(SignedQuantityFormat6)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails30) AddRegisteredBalance() *SignedQuantityFormat6 {
	c.RegisteredBalance = new(SignedQuantityFormat6)
	return c.RegisteredBalance
}

func (c *CorporateActionBalanceDetails30) AddObligatedBalance() *SignedQuantityFormat6 {
	c.ObligatedBalance = new(SignedQuantityFormat6)
	return c.ObligatedBalance
}

func (c *CorporateActionBalanceDetails30) AddPendingDeliveryBalance() *PendingBalance3 {
	newValue := new(PendingBalance3)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails30) AddPendingReceiptBalance() *PendingBalance3 {
	newValue := new(PendingBalance3)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}
