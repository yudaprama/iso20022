package model

// Provides the index used to define the rate and optionally the basis point spread.
type FloatingInterestRate4 struct {

	// Identifies the reference index for the debt instrument.
	//
	// Usage:
	// Where an identifier exists, the ISIN must be used; otherwise, names will be necessary (such as EURIBOR6M, LIBOR3M) as other identification.
	ReferenceRate *BenchmarkCurveName4Choice `xml:"RefRate"`

	// Term of the index.
	Term *InterestRateContractTerm1 `xml:"Term"`

	// Provides the number of basis points added to (if positive) or deducted from (if negative) the underlying reference rate to calculate the actual interest rate applicable for a given period at issuance of the floating rate instrument.
	//
	// Used to express differences in interest rates, for example, a difference of 0.10% is equivalent to a change of 10 basis points.
	BasisPointSpread *Number `xml:"BsisPtSprd"`
}

func (f *FloatingInterestRate4) AddReferenceRate() *BenchmarkCurveName4Choice {
	f.ReferenceRate = new(BenchmarkCurveName4Choice)
	return f.ReferenceRate
}

func (f *FloatingInterestRate4) AddTerm() *InterestRateContractTerm1 {
	f.Term = new(InterestRateContractTerm1)
	return f.Term
}

func (f *FloatingInterestRate4) SetBasisPointSpread(value string) {
	f.BasisPointSpread = (*Number)(&value)
}
