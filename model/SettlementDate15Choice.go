package model

// Choice of format for the settlement date.
type SettlementDate15Choice struct {

	// Date in ISO format.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Date and time at which the securities are to be delivered or received.
	DateCode *SettlementDateCode11Choice `xml:"DtCd"`
}

func (s *SettlementDate15Choice) AddDate() *DateAndDateTimeChoice {
	s.Date = new(DateAndDateTimeChoice)
	return s.Date
}

func (s *SettlementDate15Choice) AddDateCode() *SettlementDateCode11Choice {
	s.DateCode = new(SettlementDateCode11Choice)
	return s.DateCode
}
