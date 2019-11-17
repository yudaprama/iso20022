package model

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails31 struct {

	// Balance to which the payment applies (less or equal to the total eligible balance).
	ConfirmedBalance *BalanceFormat5Choice `xml:"ConfdBal"`

	// Total balance of securities eligible for this corporate action event. The entitlement calculation is based on this balance.
	TotalEligibleBalance *TotalEligibleBalanceFormat8 `xml:"TtlElgblBal,omitempty"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *BalanceFormat5Choice `xml:"BlckdBal,omitempty"`

	// Balance of financial instruments that have been borrowed from another party.
	BorrowedBalance *BalanceFormat5Choice `xml:"BrrwdBal,omitempty"`

	// Balance of securities that belong to a third party and that are held for the purpose of collateralisation.
	CollateralInBalance *BalanceFormat5Choice `xml:"CollInBal,omitempty"`

	// Balance of securities that belong to the safekeeping account indicated within this message, and are deposited with a third party for the purpose of collateralisation.
	CollateralOutBalance *BalanceFormat5Choice `xml:"CollOutBal,omitempty"`

	// Balance of financial instruments that have been loaned to a third party.
	OnLoanBalance *BalanceFormat5Choice `xml:"OnLnBal,omitempty"`

	// Balance of financial instruments that are pending delivery.
	PendingDeliveryBalance []*BalanceFormat6Choice `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*BalanceFormat6Choice `xml:"PdgRctBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *BalanceFormat5Choice `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance []*BalanceFormat6Choice `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *BalanceFormat5Choice `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *BalanceFormat5Choice `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *BalanceFormat5Choice `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
	RegisteredBalance *BalanceFormat5Choice `xml:"RegdBal,omitempty"`

	// Balance that has been affected by the process run through the event.
	AffectedBalance *BalanceFormat5Choice `xml:"AfctdBal,omitempty"`

	// Balance that has not been affected by the process run through the event.
	UnaffectedBalance *BalanceFormat5Choice `xml:"UafctdBal,omitempty"`
}

func (c *CorporateActionBalanceDetails31) AddConfirmedBalance() *BalanceFormat5Choice {
	c.ConfirmedBalance = new(BalanceFormat5Choice)
	return c.ConfirmedBalance
}

func (c *CorporateActionBalanceDetails31) AddTotalEligibleBalance() *TotalEligibleBalanceFormat8 {
	c.TotalEligibleBalance = new(TotalEligibleBalanceFormat8)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails31) AddBlockedBalance() *BalanceFormat5Choice {
	c.BlockedBalance = new(BalanceFormat5Choice)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails31) AddBorrowedBalance() *BalanceFormat5Choice {
	c.BorrowedBalance = new(BalanceFormat5Choice)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails31) AddCollateralInBalance() *BalanceFormat5Choice {
	c.CollateralInBalance = new(BalanceFormat5Choice)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails31) AddCollateralOutBalance() *BalanceFormat5Choice {
	c.CollateralOutBalance = new(BalanceFormat5Choice)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails31) AddOnLoanBalance() *BalanceFormat5Choice {
	c.OnLoanBalance = new(BalanceFormat5Choice)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails31) AddPendingDeliveryBalance() *BalanceFormat6Choice {
	newValue := new(BalanceFormat6Choice)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails31) AddPendingReceiptBalance() *BalanceFormat6Choice {
	newValue := new(BalanceFormat6Choice)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails31) AddOutForRegistrationBalance() *BalanceFormat5Choice {
	c.OutForRegistrationBalance = new(BalanceFormat5Choice)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails31) AddSettlementPositionBalance() *BalanceFormat6Choice {
	newValue := new(BalanceFormat6Choice)
	c.SettlementPositionBalance = append(c.SettlementPositionBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails31) AddStreetPositionBalance() *BalanceFormat5Choice {
	c.StreetPositionBalance = new(BalanceFormat5Choice)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails31) AddTradeDatePositionBalance() *BalanceFormat5Choice {
	c.TradeDatePositionBalance = new(BalanceFormat5Choice)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails31) AddInTransshipmentBalance() *BalanceFormat5Choice {
	c.InTransshipmentBalance = new(BalanceFormat5Choice)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails31) AddRegisteredBalance() *BalanceFormat5Choice {
	c.RegisteredBalance = new(BalanceFormat5Choice)
	return c.RegisteredBalance
}

func (c *CorporateActionBalanceDetails31) AddAffectedBalance() *BalanceFormat5Choice {
	c.AffectedBalance = new(BalanceFormat5Choice)
	return c.AffectedBalance
}

func (c *CorporateActionBalanceDetails31) AddUnaffectedBalance() *BalanceFormat5Choice {
	c.UnaffectedBalance = new(BalanceFormat5Choice)
	return c.UnaffectedBalance
}
