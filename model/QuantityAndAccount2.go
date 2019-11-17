package model

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount2 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *Quantity6Choice `xml:"SttldQty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity1Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity1Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Amount of money previously settled.
	PreviouslySettledAmount *AmountAndDirection7 `xml:"PrevslySttldAmt,omitempty"`

	// Amount of money remaining to be settled.
	RemainingToBeSettledAmount *AmountAndDirection7 `xml:"RmngToBeSttldAmt,omitempty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown3 `xml:"QtyBrkdwn,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount2) AddSettledQuantity() *Quantity6Choice {
	q.SettledQuantity = new(Quantity6Choice)
	return q.SettledQuantity
}

func (q *QuantityAndAccount2) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity1Choice {
	q.PreviouslySettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.PreviouslySettledQuantity
}

func (q *QuantityAndAccount2) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity1Choice {
	q.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.RemainingToBeSettledQuantity
}

func (q *QuantityAndAccount2) AddPreviouslySettledAmount() *AmountAndDirection7 {
	q.PreviouslySettledAmount = new(AmountAndDirection7)
	return q.PreviouslySettledAmount
}

func (q *QuantityAndAccount2) AddRemainingToBeSettledAmount() *AmountAndDirection7 {
	q.RemainingToBeSettledAmount = new(AmountAndDirection7)
	return q.RemainingToBeSettledAmount
}

func (q *QuantityAndAccount2) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount2) AddAccountOwner() *PartyIdentification13Choice {
	q.AccountOwner = new(PartyIdentification13Choice)
	return q.AccountOwner
}

func (q *QuantityAndAccount2) AddSafekeepingAccount() *SecuritiesAccount13 {
	q.SafekeepingAccount = new(SecuritiesAccount13)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount2) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount2) AddQuantityBreakdown() *QuantityBreakdown3 {
	newValue := new(QuantityBreakdown3)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

func (q *QuantityAndAccount2) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	q.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return q.SafekeepingPlace
}
