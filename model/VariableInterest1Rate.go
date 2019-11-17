package model

// Provides the index used to define the rate and optionaly the basis point spread.
type VariableInterest1Rate struct {

	// Specifies the index taken into account to calculate the variable interest rate.
	Index *Max35Text `xml:"Indx"`

	// Used to express differences in interest rates, for example, a difference of 0.10% is equivalent to a change of 10 basis points.
	BasisPointSpread *Number `xml:"BsisPtSprd,omitempty"`
}

func (v *VariableInterest1Rate) SetIndex(value string) {
	v.Index = (*Max35Text)(&value)
}

func (v *VariableInterest1Rate) SetBasisPointSpread(value string) {
	v.BasisPointSpread = (*Number)(&value)
}
