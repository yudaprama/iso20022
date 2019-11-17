package model

// Indicates the independent amount and how it was applied in the calculation. It also provides a description of the amount type.
type IndependentAmount2 struct {

	// Description of the other amount used in the calculation of the independent amount.
	Description *Max140Text `xml:"Desc,omitempty"`

	// Provides the independant amount.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Determines how the independent amount was applied in the calculation.
	// It is either:
	// - before threshold, effectively acting as an add on to exposure,
	// - after threshold where the amount is an add on to the credit support amount and forms part of the variation margin requirement,
	// - segregated where it is treated independently of variation margin for segregation purposes.
	Convention *IndependentAmountConventionType1Code `xml:"Cnvntn"`
}

func (i *IndependentAmount2) SetDescription(value string) {
	i.Description = (*Max140Text)(&value)
}

func (i *IndependentAmount2) SetAmount(value, currency string) {
	i.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (i *IndependentAmount2) SetConvention(value string) {
	i.Convention = (*IndependentAmountConventionType1Code)(&value)
}
