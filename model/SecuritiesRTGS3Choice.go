package model

// Choice of format for the securities RTGS information.
type SecuritiesRTGS3Choice struct {

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Securities RTGS information expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (s *SecuritiesRTGS3Choice) SetIndicator(value string) {
	s.Indicator = (*YesNoIndicator)(&value)
}

func (s *SecuritiesRTGS3Choice) AddProprietary() *GenericIdentification38 {
	s.Proprietary = new(GenericIdentification38)
	return s.Proprietary
}
