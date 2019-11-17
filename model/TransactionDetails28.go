package model

// Identifies the details of the transaction.
type TransactionDetails28 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection8 `xml:"SttlmAmt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties12 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties12 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification37Choice `xml:"Invstr,omitempty"`
}

func (t *TransactionDetails28) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails28) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails28) AddSettlementDate() *SettlementDate1Choice {
	t.SettlementDate = new(SettlementDate1Choice)
	return t.SettlementDate
}

func (t *TransactionDetails28) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails28) AddSettlementAmount() *AmountAndDirection8 {
	t.SettlementAmount = new(AmountAndDirection8)
	return t.SettlementAmount
}

func (t *TransactionDetails28) AddDeliveringSettlementParties() *SettlementParties12 {
	t.DeliveringSettlementParties = new(SettlementParties12)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails28) AddReceivingSettlementParties() *SettlementParties12 {
	t.ReceivingSettlementParties = new(SettlementParties12)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails28) AddInvestor() *PartyIdentification37Choice {
	t.Investor = new(PartyIdentification37Choice)
	return t.Investor
}
