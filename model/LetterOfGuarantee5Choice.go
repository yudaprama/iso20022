package model

// Choice of format for the letter of guarantee information.
type LetterOfGuarantee5Choice struct {

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Letter of guarantee information expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (l *LetterOfGuarantee5Choice) SetIndicator(value string) {
	l.Indicator = (*YesNoIndicator)(&value)
}

func (l *LetterOfGuarantee5Choice) AddProprietary() *GenericIdentification47 {
	l.Proprietary = new(GenericIdentification47)
	return l.Proprietary
}
