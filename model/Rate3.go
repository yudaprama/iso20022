package model

// Set of elements used to qualify the interest rate.
type Rate3 struct {

	// Specifies the type of interest rate.
	Type *RateType4Choice `xml:"Tp"`

	// An amount range where the interest rate is applicable.
	ValidityRange *CurrencyAndAmountRange2 `xml:"VldtyRg,omitempty"`
}

func (r *Rate3) AddType() *RateType4Choice {
	r.Type = new(RateType4Choice)
	return r.Type
}

func (r *Rate3) AddValidityRange() *CurrencyAndAmountRange2 {
	r.ValidityRange = new(CurrencyAndAmountRange2)
	return r.ValidityRange
}
