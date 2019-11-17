package model

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount7 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount7) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount7) AddAccountOwner() *PartyIdentification13Choice {
	q.AccountOwner = new(PartyIdentification13Choice)
	return q.AccountOwner
}

func (q *QuantityAndAccount7) AddSafekeepingAccount() *SecuritiesAccount13 {
	q.SafekeepingAccount = new(SecuritiesAccount13)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount7) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount7) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	q.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return q.SafekeepingPlace
}
