package model

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails32 struct {

	// Total quantity of financial instruments of the balance.
	TotalEligibleBalance *Quantity17Choice `xml:"TtlElgblBal,omitempty"`

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
	PendingDeliveryBalance []*BalanceFormat5Choice `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*BalanceFormat5Choice `xml:"PdgRctBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *BalanceFormat5Choice `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance *BalanceFormat5Choice `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *BalanceFormat5Choice `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *BalanceFormat5Choice `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *BalanceFormat5Choice `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
	RegisteredBalance *BalanceFormat5Choice `xml:"RegdBal,omitempty"`
}

func (c *CorporateActionBalanceDetails32) AddTotalEligibleBalance() *Quantity17Choice {
	c.TotalEligibleBalance = new(Quantity17Choice)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails32) AddBlockedBalance() *BalanceFormat5Choice {
	c.BlockedBalance = new(BalanceFormat5Choice)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails32) AddBorrowedBalance() *BalanceFormat5Choice {
	c.BorrowedBalance = new(BalanceFormat5Choice)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails32) AddCollateralInBalance() *BalanceFormat5Choice {
	c.CollateralInBalance = new(BalanceFormat5Choice)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails32) AddCollateralOutBalance() *BalanceFormat5Choice {
	c.CollateralOutBalance = new(BalanceFormat5Choice)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails32) AddOnLoanBalance() *BalanceFormat5Choice {
	c.OnLoanBalance = new(BalanceFormat5Choice)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails32) AddPendingDeliveryBalance() *BalanceFormat5Choice {
	newValue := new(BalanceFormat5Choice)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails32) AddPendingReceiptBalance() *BalanceFormat5Choice {
	newValue := new(BalanceFormat5Choice)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails32) AddOutForRegistrationBalance() *BalanceFormat5Choice {
	c.OutForRegistrationBalance = new(BalanceFormat5Choice)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails32) AddSettlementPositionBalance() *BalanceFormat5Choice {
	c.SettlementPositionBalance = new(BalanceFormat5Choice)
	return c.SettlementPositionBalance
}

func (c *CorporateActionBalanceDetails32) AddStreetPositionBalance() *BalanceFormat5Choice {
	c.StreetPositionBalance = new(BalanceFormat5Choice)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails32) AddTradeDatePositionBalance() *BalanceFormat5Choice {
	c.TradeDatePositionBalance = new(BalanceFormat5Choice)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails32) AddInTransshipmentBalance() *BalanceFormat5Choice {
	c.InTransshipmentBalance = new(BalanceFormat5Choice)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails32) AddRegisteredBalance() *BalanceFormat5Choice {
	c.RegisteredBalance = new(BalanceFormat5Choice)
	return c.RegisteredBalance
}
