package model

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails35 struct {

	// Balance to which the payment applies (less or equal to the total eligible balance).
	ConfirmedBalance *BalanceFormat7Choice `xml:"ConfdBal"`

	// Total balance of securities eligible for this corporate action event. The entitlement calculation is based on this balance.
	TotalEligibleBalance *TotalEligibleBalanceFormat9 `xml:"TtlElgblBal,omitempty"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *BalanceFormat7Choice `xml:"BlckdBal,omitempty"`

	// Balance of financial instruments that have been borrowed from another party.
	BorrowedBalance *BalanceFormat7Choice `xml:"BrrwdBal,omitempty"`

	// Balance of securities that belong to a third party and that are held for the purpose of collateralisation.
	CollateralInBalance *BalanceFormat7Choice `xml:"CollInBal,omitempty"`

	// Balance of securities that belong to the safekeeping account indicated within this message, and are deposited with a third party for the purpose of collateralisation.
	CollateralOutBalance *BalanceFormat7Choice `xml:"CollOutBal,omitempty"`

	// Balance of financial instruments that have been loaned to a third party.
	OnLoanBalance *BalanceFormat7Choice `xml:"OnLnBal,omitempty"`

	// Balance of financial instruments that are pending delivery.
	PendingDeliveryBalance []*BalanceFormat10Choice `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*BalanceFormat10Choice `xml:"PdgRctBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *BalanceFormat7Choice `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance []*BalanceFormat10Choice `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *BalanceFormat7Choice `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *BalanceFormat7Choice `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *BalanceFormat7Choice `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
	RegisteredBalance *BalanceFormat7Choice `xml:"RegdBal,omitempty"`

	// Balance that has been affected by the process run through the event.
	AffectedBalance *BalanceFormat7Choice `xml:"AfctdBal,omitempty"`

	// Balance that has not been affected by the process run through the event.
	UnaffectedBalance *BalanceFormat7Choice `xml:"UafctdBal,omitempty"`
}

func (c *CorporateActionBalanceDetails35) AddConfirmedBalance() *BalanceFormat7Choice {
	c.ConfirmedBalance = new(BalanceFormat7Choice)
	return c.ConfirmedBalance
}

func (c *CorporateActionBalanceDetails35) AddTotalEligibleBalance() *TotalEligibleBalanceFormat9 {
	c.TotalEligibleBalance = new(TotalEligibleBalanceFormat9)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails35) AddBlockedBalance() *BalanceFormat7Choice {
	c.BlockedBalance = new(BalanceFormat7Choice)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails35) AddBorrowedBalance() *BalanceFormat7Choice {
	c.BorrowedBalance = new(BalanceFormat7Choice)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails35) AddCollateralInBalance() *BalanceFormat7Choice {
	c.CollateralInBalance = new(BalanceFormat7Choice)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails35) AddCollateralOutBalance() *BalanceFormat7Choice {
	c.CollateralOutBalance = new(BalanceFormat7Choice)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails35) AddOnLoanBalance() *BalanceFormat7Choice {
	c.OnLoanBalance = new(BalanceFormat7Choice)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails35) AddPendingDeliveryBalance() *BalanceFormat10Choice {
	newValue := new(BalanceFormat10Choice)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails35) AddPendingReceiptBalance() *BalanceFormat10Choice {
	newValue := new(BalanceFormat10Choice)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails35) AddOutForRegistrationBalance() *BalanceFormat7Choice {
	c.OutForRegistrationBalance = new(BalanceFormat7Choice)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails35) AddSettlementPositionBalance() *BalanceFormat10Choice {
	newValue := new(BalanceFormat10Choice)
	c.SettlementPositionBalance = append(c.SettlementPositionBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails35) AddStreetPositionBalance() *BalanceFormat7Choice {
	c.StreetPositionBalance = new(BalanceFormat7Choice)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails35) AddTradeDatePositionBalance() *BalanceFormat7Choice {
	c.TradeDatePositionBalance = new(BalanceFormat7Choice)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails35) AddInTransshipmentBalance() *BalanceFormat7Choice {
	c.InTransshipmentBalance = new(BalanceFormat7Choice)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails35) AddRegisteredBalance() *BalanceFormat7Choice {
	c.RegisteredBalance = new(BalanceFormat7Choice)
	return c.RegisteredBalance
}

func (c *CorporateActionBalanceDetails35) AddAffectedBalance() *BalanceFormat7Choice {
	c.AffectedBalance = new(BalanceFormat7Choice)
	return c.AffectedBalance
}

func (c *CorporateActionBalanceDetails35) AddUnaffectedBalance() *BalanceFormat7Choice {
	c.UnaffectedBalance = new(BalanceFormat7Choice)
	return c.UnaffectedBalance
}
