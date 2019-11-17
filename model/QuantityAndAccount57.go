package model

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount57 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *Quantity10Choice `xml:"SttldQty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity15Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity15Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Amount of money previously settled.
	PreviouslySettledAmount *AmountAndDirection57 `xml:"PrevslySttldAmt,omitempty"`

	// Amount of money remaining to be settled.
	RemainingToBeSettledAmount *AmountAndDirection57 `xml:"RmngToBeSttldAmt,omitempty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *RestrictedFINXMax210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification119 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount27 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown44 `xml:"QtyBrkdwn,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount57) AddSettledQuantity() *Quantity10Choice {
	q.SettledQuantity = new(Quantity10Choice)
	return q.SettledQuantity
}

func (q *QuantityAndAccount57) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity15Choice {
	q.PreviouslySettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.PreviouslySettledQuantity
}

func (q *QuantityAndAccount57) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity15Choice {
	q.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.RemainingToBeSettledQuantity
}

func (q *QuantityAndAccount57) AddPreviouslySettledAmount() *AmountAndDirection57 {
	q.PreviouslySettledAmount = new(AmountAndDirection57)
	return q.PreviouslySettledAmount
}

func (q *QuantityAndAccount57) AddRemainingToBeSettledAmount() *AmountAndDirection57 {
	q.RemainingToBeSettledAmount = new(AmountAndDirection57)
	return q.RemainingToBeSettledAmount
}

func (q *QuantityAndAccount57) SetDenominationChoice(value string) {
	q.DenominationChoice = (*RestrictedFINXMax210Text)(&value)
}

func (q *QuantityAndAccount57) AddAccountOwner() *PartyIdentification119 {
	q.AccountOwner = new(PartyIdentification119)
	return q.AccountOwner
}

func (q *QuantityAndAccount57) AddSafekeepingAccount() *SecuritiesAccount27 {
	q.SafekeepingAccount = new(SecuritiesAccount27)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount57) AddCashAccount() *CashAccountIdentification6Choice {
	q.CashAccount = new(CashAccountIdentification6Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount57) AddQuantityBreakdown() *QuantityBreakdown44 {
	newValue := new(QuantityBreakdown44)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

func (q *QuantityAndAccount57) AddSafekeepingPlace() *SafeKeepingPlace2 {
	q.SafekeepingPlace = new(SafeKeepingPlace2)
	return q.SafekeepingPlace
}
