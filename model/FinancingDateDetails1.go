package model

// Includes a set of dates (e.g. credit date) related to settlement of the financing amount.
type FinancingDateDetails1 struct {

	// Date on which the financing transaction has been booked in an account.
	BookDate []*ISODate `xml:"BookDt,omitempty"`

	// Date on which a financed amount has been credited.
	CreditDate *ISODate `xml:"CdtDt"`

	// Date on which a financed amount has been debited.
	DebitDate *ISODate `xml:"DbtDt,omitempty"`
}

func (f *FinancingDateDetails1) AddBookDate(value string) {
	f.BookDate = append(f.BookDate, (*ISODate)(&value))
}

func (f *FinancingDateDetails1) SetCreditDate(value string) {
	f.CreditDate = (*ISODate)(&value)
}

func (f *FinancingDateDetails1) SetDebitDate(value string) {
	f.DebitDate = (*ISODate)(&value)
}
