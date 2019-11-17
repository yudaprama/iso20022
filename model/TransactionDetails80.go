package model

// Identifies the details of the transaction.
type TransactionDetails80 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount24 `xml:"SfkpgAcct"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection51 `xml:"SttlmAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate10Choice `xml:"SttlmDt"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties40 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties40 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification99 `xml:"Invstr,omitempty"`
}

func (t *TransactionDetails80) AddAccountOwner() *PartyIdentification98 {
	t.AccountOwner = new(PartyIdentification98)
	return t.AccountOwner
}

func (t *TransactionDetails80) AddSafekeepingAccount() *SecuritiesAccount24 {
	t.SafekeepingAccount = new(SecuritiesAccount24)
	return t.SafekeepingAccount
}

func (t *TransactionDetails80) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails80) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails80) AddSettlementAmount() *AmountAndDirection51 {
	t.SettlementAmount = new(AmountAndDirection51)
	return t.SettlementAmount
}

func (t *TransactionDetails80) AddTradeDate() *TradeDate5Choice {
	t.TradeDate = new(TradeDate5Choice)
	return t.TradeDate
}

func (t *TransactionDetails80) AddSettlementDate() *SettlementDate10Choice {
	t.SettlementDate = new(SettlementDate10Choice)
	return t.SettlementDate
}

func (t *TransactionDetails80) AddDeliveringSettlementParties() *SettlementParties40 {
	t.DeliveringSettlementParties = new(SettlementParties40)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails80) AddReceivingSettlementParties() *SettlementParties40 {
	t.ReceivingSettlementParties = new(SettlementParties40)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails80) AddInvestor() *PartyIdentification99 {
	t.Investor = new(PartyIdentification99)
	return t.Investor
}
