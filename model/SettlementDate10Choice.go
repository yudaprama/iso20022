package model

// Choice of format for the settlement date.
type SettlementDate10Choice struct {

	// Date in ISO format.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Date and time at which the securities are to be delivered or received.
	DateCode *SettlementDateCode8Choice `xml:"DtCd"`
}

func (s *SettlementDate10Choice) AddDate() *DateAndDateTimeChoice {
	s.Date = new(DateAndDateTimeChoice)
	return s.Date
}

func (s *SettlementDate10Choice) AddDateCode() *SettlementDateCode8Choice {
	s.DateCode = new(SettlementDateCode8Choice)
	return s.DateCode
}
