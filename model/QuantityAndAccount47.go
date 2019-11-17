package model

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount47 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity10Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *RestrictedFINXMax210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown38 `xml:"QtyBrkdwn,omitempty"`
}

func (q *QuantityAndAccount47) AddSettlementQuantity() *Quantity10Choice {
	q.SettlementQuantity = new(Quantity10Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount47) SetDenominationChoice(value string) {
	q.DenominationChoice = (*RestrictedFINXMax210Text)(&value)
}

func (q *QuantityAndAccount47) AddAccountOwner() *PartyIdentification109 {
	q.AccountOwner = new(PartyIdentification109)
	return q.AccountOwner
}

func (q *QuantityAndAccount47) AddSafekeepingAccount() *SecuritiesAccount30 {
	q.SafekeepingAccount = new(SecuritiesAccount30)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount47) AddCashAccount() *CashAccountIdentification6Choice {
	q.CashAccount = new(CashAccountIdentification6Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount47) AddSafekeepingPlace() *SafeKeepingPlace2 {
	q.SafekeepingPlace = new(SafeKeepingPlace2)
	return q.SafekeepingPlace
}

func (q *QuantityAndAccount47) AddQuantityBreakdown() *QuantityBreakdown38 {
	newValue := new(QuantityBreakdown38)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}
