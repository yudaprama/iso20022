package model

// Specifies the service level of the transaction.
type ServiceLevel8Choice struct {

	// Specifies a pre-agreed service or level of service between the parties, as published in an external service level code list.
	Code *ExternalServiceLevel1Code `xml:"Cd"`

	// Specifies a pre-agreed service or level of service between the parties, as a proprietary code.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (s *ServiceLevel8Choice) SetCode(value string) {
	s.Code = (*ExternalServiceLevel1Code)(&value)
}

func (s *ServiceLevel8Choice) SetProprietary(value string) {
	s.Proprietary = (*Max35Text)(&value)
}
