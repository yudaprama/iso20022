package model

// Specification of the fund order type.
type FundOrderType1 struct {

	// Structured format.
	Structured *FundOrderType2Code `xml:"Strd"`

	// Additional information about the type of identification
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (f *FundOrderType1) SetStructured(value string) {
	f.Structured = (*FundOrderType2Code)(&value)
}

func (f *FundOrderType1) SetAdditionalInformation(value string) {
	f.AdditionalInformation = (*Max350Text)(&value)
}
