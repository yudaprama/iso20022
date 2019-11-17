package model

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount44 struct {

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown30 `xml:"QtyBrkdwn,omitempty"`
}

func (q *QuantityAndAccount44) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount44) AddAccountOwner() *PartyIdentification98 {
	q.AccountOwner = new(PartyIdentification98)
	return q.AccountOwner
}

func (q *QuantityAndAccount44) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount44) AddSafekeepingPlace() *SafeKeepingPlace1 {
	q.SafekeepingPlace = new(SafeKeepingPlace1)
	return q.SafekeepingPlace
}

func (q *QuantityAndAccount44) AddQuantityBreakdown() *QuantityBreakdown30 {
	newValue := new(QuantityBreakdown30)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}
