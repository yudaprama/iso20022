package model

// Summation of the call amounts either due to A or due to B.
type Result1 struct {

	// Amount payable by party B to party A.
	DueToPartyA *ActiveCurrencyAndAmount `xml:"DueToPtyA,omitempty"`

	// Amount payable by party A to party B.
	DueToPartyB *ActiveCurrencyAndAmount `xml:"DueToPtyB,omitempty"`

	// Provides additional information related to the collateral that may be requested.
	AdditionalInformation *Max210Text `xml:"AddtlInf,omitempty"`
}

func (r *Result1) SetDueToPartyA(value, currency string) {
	r.DueToPartyA = NewActiveCurrencyAndAmount(value, currency)
}

func (r *Result1) SetDueToPartyB(value, currency string) {
	r.DueToPartyB = NewActiveCurrencyAndAmount(value, currency)
}

func (r *Result1) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max210Text)(&value)
}
