package model

// Information needed due to regulatory and statutory requirements.
type StructuredRegulatoryReporting3 struct {

	// Specifies the type of the information supplied in the regulatory reporting details.
	Type *Max35Text `xml:"Tp,omitempty"`

	// Date related to the specified type of regulatory reporting details.
	Date *ISODate `xml:"Dt,omitempty"`

	// Country related to the specified type of regulatory reporting details.
	Country *CountryCode `xml:"Ctry,omitempty"`

	// Specifies the nature, purpose, and reason for the transaction to be reported for regulatory and statutory requirements in a coded form.
	Code *Max10Text `xml:"Cd,omitempty"`

	// Amount of money to be reported for regulatory and statutory requirements.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`

	// Additional details that cater for specific domestic regulatory requirements.
	Information []*Max35Text `xml:"Inf,omitempty"`
}

func (s *StructuredRegulatoryReporting3) SetType(value string) {
	s.Type = (*Max35Text)(&value)
}

func (s *StructuredRegulatoryReporting3) SetDate(value string) {
	s.Date = (*ISODate)(&value)
}

func (s *StructuredRegulatoryReporting3) SetCountry(value string) {
	s.Country = (*CountryCode)(&value)
}

func (s *StructuredRegulatoryReporting3) SetCode(value string) {
	s.Code = (*Max10Text)(&value)
}

func (s *StructuredRegulatoryReporting3) SetAmount(value, currency string) {
	s.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *StructuredRegulatoryReporting3) AddInformation(value string) {
	s.Information = append(s.Information, (*Max35Text)(&value))
}
