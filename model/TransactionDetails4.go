package model

// Identifies the details of the transaction.
type TransactionDetails4 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification11 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection8 `xml:"SttlmAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate2Choice `xml:"SttlmDt"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties2 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties2 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification14Choice `xml:"Invstr,omitempty"`
}

func (t *TransactionDetails4) AddAccountOwner() *PartyIdentification13Choice {
	t.AccountOwner = new(PartyIdentification13Choice)
	return t.AccountOwner
}

func (t *TransactionDetails4) AddSafekeepingAccount() *SecuritiesAccount13 {
	t.SafekeepingAccount = new(SecuritiesAccount13)
	return t.SafekeepingAccount
}

func (t *TransactionDetails4) AddFinancialInstrumentIdentification() *SecurityIdentification11 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification11)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails4) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails4) AddSettlementAmount() *AmountAndDirection8 {
	t.SettlementAmount = new(AmountAndDirection8)
	return t.SettlementAmount
}

func (t *TransactionDetails4) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails4) AddSettlementDate() *SettlementDate2Choice {
	t.SettlementDate = new(SettlementDate2Choice)
	return t.SettlementDate
}

func (t *TransactionDetails4) AddDeliveringSettlementParties() *SettlementParties2 {
	t.DeliveringSettlementParties = new(SettlementParties2)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails4) AddReceivingSettlementParties() *SettlementParties2 {
	t.ReceivingSettlementParties = new(SettlementParties2)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails4) AddInvestor() *PartyIdentification14Choice {
	t.Investor = new(PartyIdentification14Choice)
	return t.Investor
}
