package model

// Provides details on the type of margin amounts.
type Margin4 struct {

	// Specifies the type of margin that is calculated.
	Type *MarginType1Choice `xml:"Tp"`

	// Provides the margin amount in the reporting currency and optionally in the original currency.
	Amount *Amount2 `xml:"Amt"`

	// Specifies whether the margin type position is short or long, that is, whether the balance is a negative or positive balance.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`
}

func (m *Margin4) AddType() *MarginType1Choice {
	m.Type = new(MarginType1Choice)
	return m.Type
}

func (m *Margin4) AddAmount() *Amount2 {
	m.Amount = new(Amount2)
	return m.Amount
}

func (m *Margin4) SetCreditDebitIndicator(value string) {
	m.CreditDebitIndicator = (*CreditDebitCode)(&value)
}
