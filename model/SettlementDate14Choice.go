package model

// Choice of format for the settlement date.
type SettlementDate14Choice struct {

	// Date in ISO format.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Date and time at which the securities are to be delivered or received.
	DateCode *GenericIdentification47 `xml:"DtCd"`
}

func (s *SettlementDate14Choice) AddDate() *DateAndDateTimeChoice {
	s.Date = new(DateAndDateTimeChoice)
	return s.Date
}

func (s *SettlementDate14Choice) AddDateCode() *GenericIdentification47 {
	s.DateCode = new(GenericIdentification47)
	return s.DateCode
}
