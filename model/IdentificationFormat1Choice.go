package model

// Choice between either a short, long or a proprietary identification format.
type IdentificationFormat1Choice struct {

	// Format expressed as a short identification.
	ShortIdentification *Exact3UpperCaseAlphaNumericText `xml:"ShrtId"`

	// Format expressed as a long identification.
	LongIdentification *Max30Text `xml:"LngId"`

	// Format expressed as a proprietary identification.
	ProprietaryIdentification *GenericIdentification19 `xml:"PrtryId"`
}

func (i *IdentificationFormat1Choice) SetShortIdentification(value string) {
	i.ShortIdentification = (*Exact3UpperCaseAlphaNumericText)(&value)
}

func (i *IdentificationFormat1Choice) SetLongIdentification(value string) {
	i.LongIdentification = (*Max30Text)(&value)
}

func (i *IdentificationFormat1Choice) AddProprietaryIdentification() *GenericIdentification19 {
	i.ProprietaryIdentification = new(GenericIdentification19)
	return i.ProprietaryIdentification
}
