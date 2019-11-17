package model

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount59 struct {

	// Denomination of the security to be received or delivered.
	DenominationChoice *RestrictedFINXMax210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification119 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown38 `xml:"QtyBrkdwn,omitempty"`
}

func (q *QuantityAndAccount59) SetDenominationChoice(value string) {
	q.DenominationChoice = (*RestrictedFINXMax210Text)(&value)
}

func (q *QuantityAndAccount59) AddAccountOwner() *PartyIdentification119 {
	q.AccountOwner = new(PartyIdentification119)
	return q.AccountOwner
}

func (q *QuantityAndAccount59) AddCashAccount() *CashAccountIdentification6Choice {
	q.CashAccount = new(CashAccountIdentification6Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount59) AddSafekeepingPlace() *SafeKeepingPlace2 {
	q.SafekeepingPlace = new(SafeKeepingPlace2)
	return q.SafekeepingPlace
}

func (q *QuantityAndAccount59) AddQuantityBreakdown() *QuantityBreakdown38 {
	newValue := new(QuantityBreakdown38)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}
