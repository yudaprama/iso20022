package model

// Choice of format for the settlement date.
type SettlementDate3Choice struct {

	// Date in ISO format.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Date and time at which the securities are to be delivered or received.
	DateCode *GenericIdentification20 `xml:"DtCd"`
}

func (s *SettlementDate3Choice) AddDate() *DateAndDateTimeChoice {
	s.Date = new(DateAndDateTimeChoice)
	return s.Date
}

func (s *SettlementDate3Choice) AddDateCode() *GenericIdentification20 {
	s.DateCode = new(GenericIdentification20)
	return s.DateCode
}
