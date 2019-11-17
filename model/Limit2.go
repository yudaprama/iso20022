package model

// Specifies the minimum value of entries to be reported in the requested message.
type Limit2 struct {

	// Minimum transaction amount to be reported in the requested message.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether the floor limit applies to credit, to debit or to both credit and debit entries.
	CreditDebitIndicator *FloorLimitType1Code `xml:"CdtDbtInd"`
}

func (l *Limit2) SetAmount(value, currency string) {
	l.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (l *Limit2) SetCreditDebitIndicator(value string) {
	l.CreditDebitIndicator = (*FloorLimitType1Code)(&value)
}
