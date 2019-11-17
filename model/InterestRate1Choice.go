package model

// Choice between a fixed rate and a variable rate.
type InterestRate1Choice struct {

	// Indicates that the rate is fixed.
	FixedInterestRate *PercentageRate `xml:"FxdIntrstRate"`

	// Provides details about the variable rate.
	VariableInterestRate *VariableInterest1Rate `xml:"VarblIntrstRate"`
}

func (i *InterestRate1Choice) SetFixedInterestRate(value string) {
	i.FixedInterestRate = (*PercentageRate)(&value)
}

func (i *InterestRate1Choice) AddVariableInterestRate() *VariableInterest1Rate {
	i.VariableInterestRate = new(VariableInterest1Rate)
	return i.VariableInterestRate
}
