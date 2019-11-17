package model

// Specifies the service level of the transaction.
type ServiceLevel2Choice struct {

	// Identification of a pre-agreed level of service between the parties in a coded form.
	Code *ServiceLevel1Code `xml:"Cd"`

	// Proprietary identification of a pre-agreed level of service between the parties.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (s *ServiceLevel2Choice) SetCode(value string) {
	s.Code = (*ServiceLevel1Code)(&value)
}

func (s *ServiceLevel2Choice) SetProprietary(value string) {
	s.Proprietary = (*Max35Text)(&value)
}
