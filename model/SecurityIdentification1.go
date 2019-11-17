package model

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type SecurityIdentification1 struct {

	// Identification of a security by an ISIN.
	Identification *SecurityIdentification7 `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm"`

	// Features of units offered by a fund. For example, a unit may have a specific load structure, eg, front end or back end, an income policy, eg, pay out or accumulate, or a trailer policy, eg, with or without. Fund classes are typically denoted by a single character, eg, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Name of the umbrella fund in which financial instrument is contained.
	UmbrellaName *Max35Text `xml:"UmbrllNm,omitempty"`

	// Currency of the investment fund class.
	BaseCurrency *ActiveCurrencyCode `xml:"BaseCcy"`

	// Country where the fund has legal domicile as reflected in the ISIN classification.
	CountryOfDomicile *CountryCode `xml:"CtryOfDmcl"`

	// Countries where the fund is registered for distribution.
	RegisteredDistributionCountry []*CountryCode `xml:"RegdDstrbtnCtry"`
}

func (s *SecurityIdentification1) AddIdentification() *SecurityIdentification7 {
	s.Identification = new(SecurityIdentification7)
	return s.Identification
}

func (s *SecurityIdentification1) SetName(value string) {
	s.Name = (*Max350Text)(&value)
}

func (s *SecurityIdentification1) SetClassType(value string) {
	s.ClassType = (*Max35Text)(&value)
}

func (s *SecurityIdentification1) SetUmbrellaName(value string) {
	s.UmbrellaName = (*Max35Text)(&value)
}

func (s *SecurityIdentification1) SetBaseCurrency(value string) {
	s.BaseCurrency = (*ActiveCurrencyCode)(&value)
}

func (s *SecurityIdentification1) SetCountryOfDomicile(value string) {
	s.CountryOfDomicile = (*CountryCode)(&value)
}

func (s *SecurityIdentification1) AddRegisteredDistributionCountry(value string) {
	s.RegisteredDistributionCountry = append(s.RegisteredDistributionCountry, (*CountryCode)(&value))
}
