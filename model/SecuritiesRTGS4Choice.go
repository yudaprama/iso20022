package model

// Choice of format for the securities RTGS information.
type SecuritiesRTGS4Choice struct {

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Securities RTGS information expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SecuritiesRTGS4Choice) SetIndicator(value string) {
	s.Indicator = (*YesNoIndicator)(&value)
}

func (s *SecuritiesRTGS4Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
