package model

// Specifies the components of a settlement rate source for a non delvierable trade.
type SettlementRateSource1 struct {

	// Specifies the rate source for the non deliverable trade.
	RateSource *RateSourceText `xml:"RateSrc"`

	// Specifies the time "HHMM" associated with the rate source.
	Time *Exact4NumericText `xml:"Tm,omitempty"`

	// Specifies the country code for the quoted rate source.
	CountryCode *CountryCode `xml:"CtryCd,omitempty"`

	// The location expressed as a 2 character code.
	LocationCode *Exact2AlphaNumericText `xml:"LctnCd,omitempty"`
}

func (s *SettlementRateSource1) SetRateSource(value string) {
	s.RateSource = (*RateSourceText)(&value)
}

func (s *SettlementRateSource1) SetTime(value string) {
	s.Time = (*Exact4NumericText)(&value)
}

func (s *SettlementRateSource1) SetCountryCode(value string) {
	s.CountryCode = (*CountryCode)(&value)
}

func (s *SettlementRateSource1) SetLocationCode(value string) {
	s.LocationCode = (*Exact2AlphaNumericText)(&value)
}
