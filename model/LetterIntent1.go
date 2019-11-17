package model

// Specifies information about the letter of intent.
type LetterIntent1 struct {

	// Reference of a letter of intent program, in which sales commissions are reduced based on the aggregate of a customer's actual purchase and anticipated purchases, over a specific period of time, and as agreed by the customer.
	LetterIntentReference *Max35Text `xml:"LttrInttRef"`

	// Amount stated on the letter of intent.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt,omitempty"`

	// Start date stated on the letter of intent.
	StartDate *ISODate `xml:"StartDt,omitempty"`

	// End date stated on the letter of intent.
	EndDate *ISODate `xml:"EndDt,omitempty"`
}

func (l *LetterIntent1) SetLetterIntentReference(value string) {
	l.LetterIntentReference = (*Max35Text)(&value)
}

func (l *LetterIntent1) SetAmount(value, currency string) {
	l.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (l *LetterIntent1) SetStartDate(value string) {
	l.StartDate = (*ISODate)(&value)
}

func (l *LetterIntent1) SetEndDate(value string) {
	l.EndDate = (*ISODate)(&value)
}
