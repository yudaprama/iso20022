package model

// Choice of format for the settlement date.
type SettlementDate2Choice struct {

	// Date in ISO format.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Date and time at which the securities are to be delivered or received.
	DateCode *SettlementDateCode2Choice `xml:"DtCd"`
}

func (s *SettlementDate2Choice) AddDate() *DateAndDateTimeChoice {
	s.Date = new(DateAndDateTimeChoice)
	return s.Date
}

func (s *SettlementDate2Choice) AddDateCode() *SettlementDateCode2Choice {
	s.DateCode = new(SettlementDateCode2Choice)
	return s.DateCode
}
