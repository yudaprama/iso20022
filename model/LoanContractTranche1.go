package model

// Provides details on the tranches defined for the loan contract.
type LoanContractTranche1 struct {

	// Unique sequence number of the tranche.
	TrancheNumber *Number `xml:"TrchNb"`

	// Expected tranche payment date.
	ExpectedDate *ISODate `xml:"XpctdDt"`

	// Amount of the tranche as defined in the loan contract.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Loan tranche due date.
	DueDate *ISODate `xml:"DueDt,omitempty"`

	// Loan tranche duration in a coded form.
	DurationCode *Exact1NumericText `xml:"DrtnCd,omitempty"`

	// Indicates whether this tranche is the last tranche of the full report.
	LastTrancheIndicator *YesNoIndicator `xml:"LastTrchInd,omitempty"`
}

func (l *LoanContractTranche1) SetTrancheNumber(value string) {
	l.TrancheNumber = (*Number)(&value)
}

func (l *LoanContractTranche1) SetExpectedDate(value string) {
	l.ExpectedDate = (*ISODate)(&value)
}

func (l *LoanContractTranche1) SetAmount(value, currency string) {
	l.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (l *LoanContractTranche1) SetDueDate(value string) {
	l.DueDate = (*ISODate)(&value)
}

func (l *LoanContractTranche1) SetDurationCode(value string) {
	l.DurationCode = (*Exact1NumericText)(&value)
}

func (l *LoanContractTranche1) SetLastTrancheIndicator(value string) {
	l.LastTrancheIndicator = (*YesNoIndicator)(&value)
}
