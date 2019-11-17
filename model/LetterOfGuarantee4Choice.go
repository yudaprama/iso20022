package model

// Choice of format for the letter of guarantee information.
type LetterOfGuarantee4Choice struct {

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Letter of guarantee information expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (l *LetterOfGuarantee4Choice) SetIndicator(value string) {
	l.Indicator = (*YesNoIndicator)(&value)
}

func (l *LetterOfGuarantee4Choice) AddProprietary() *GenericIdentification30 {
	l.Proprietary = new(GenericIdentification30)
	return l.Proprietary
}
