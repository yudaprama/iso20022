package model

// Provides the details of a debt instrument in which the periodic interest payments are calculated on the basis of the value (that is, fixing of an underlying reference rate such as the Euribor) on predefined dates (that is, fixing dates) and which has a maturity of no more than one year.
type FloatingRateNote2 struct {

	// Underlying reference rate on the basis on which the periodic interest payments are calculated.
	ReferenceRateIndex *ISINOct2015Identifier `xml:"RefRateIndx"`

	// Number of basis points added to (if positive) or deducted from (if negative) the underlying reference rate to calculate the actual interest rate applicable for a given period at issuance of the floating rate instrument.
	//
	// Used to express differences in interest rates, for example, a difference of 0.10% is equivalent to a change of 10 basis points.
	BasisPointSpread *Number `xml:"BsisPtSprd"`
}

func (f *FloatingRateNote2) SetReferenceRateIndex(value string) {
	f.ReferenceRateIndex = (*ISINOct2015Identifier)(&value)
}

func (f *FloatingRateNote2) SetBasisPointSpread(value string) {
	f.BasisPointSpread = (*Number)(&value)
}
