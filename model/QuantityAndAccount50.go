package model

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount50 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity15Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *RestrictedFINXMax210Text `xml:"DnmtnChc,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown38 `xml:"QtyBrkdwn,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount50) AddSettlementQuantity() *FinancialInstrumentQuantity15Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount50) SetDenominationChoice(value string) {
	q.DenominationChoice = (*RestrictedFINXMax210Text)(&value)
}

func (q *QuantityAndAccount50) AddCashAccount() *CashAccountIdentification6Choice {
	q.CashAccount = new(CashAccountIdentification6Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount50) AddQuantityBreakdown() *QuantityBreakdown38 {
	newValue := new(QuantityBreakdown38)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

func (q *QuantityAndAccount50) AddSafekeepingPlace() *SafeKeepingPlace2 {
	q.SafekeepingPlace = new(SafeKeepingPlace2)
	return q.SafekeepingPlace
}
