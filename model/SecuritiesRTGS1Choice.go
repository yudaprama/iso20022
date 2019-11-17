package model

// Choice of format for the securities RTGS information.
type SecuritiesRTGS1Choice struct {

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Securities RTGS information expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SecuritiesRTGS1Choice) SetIndicator(value string) {
	s.Indicator = (*YesNoIndicator)(&value)
}

func (s *SecuritiesRTGS1Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
