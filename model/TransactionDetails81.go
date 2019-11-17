package model

// Identifies the details of the transaction.
type TransactionDetails81 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, eg, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection51 `xml:"SttlmAmt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate10Choice `xml:"SttlmDt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties40 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties40 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification99 `xml:"Invstr,omitempty"`
}

func (t *TransactionDetails81) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails81) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails81) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails81) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails81) AddSettlementAmount() *AmountAndDirection51 {
	t.SettlementAmount = new(AmountAndDirection51)
	return t.SettlementAmount
}

func (t *TransactionDetails81) AddSettlementDate() *SettlementDate10Choice {
	t.SettlementDate = new(SettlementDate10Choice)
	return t.SettlementDate
}

func (t *TransactionDetails81) AddTradeDate() *TradeDate5Choice {
	t.TradeDate = new(TradeDate5Choice)
	return t.TradeDate
}

func (t *TransactionDetails81) AddDeliveringSettlementParties() *SettlementParties40 {
	t.DeliveringSettlementParties = new(SettlementParties40)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails81) AddReceivingSettlementParties() *SettlementParties40 {
	t.ReceivingSettlementParties = new(SettlementParties40)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails81) AddInvestor() *PartyIdentification99 {
	t.Investor = new(PartyIdentification99)
	return t.Investor
}
