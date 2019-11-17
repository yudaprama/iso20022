package model

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount17 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown3 `xml:"QtyBrkdwn,omitempty"`
}

func (q *QuantityAndAccount17) AddSettlementQuantity() *Quantity6Choice {
	q.SettlementQuantity = new(Quantity6Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount17) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount17) AddAccountOwner() *PartyIdentification36Choice {
	q.AccountOwner = new(PartyIdentification36Choice)
	return q.AccountOwner
}

func (q *QuantityAndAccount17) AddSafekeepingAccount() *SecuritiesAccount13 {
	q.SafekeepingAccount = new(SecuritiesAccount13)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount17) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount17) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	q.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return q.SafekeepingPlace
}

func (q *QuantityAndAccount17) AddQuantityBreakdown() *QuantityBreakdown3 {
	newValue := new(QuantityBreakdown3)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}
