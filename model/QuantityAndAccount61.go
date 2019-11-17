package model

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount61 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity15Choice `xml:"SttlmQty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification119 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount61) AddSettlementQuantity() *FinancialInstrumentQuantity15Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount61) AddAccountOwner() *PartyIdentification119 {
	q.AccountOwner = new(PartyIdentification119)
	return q.AccountOwner
}

func (q *QuantityAndAccount61) AddSafekeepingAccount() *SecuritiesAccount30 {
	q.SafekeepingAccount = new(SecuritiesAccount30)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount61) AddCashAccount() *CashAccountIdentification6Choice {
	q.CashAccount = new(CashAccountIdentification6Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount61) AddSafekeepingPlace() *SafeKeepingPlace2 {
	q.SafekeepingPlace = new(SafeKeepingPlace2)
	return q.SafekeepingPlace
}
