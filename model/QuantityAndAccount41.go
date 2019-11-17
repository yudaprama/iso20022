package model

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount41 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *Quantity6Choice `xml:"SttldQty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity1Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity1Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Amount of money previously settled.
	PreviouslySettledAmount *AmountAndDirection52 `xml:"PrevslySttldAmt,omitempty"`

	// Amount of money remaining to be settled.
	RemainingToBeSettledAmount *AmountAndDirection52 `xml:"RmngToBeSttldAmt,omitempty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount24 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown29 `xml:"QtyBrkdwn,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount41) AddSettledQuantity() *Quantity6Choice {
	q.SettledQuantity = new(Quantity6Choice)
	return q.SettledQuantity
}

func (q *QuantityAndAccount41) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity1Choice {
	q.PreviouslySettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.PreviouslySettledQuantity
}

func (q *QuantityAndAccount41) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity1Choice {
	q.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.RemainingToBeSettledQuantity
}

func (q *QuantityAndAccount41) AddPreviouslySettledAmount() *AmountAndDirection52 {
	q.PreviouslySettledAmount = new(AmountAndDirection52)
	return q.PreviouslySettledAmount
}

func (q *QuantityAndAccount41) AddRemainingToBeSettledAmount() *AmountAndDirection52 {
	q.RemainingToBeSettledAmount = new(AmountAndDirection52)
	return q.RemainingToBeSettledAmount
}

func (q *QuantityAndAccount41) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount41) AddAccountOwner() *PartyIdentification98 {
	q.AccountOwner = new(PartyIdentification98)
	return q.AccountOwner
}

func (q *QuantityAndAccount41) AddSafekeepingAccount() *SecuritiesAccount24 {
	q.SafekeepingAccount = new(SecuritiesAccount24)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount41) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount41) AddQuantityBreakdown() *QuantityBreakdown29 {
	newValue := new(QuantityBreakdown29)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

func (q *QuantityAndAccount41) AddSafekeepingPlace() *SafeKeepingPlace1 {
	q.SafekeepingPlace = new(SafeKeepingPlace1)
	return q.SafekeepingPlace
}
