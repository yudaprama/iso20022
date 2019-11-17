package model

// Choice of format for the settlement date.
type SettlementDate9Choice struct {

	// Date in ISO format.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Date and time at which the securities are to be delivered or received.
	DateCode *SettlementDateCode7Choice `xml:"DtCd"`
}

func (s *SettlementDate9Choice) AddDate() *DateAndDateTimeChoice {
	s.Date = new(DateAndDateTimeChoice)
	return s.Date
}

func (s *SettlementDate9Choice) AddDateCode() *SettlementDateCode7Choice {
	s.DateCode = new(SettlementDateCode7Choice)
	return s.DateCode
}
