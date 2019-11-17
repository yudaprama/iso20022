package model

// Identifies the details of the transaction.
type TransactionDetails83 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate12Choice `xml:"SttlmDt"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity10Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection67 `xml:"SttlmAmt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties49 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties49 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification110 `xml:"Invstr,omitempty"`
}

func (t *TransactionDetails83) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails83) AddTradeDate() *TradeDate6Choice {
	t.TradeDate = new(TradeDate6Choice)
	return t.TradeDate
}

func (t *TransactionDetails83) AddSettlementDate() *SettlementDate12Choice {
	t.SettlementDate = new(SettlementDate12Choice)
	return t.SettlementDate
}

func (t *TransactionDetails83) AddSettlementQuantity() *Quantity10Choice {
	t.SettlementQuantity = new(Quantity10Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails83) AddSettlementAmount() *AmountAndDirection67 {
	t.SettlementAmount = new(AmountAndDirection67)
	return t.SettlementAmount
}

func (t *TransactionDetails83) AddDeliveringSettlementParties() *SettlementParties49 {
	t.DeliveringSettlementParties = new(SettlementParties49)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails83) AddReceivingSettlementParties() *SettlementParties49 {
	t.ReceivingSettlementParties = new(SettlementParties49)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails83) AddInvestor() *PartyIdentification110 {
	t.Investor = new(PartyIdentification110)
	return t.Investor
}
