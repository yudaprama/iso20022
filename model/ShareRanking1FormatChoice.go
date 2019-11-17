package model

// Choice of formats to  express the share ranking.
type ShareRanking1FormatChoice struct {

	// Standard code to specify the share ranking.
	Code *ShareRanking1Code `xml:"Cd"`

	// Proprietary code to express the share ranking.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (s *ShareRanking1FormatChoice) SetCode(value string) {
	s.Code = (*ShareRanking1Code)(&value)
}

func (s *ShareRanking1FormatChoice) AddProprietary() *GenericIdentification13 {
	s.Proprietary = new(GenericIdentification13)
	return s.Proprietary
}
