package model

// Choice of format for the settlement date.
type SettlementDate11Choice struct {

	// Date in ISO format.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Date and time at which the securities are to be delivered or received.
	DateCode *GenericIdentification30 `xml:"DtCd"`
}

func (s *SettlementDate11Choice) AddDate() *DateAndDateTimeChoice {
	s.Date = new(DateAndDateTimeChoice)
	return s.Date
}

func (s *SettlementDate11Choice) AddDateCode() *GenericIdentification30 {
	s.DateCode = new(GenericIdentification30)
	return s.DateCode
}
