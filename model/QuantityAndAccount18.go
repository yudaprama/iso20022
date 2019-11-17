package model

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount18 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *Quantity6Choice `xml:"SttldQty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity1Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity1Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Amount of money previously settled.
	PreviouslySettledAmount *AmountAndDirection5 `xml:"PrevslySttldAmt,omitempty"`

	// Amount of money remaining to be settled.
	RemainingToBeSettledAmount *AmountAndDirection5 `xml:"RmngToBeSttldAmt,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount18) AddSettledQuantity() *Quantity6Choice {
	q.SettledQuantity = new(Quantity6Choice)
	return q.SettledQuantity
}

func (q *QuantityAndAccount18) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity1Choice {
	q.PreviouslySettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.PreviouslySettledQuantity
}

func (q *QuantityAndAccount18) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity1Choice {
	q.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.RemainingToBeSettledQuantity
}

func (q *QuantityAndAccount18) AddPreviouslySettledAmount() *AmountAndDirection5 {
	q.PreviouslySettledAmount = new(AmountAndDirection5)
	return q.PreviouslySettledAmount
}

func (q *QuantityAndAccount18) AddRemainingToBeSettledAmount() *AmountAndDirection5 {
	q.RemainingToBeSettledAmount = new(AmountAndDirection5)
	return q.RemainingToBeSettledAmount
}

func (q *QuantityAndAccount18) AddAccountOwner() *PartyIdentification36Choice {
	q.AccountOwner = new(PartyIdentification36Choice)
	return q.AccountOwner
}

func (q *QuantityAndAccount18) AddSafekeepingAccount() *SecuritiesAccount13 {
	q.SafekeepingAccount = new(SecuritiesAccount13)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount18) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount18) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	q.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return q.SafekeepingPlace
}
