package model

// Choice of format for the settlement date.
type SettlementDate12Choice struct {

	// Date in ISO format.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Date and time at which the securities are to be delivered or received.
	DateCode *SettlementDateCode9Choice `xml:"DtCd"`
}

func (s *SettlementDate12Choice) AddDate() *DateAndDateTimeChoice {
	s.Date = new(DateAndDateTimeChoice)
	return s.Date
}

func (s *SettlementDate12Choice) AddDateCode() *SettlementDateCode9Choice {
	s.DateCode = new(SettlementDateCode9Choice)
	return s.DateCode
}
