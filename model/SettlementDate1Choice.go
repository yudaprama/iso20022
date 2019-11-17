package model

// Choice of format for the settlement date.
type SettlementDate1Choice struct {

	// Date in ISO format.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Date and time at which the securities are to be delivered or received.
	DateCode *SettlementDateCode1Choice `xml:"DtCd"`
}

func (s *SettlementDate1Choice) AddDate() *DateAndDateTimeChoice {
	s.Date = new(DateAndDateTimeChoice)
	return s.Date
}

func (s *SettlementDate1Choice) AddDateCode() *SettlementDateCode1Choice {
	s.DateCode = new(SettlementDateCode1Choice)
	return s.DateCode
}
