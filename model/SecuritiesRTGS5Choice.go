package model

// Choice of format for the securities RTGS information.
type SecuritiesRTGS5Choice struct {

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Securities RTGS information expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SecuritiesRTGS5Choice) SetIndicator(value string) {
	s.Indicator = (*YesNoIndicator)(&value)
}

func (s *SecuritiesRTGS5Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
