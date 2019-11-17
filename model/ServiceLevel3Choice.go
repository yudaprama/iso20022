package model

// Specifies the service level of the transaction.
type ServiceLevel3Choice struct {

	// Identification of a pre-agreed level of service between the parties in a coded form.
	Code *ServiceLevel2Code `xml:"Cd"`

	// Proprietary identification of a pre-agreed level of service between the parties.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (s *ServiceLevel3Choice) SetCode(value string) {
	s.Code = (*ServiceLevel2Code)(&value)
}

func (s *ServiceLevel3Choice) SetProprietary(value string) {
	s.Proprietary = (*Max35Text)(&value)
}
