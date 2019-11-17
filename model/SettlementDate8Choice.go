package model

// Choice of format for the settlement date.
type SettlementDate8Choice struct {

	// Numeric representation of the day of the month and year.
	Date *DateAndDateTime1Choice `xml:"Dt"`

	// Date and time at which the securities are to be delivered or received.
	Code *SettlementDateCode5Choice `xml:"Cd"`
}

func (s *SettlementDate8Choice) AddDate() *DateAndDateTime1Choice {
	s.Date = new(DateAndDateTime1Choice)
	return s.Date
}

func (s *SettlementDate8Choice) AddCode() *SettlementDateCode5Choice {
	s.Code = new(SettlementDateCode5Choice)
	return s.Code
}
