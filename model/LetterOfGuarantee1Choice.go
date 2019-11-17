package model

// Choice of format for the letter of guarantee information.
type LetterOfGuarantee1Choice struct {

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Letter of guarantee information expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (l *LetterOfGuarantee1Choice) SetIndicator(value string) {
	l.Indicator = (*YesNoIndicator)(&value)
}

func (l *LetterOfGuarantee1Choice) AddProprietary() *GenericIdentification20 {
	l.Proprietary = new(GenericIdentification20)
	return l.Proprietary
}
