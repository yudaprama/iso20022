package model

// Describes how interest rates are reported.
type InterestRateContractTerm1 struct {

	// Unit for the rate basis.
	Unit *RateBasis1Code `xml:"Unit"`

	// Value of the contract term in number of units.
	Value *Number `xml:"Val"`
}

func (i *InterestRateContractTerm1) SetUnit(value string) {
	i.Unit = (*RateBasis1Code)(&value)
}

func (i *InterestRateContractTerm1) SetValue(value string) {
	i.Value = (*Number)(&value)
}
