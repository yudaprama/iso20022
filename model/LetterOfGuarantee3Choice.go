package model

// Choice of format for the letter of guarantee information.
type LetterOfGuarantee3Choice struct {

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Letter of guarantee information expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (l *LetterOfGuarantee3Choice) SetIndicator(value string) {
	l.Indicator = (*YesNoIndicator)(&value)
}

func (l *LetterOfGuarantee3Choice) AddProprietary() *GenericIdentification38 {
	l.Proprietary = new(GenericIdentification38)
	return l.Proprietary
}
