package model

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount46 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *Quantity10Choice `xml:"SttldQty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity15Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity15Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Amount of money previously settled.
	PreviouslySettledAmount *AmountAndDirection19 `xml:"PrevslySttldAmt,omitempty"`

	// Amount of money remaining to be settled.
	RemainingToBeSettledAmount *AmountAndDirection19 `xml:"RmngToBeSttldAmt,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount46) AddSettledQuantity() *Quantity10Choice {
	q.SettledQuantity = new(Quantity10Choice)
	return q.SettledQuantity
}

func (q *QuantityAndAccount46) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity15Choice {
	q.PreviouslySettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.PreviouslySettledQuantity
}

func (q *QuantityAndAccount46) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity15Choice {
	q.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.RemainingToBeSettledQuantity
}

func (q *QuantityAndAccount46) AddPreviouslySettledAmount() *AmountAndDirection19 {
	q.PreviouslySettledAmount = new(AmountAndDirection19)
	return q.PreviouslySettledAmount
}

func (q *QuantityAndAccount46) AddRemainingToBeSettledAmount() *AmountAndDirection19 {
	q.RemainingToBeSettledAmount = new(AmountAndDirection19)
	return q.RemainingToBeSettledAmount
}

func (q *QuantityAndAccount46) AddAccountOwner() *PartyIdentification109 {
	q.AccountOwner = new(PartyIdentification109)
	return q.AccountOwner
}

func (q *QuantityAndAccount46) AddSafekeepingAccount() *SecuritiesAccount30 {
	q.SafekeepingAccount = new(SecuritiesAccount30)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount46) AddCashAccount() *CashAccountIdentification6Choice {
	q.CashAccount = new(CashAccountIdentification6Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount46) AddSafekeepingPlace() *SafeKeepingPlace2 {
	q.SafekeepingPlace = new(SafeKeepingPlace2)
	return q.SafekeepingPlace
}
