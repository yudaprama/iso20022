package model

// Information needed due to regulatory and statutory requirements.
type StructuredRegulatoryReporting2 struct {

	// Specifies the nature, purpose, and reason for the transaction to be reported for regulatory and statutory requirements in a coded form.
	Code *Max3Text `xml:"Cd,omitempty"`

	// Amount of money to be reported for regulatory and statutory requirements.
	Amount *CurrencyAndAmount `xml:"Amt,omitempty"`

	// Additional details that cater for specific domestic regulatory requirements.
	//
	// Usage: Information is used to provide details that are not catered for in the Code or/and Amount elements.
	Information *Max35Text `xml:"Inf,omitempty"`
}

func (s *StructuredRegulatoryReporting2) SetCode(value string) {
	s.Code = (*Max3Text)(&value)
}

func (s *StructuredRegulatoryReporting2) SetAmount(value, currency string) {
	s.Amount = NewCurrencyAndAmount(value, currency)
}

func (s *StructuredRegulatoryReporting2) SetInformation(value string) {
	s.Information = (*Max35Text)(&value)
}
