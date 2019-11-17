package model

// Identifies the details of the transaction.
type TransactionDetails92 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, eg, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity10Choice `xml:"SttlmQty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection67 `xml:"SttlmAmt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate15Choice `xml:"SttlmDt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties49 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties49 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification110 `xml:"Invstr,omitempty"`
}

func (t *TransactionDetails92) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails92) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails92) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails92) AddSettlementQuantity() *Quantity10Choice {
	t.SettlementQuantity = new(Quantity10Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails92) AddSafekeepingAccount() *SecuritiesAccount30 {
	t.SafekeepingAccount = new(SecuritiesAccount30)
	return t.SafekeepingAccount
}

func (t *TransactionDetails92) AddSettlementAmount() *AmountAndDirection67 {
	t.SettlementAmount = new(AmountAndDirection67)
	return t.SettlementAmount
}

func (t *TransactionDetails92) AddSettlementDate() *SettlementDate15Choice {
	t.SettlementDate = new(SettlementDate15Choice)
	return t.SettlementDate
}

func (t *TransactionDetails92) AddTradeDate() *TradeDate6Choice {
	t.TradeDate = new(TradeDate6Choice)
	return t.TradeDate
}

func (t *TransactionDetails92) AddDeliveringSettlementParties() *SettlementParties49 {
	t.DeliveringSettlementParties = new(SettlementParties49)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails92) AddReceivingSettlementParties() *SettlementParties49 {
	t.ReceivingSettlementParties = new(SettlementParties49)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails92) AddInvestor() *PartyIdentification110 {
	t.Investor = new(PartyIdentification110)
	return t.Investor
}
