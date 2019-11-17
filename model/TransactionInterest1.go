package model

// Provides transaction specific interest information that applies to the underlying transaction.
type TransactionInterest1 struct {

	// Identifies the amount of interest included in the entry amount.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Identifies whether the interest amount included in the entry amount is positive (CRDT) or negative (DBIT).
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies the type of interest.
	Type *InterestType1Choice `xml:"Tp,omitempty"`

	// Set of elements qualifying the interest rate.
	Rate []*Rate1 `xml:"Rate,omitempty"`

	// Range of time between a start date and an end date.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Underlying reason for the interest, eg, yearly credit interest on a savings account.
	Reason *Max35Text `xml:"Rsn,omitempty"`
}

func (t *TransactionInterest1) SetAmount(value, currency string) {
	t.Amount = NewCurrencyAndAmount(value, currency)
}

func (t *TransactionInterest1) SetCreditDebitIndicator(value string) {
	t.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (t *TransactionInterest1) AddType() *InterestType1Choice {
	t.Type = new(InterestType1Choice)
	return t.Type
}

func (t *TransactionInterest1) AddRate() *Rate1 {
	newValue := new(Rate1)
	t.Rate = append(t.Rate, newValue)
	return newValue
}

func (t *TransactionInterest1) AddFromToDate() *DateTimePeriodDetails {
	t.FromToDate = new(DateTimePeriodDetails)
	return t.FromToDate
}

func (t *TransactionInterest1) SetReason(value string) {
	t.Reason = (*Max35Text)(&value)
}
