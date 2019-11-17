package model

// Identifies the details of the transaction.
type TransactionDetails10 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification11 `xml:"FinInstrmId"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection8 `xml:"SttlmAmt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties2 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties2 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification14Choice `xml:"Invstr,omitempty"`
}

func (t *TransactionDetails10) AddFinancialInstrumentIdentification() *SecurityIdentification11 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification11)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails10) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails10) AddSettlementDate() *SettlementDate1Choice {
	t.SettlementDate = new(SettlementDate1Choice)
	return t.SettlementDate
}

func (t *TransactionDetails10) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails10) AddSettlementAmount() *AmountAndDirection8 {
	t.SettlementAmount = new(AmountAndDirection8)
	return t.SettlementAmount
}

func (t *TransactionDetails10) AddDeliveringSettlementParties() *SettlementParties2 {
	t.DeliveringSettlementParties = new(SettlementParties2)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails10) AddReceivingSettlementParties() *SettlementParties2 {
	t.ReceivingSettlementParties = new(SettlementParties2)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails10) AddInvestor() *PartyIdentification14Choice {
	t.Investor = new(PartyIdentification14Choice)
	return t.Investor
}
